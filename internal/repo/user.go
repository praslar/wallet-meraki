package repo

import (
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

func (s *UserRepo) CreateUser(newUser *model.User) error {
	result := s.db.Create(newUser)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *UserRepo) CheckEmail(newEmail string) bool {
	rs := model.User{}
	err := s.db.Model(&model.User{}).Where("email = ?", newEmail).First(&rs).Error
	if err != nil {
		// email chưa tôn tại, nên không tìm thấy
		return false
	}
	return true
}
