package repo

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"wallet/internal/model"
)

type TokenRepo struct {
	db *gorm.DB
}

type TokenRepoInterface interface {
	SymbolUnique(symbol string) bool
	CreateToken(newToken *model.Token) error
	UpdateToken(newToken *model.Token) error
	DeleteToken(newToken *model.Token) error
	ValidateTokenInUse(tokenaddress uuid.UUID) bool
}

func NewTokenRepo(db *gorm.DB) TokenRepoInterface {
	return &TokenRepo{
		db: db,
	}
}

func (r *TokenRepo) SymbolUnique(symbol string) bool {
	var token *model.Token
	result := r.db.Model(&model.Token{}).Where("symbol = ?", symbol).First(token)
	if result == nil {
		logrus.Infof("token bi duplicate. Khong tao được ")
		return false
		// tim ko co thi chua co token
	}
	logrus.Infof("Khong tim thấy token. Co thể tao ")
	return true
	// tim co thi co token
}

func (r *TokenRepo) CreateToken(newToken *model.Token) error {
	result := r.db.Create(&newToken)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *TokenRepo) UpdateToken(newToken *model.Token) error {
	result := r.db.Model(&newToken).Where("address = ?", newToken.Address).Updates(map[string]interface{}{"symbol": newToken.Symbol, "price": newToken.Price})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *TokenRepo) DeleteToken(newToken *model.Token) error {
	result := r.db.Model(&newToken).Where("address = ? ", newToken.Address).Delete("symbol", newToken.Symbol)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *TokenRepo) ValidateTokenInUse(tokenaddress uuid.UUID) bool {
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
