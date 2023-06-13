package repo

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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
	return r.db.Create(&user).Error
}

func (r *UserRepo) GetAllUser() ([]model.User, error) {
	var rs []model.User
	if err := r.db.Find(&rs).Error; err != nil {
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
	var user *model.User
	fmt.Print(r.db.Name())
	if err := r.db.Model(&model.User{}).Where("id = ?", id).Take(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepo) CreateWallet(newWallet *model.Wallet) error {
	result := r.db.Create(&newWallet)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepo) CreateToken(newToken *model.Token) error {
	result := r.db.Create(&newToken)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepo) SymbolUnit(symbol string) bool {
	token := &model.Token{}
	result := r.db.Model(token).Where("symbol", symbol).First(token).Error
	if result != nil {
		logrus.Infof("Khong tim thấy token. ")
		return false
		// tim ko co thi chua co token
	}
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

func (r *UserRepo) UpdateToken(newToken *model.Token) error {
	result := r.db.Model(&newToken).Where("address = ?", newToken.Address).Updates(map[string]interface{}{"symbol": newToken.Symbol, "price": newToken.Price})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepo) SendUserToken(newtransaction *model.Transaction) error {

	// Save the new token and transaction to the database
	if err := r.db.Create(&newtransaction).Error; err != nil {
		return fmt.Errorf("Failed to save transaction: %v. ", err)
	}

	return nil
}

func (r *UserRepo) ValidateWallet(address uuid.UUID) bool {
	wallet := &model.Wallet{}
	result := r.db.Model(wallet).Where("address", address).First(wallet).Error
	if result != nil {
		logrus.Infof("Khong tìm thấy wallet. ")
		return false
		// tim ko co thi chua co token
	}
	return true
	// tim co thi co token
}

func (r *UserRepo) ValidateToken(address uuid.UUID) bool {
	token := &model.Token{}
	result := r.db.Model(token).Where("address", address).First(token).Error
	if result != nil {
		logrus.Infof("Không tìm thấy token. ")
		return false
		// tim ko co thi chua co token
	}
	return true
	// tim co thi co token
}

func (r *UserRepo) ValidateTokenInUse(address uuid.UUID) bool {
	transaction := &model.Transaction{}
	result := r.db.Model(transaction).Where("address", transaction.TokenAddress).First(transaction).Error
	if result != nil {
		logrus.Infof("Không tìm thấy token InUse. ")
		return false
		// tim ko co thi chua co token
	}
	return true
	// tim co thi co token
}
