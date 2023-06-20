package repo

import (
	"fmt"
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

func (r *WalletRepo) CheckWalletExist(name string) error {
	rs := model.Wallet{}
	err := r.db.Model(&model.Wallet{}).Where("name = ?", name).First(&rs).Error
	if err != nil {
		return fmt.Errorf("Wallet does not exist")
	}
	// Wallet tồn tại
	return nil
}

func (r *WalletRepo) GetOneWallet(name string, userID string) ([]model.Wallet, error) {
	rs := []model.Wallet{}
	if err := r.db.Preload("User").Preload("User.Role").Where("name = ? AND user_id = ?", name, userID).First(&rs).Error; err != nil {
		return nil, err
	}
	return rs, nil
}

func (r *WalletRepo) GetAllWallet(order string, name string, userID string, pageSize, page int) ([]model.Wallet, error) {
	rs := []model.Wallet{}
	tx := r.db.Preload("User").Preload("User.Role")

	if name != "" {
		tx = tx.Where("name ILIKE ?", "%"+name+"%")
	}

	if userID != "" {
		tx = tx.Where("user_id = ?", userID)
	}

	if err := tx.Order(order).Limit(pageSize).Offset((page - 1) * pageSize).Find(&rs).Error; err != nil {
		return nil, err
	}
	return rs, nil
}

func (r *WalletRepo) DeleteWallet(userId string, name string) error {
	var wallet model.Wallet
	if err := r.db.Preload("User").Preload("User.Role").Where("user_id = ? AND name = ?", userId, name).First(&wallet).Error; err != nil {
		return err
	}
	if err := r.db.Delete(&wallet).Error; err != nil {
		return err
	}

	return nil
}
