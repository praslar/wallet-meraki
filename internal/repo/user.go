package repo

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math"
	"wallet/internal/model"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return UserRepo{
		db: db,
	}
}

func (r *UserRepo) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepo) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Model(&model.User{}).Where("email = ?", email).Preload("Role").First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetUserByID(id string) (*model.User, error) {
	var user model.User
	if err := r.db.Model(&model.User{}).Preload("Role").
		Where("id = ?", id).Take(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) CheckEmailExist(newEmail string) bool {
	rs := model.User{}
	err := r.db.Model(&model.User{}).Where("email = ?", newEmail).First(&rs).Error
	if err != nil {
		// email chưa tôn tại, nên không tìm thấy
		return false
	}
	return true
}

func (r *UserRepo) GetRoleID(namerole string) (uuid.UUID, error) {
	var roleID model.Role
	err := r.db.Where("name = ?", namerole).First(&roleID).Error
	if err != nil {
		return uuid.Nil, err
	}

	return roleID.ID, nil
}


func (r *UserRepo) DeleteUser(userID uuid.UUID) error {
	user := model.User{}
	if err := r.db.Model(&user).Where("id = ?", userID).Take(&user).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) GetAllUsers(filterEmail string, sortOrder string, page int, limit int) ([]model.User, int, error) {
	var users []model.User

	// Create a query builder
	query := r.db.Model(&model.User{}).
		Preload("Role").
		Preload("Wallets").
		Offset((page - 1) * limit).
		Limit(limit)

	// Apply filtering if filterName is provided
	if filterEmail != "" {
		query = query.Where("email ILIKE ?", fmt.Sprintf("%%%s%%", filterEmail))
	}

	// Apply sorting if sortOrder is provided
	if sortOrder != "" {
		query = query.Order(fmt.Sprintf("email %s", sortOrder))
	}

	// Count the total number of users (without pagination)
	var count int64
	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	// Retrieve the users
	if err := query.Find(&users).Error; err != nil {
		return nil, 0, err
	}

	// Iterate over users and populate the wallet count
	for i := range users {
		user := &users[i]
		walletCount := len(user.Wallets)
		user.WalletCount = walletCount
	}

	// Calculate the total number of pages based on the limit
	totalPages := int(math.Ceil(float64(count) / float64(limit)))

	return users, totalPages, nil
}

func (r *UserRepo) GetUser(userID uuid.UUID) (*model.User, error) {
	var user model.User
	if err := r.db.Preload("Role").Preload("Wallets").First(&user, userID).Error; err != nil {
		return nil, err
	}

	walletCount := len(user.Wallets)
	user.WalletCount = walletCount

	return &user, nil
}

func (r *UserRepo) UpdateUserRole(userID uuid.UUID, role string) error {
	user := model.User{}
	if err := r.db.Model(&user).Where("id = ?", userID).Take(&user).Error; err != nil {
		return err
	}

	roleID, err := r.GetRoleID(role)
	if err != nil {
		return err
	}

	user.RoleID = roleID

	if err := r.db.Save(&user).Error; err != nil {
		return err
	}

	return nil

func (r *UserRepo) GetTransactionID(id string) ([]model.Transaction, error) {
	var data []model.Transaction
	tx := r.db.Preload("Token").Preload("WalletTo.User.Role").Preload("WalletFrom.User.Role")
	tx = tx.Table("transactions t").
		Joins("join wallets w on t.from_address = w.address").
		Joins("join users u on w.user_id = u.id").
		Where(" u.id = ?", id).Find(&data)

	return data, nil
}

func (r *UserRepo) GetAllTransaction(formWallet string, toWallet string, email string, tokenAddress string, orderBy string, amount int, pageSize int, page int) ([]model.Transaction, error) {
	var data []model.Transaction

	tx := r.db.Preload("Token").Preload("WalletTo.User.Role").Preload("WalletFrom.User.Role")

	//Xu li logic get all user
	if amount != 0 {
		tx = tx.Where("amount > ?", amount)
	}

	if orderBy != "" {
		tx = tx.Order(orderBy)
	}

	if formWallet != "" || toWallet != "" {
		tx = tx.Preload("user_id").Where("from_address = ? AND to_address = ?", formWallet, toWallet)
	}

	if email != "" {
		tx = tx.Table("transactions t").Joins(`join wallets w on t.from_address = w.address`).
			Joins(`join users u on w.user_id  = u.id `).
			Where(`email = ?`, email).Scan(&data)
	}

	if tokenAddress != "" {
		tx = tx.Where("token_address = ?", tokenAddress)
	}

	//xu li paging
	if err := tx.Limit(pageSize).Offset((page - 1) * pageSize).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r *UserRepo) CreateToken(newToken *model.Token) error {
	result := r.db.Create(&newToken)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepo) SymbolUnique(symbol string) bool {
	var token *model.Token
	result := r.db.Model(&model.Token{}).Where("symbol = ?", symbol).First(&token)
	if result == nil {
		logrus.Infof("token bi duplicate. Khong tao được ")
		return false
		// tim ko co thi chua co token
	}
	logrus.Infof("Khong tim thấy token. Co thể tao ")
	return true
	// tim co thi co token
}

func (r *UserRepo) DeleteToken(newToken *model.Token) error {
	result := r.db.Model(&newToken).Where("address = ? ", newToken.Address).Delete("symbol", newToken.Symbol)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepo) ValidateTokenInUse(tokenaddress uuid.UUID) bool {
	var transaction *model.Transaction
	result := r.db.Model(&model.Transaction{}).Where("token_address = ?", tokenaddress).First(&transaction)
	if result == nil {
		logrus.Infof("Token InUsed. Không Thể Xoá ")
		return false
	}
	logrus.Infof("Không tìm thấy Token InUsed. Có thể Xoá ")
	return true
	// tim co thi co token
}

func (r *UserRepo) UpdateToken(newToken *model.Token) error {
	result := r.db.Model(&newToken).Where("address = ?", newToken.Address).Updates(map[string]interface{}{"symbol": newToken.Symbol, "price": newToken.Price})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepo) SendUserToken(newtransaction *model.Transaction) error {
	if err := r.db.Create(&newtransaction).Error; err != nil {
		return fmt.Errorf("Failed to save transaction: %v. ", err)
	}
	return nil
}

func (r *UserRepo) ValidateWallet(address uuid.UUID) bool {
	wallet := &model.Wallet{}
	result := r.db.Model(wallet).Where("address", address).First(&wallet).Error
	if result != nil {
		logrus.Infof("Khong tìm thấy wallet. ")
		return false
	}
	return true

}

func (r *UserRepo) ValidateToken(address uuid.UUID) bool {
	token := &model.Token{}
	result := r.db.Model(token).Where("address", address).First(&token).Error
	if result != nil {
		logrus.Infof("Không tìm thấy token. ")
		return false
	}
	return true
}
