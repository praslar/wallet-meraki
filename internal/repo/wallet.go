package repo

import (
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

func (r *WalletRepo) GetAllWallet() ([]model.Wallet, error) {
	rs := []model.Wallet{}
	if err := r.db.Find(&rs).Error; err != nil {
		return nil, err
	}
	return rs, nil
}

func (r *WalletRepo) DeleteWallet(address string) error {
	wallet := model.Wallet{}
	if err := r.db.Where("address = ?", address).First(&wallet).Error; err != nil {
		return err
	}
	if err := r.db.Delete(&wallet).Error; err != nil {
		return err
	}

	return nil
}

func (r *WalletRepo) CheckWalletExist(address string) bool {
	rs := model.Wallet{}
	err := r.db.Model(&model.Wallet{}).Where("address = ?", address).First(&rs).Error
	if err != nil {
		//  không tìm thấy ví
		return false
	}
	return true
}
func (s *WalletRepo) Update(address string, name string) (*model.Wallet, error) {
	wallet := &model.Wallet{}
	if err := s.db.Where("address = ?", address).First(&wallet).Error; err != nil {
		return nil, err
	}
	wallet.Name = name
	if err := s.db.Save(&wallet).Error; err != nil {
		return nil, err
	}
	return wallet, nil
}
