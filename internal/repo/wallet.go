package repo

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"wallet/internal/model"
)

type WalletRepo struct {
	db *gorm.DB
}

//func NewWalletRepo(db *gorm.DB) WalletRepo {
//	return WalletRepo{
//		db: db,
//	}
//}

func NewWalletRepo(db *gorm.DB) WalletRepoInterface {
	return &WalletRepo{db: db}

}

type WalletRepoInterface interface {
	CreateWallet(newWallet *model.Wallet) error
	CheckWalletExist(name string) error
	GetOneWallet(name string, userID string) ([]model.Wallet, error)
	GetAllWallet(order string, name string, userID string, pageSize, page int) ([]model.Wallet, error)
	DeleteWallet(userId string, name string) error
	Update(userid string, name string, updateName string) ([]model.Wallet, error)
	AirdropToken(airdroptransaction *model.Transaction) error
	GetUserWalletAddress(userid string, name string) uuid.UUID
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

func (r *WalletRepo) Update(userid string, name string, updateName string) ([]model.Wallet, error) {
	var wallet []model.Wallet
	if err := r.db.Model(&model.Wallet{}).Preload("User").Preload("User.Role").Where("user_id = ? AND name = ?", userid, name).Find(&wallet).Error; err != nil {
		return nil, err
	}
	for _, wallet := range wallet {
		wallet.Name = updateName
		err := r.db.Save(&wallet).Error
		if err != nil {
			return nil, err
		}
	}
	return wallet, nil
}

func (r *WalletRepo) AirdropToken(airdroptransaction *model.Transaction) error {
	if err := r.db.Create(&airdroptransaction).Error; err != nil {
		return fmt.Errorf("Failed to save transaction: %v. ", err)
	}
	return nil
}

func (r *WalletRepo) GetUserWalletAddress(userid string, name string) uuid.UUID {
	var wallet model.Wallet
	if err := r.db.Where("user_id = ? AND name = ?", userid, name).Find(&wallet).Error; err != nil {
		return uuid.Nil
	}
	return wallet.Address
}
