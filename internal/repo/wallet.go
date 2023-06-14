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

func (r *WalletRepo) GetAllWallet(name string, userID string, pageSize, page int) ([]model.Wallet, error) {
	rs := []model.Wallet{}
	tx := r.db.Preload("User")

	if name != "" {
		tx = tx.Where("name ILIKE ?", "%"+name+"%")
	}

	if userID != "" {
		tx = tx.Where("user_id = ?", "%"+userID+"%")
	}

	if err := r.db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&rs).Error; err != nil {
		return nil, err
	}
	return rs, nil
}

func (r *WalletRepo) DeleteWallet(name string) error {
	var wallet model.Wallet
	// Tìm ví
	if err := r.db.Where("name = ?", name).First(&wallet).Error; err != nil {
		return err
	}
	// Xóa ví khỏi cơ sở dữ liệu
	if err := r.db.Delete(&wallet).Error; err != nil {
		return err
	}

	return nil
}

func (r *WalletRepo) CheckWalletExist(name string) bool {
	rs := model.Wallet{}
	err := r.db.Model(&model.Wallet{}).Where("name = ?", name).First(&rs).Error
	if err != nil {
		return false
	}
	return true
}
func (s *WalletRepo) Update(userid uuid.UUID, name string) (*model.Wallet, error) {
	var wallet model.Wallet
	if err := s.db.Model(&model.Wallet{}).Where("user_id = ?", userid).Update("name", name).Error; err != nil {
		return nil, err
	}

	return &wallet, nil
}
func (r *UserRepo) CheckWalletExist(name string) bool {
	rs := model.Wallet{}
	err := r.db.Model(&model.Wallet{}).Where("name = ?", name).Find(&rs).Error
	if err != nil {
		// email chưa tôn tại, nên không tìm thấy
		return false
	}
	return true
}
