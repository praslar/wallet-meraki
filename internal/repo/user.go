package repo

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (r *UserRepo) GetAllUser(orderBy string) ([]model.User, error) {
	rs := []model.User{}
	if err := r.db.Preload("Role").Order(orderBy).Find(&rs).Error; err != nil {
		return nil, err
	}
	return rs, nil
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

func (r *UserRepo) GetTransactionID(id string) ([]model.Transaction, error) {
	var data []model.Transaction
	if err := r.db.Table("transactions t").
		Joins("join wallets w on t.from_address = w.address").
		Joins("join users u on w.user_id = u.id").
		Where(" u.id = ?", id).Scan(&data).Error; err != nil {
		return nil, err
	}
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
