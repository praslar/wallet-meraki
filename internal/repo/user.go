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

func (r *UserRepo) CreateToken(newToken *model.Token) error {
	result := r.db.Create(&newToken)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepo) SymbolUnique(symbol string) bool {
	var token *model.Token
	result := r.db.Model(&model.Token{}).Where("symbol = ?", symbol).Find(&token)
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

func (r *UserRepo) ValidateTokenInUse(tokenaddress uuid.UUID) bool {
	var transaction *model.Transaction
	result := r.db.Model(&model.Transaction{}).Where("token_address = ?", tokenaddress).Find(&transaction)
	if result != nil {
		logrus.Infof("Không tìm thấy Token InUsed. ")
		return false
	}
	logrus.Infof("Token InUsed. ")
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
