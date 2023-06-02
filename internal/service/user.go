package service

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"wallet/internal/model"
	"wallet/internal/repo"
	"wallet/internal/utils"
)

type UserService struct {
	userRepo repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) UserService {
	return UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Register(email string, password string) error {
	newUser := &model.User{
		Email:    email,
		Password: password,
	}
	if len(password) < utils.MIN_PASSWORD_LEN {
		return fmt.Errorf("Min length password: %v", utils.MIN_PASSWORD_LEN)
	}

	if !utils.ValidEmail(email) {
		return fmt.Errorf("Wrong email format")
	}

	if err := s.userRepo.CreateUser(newUser); err != nil {
		logrus.Errorf("Failed to create new user: %s", err.Error())
		return fmt.Errorf("Internal server error")
	}
	return nil
}
