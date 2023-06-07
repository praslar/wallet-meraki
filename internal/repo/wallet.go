package repo

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"wallet/internal/model"
)

type WalletRepo struct {
	db *gorm.DB
}

func NewWalletRepo(db *gorm.DB) WalletRepo {
	return WalletRepo{
		db: db,
	}
}

func (r *WalletRepo) CreateWallet(newWallet *model.Wallet) error {
	result := r.db.Create(newWallet)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *WalletRepo) GetUserIDByEmail(email string) (uuid.UUID, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return uuid.Nil, err
	}

	return user.ID, nil
}
