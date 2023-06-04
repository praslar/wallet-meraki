package service

import (
	"fmt"
	"wallet/internal/model"
	"wallet/internal/repo"
)

type UserService struct {
	repo repo.UserRepo
}

func NewUserService(repo repo.UserRepo) UserService {
	return UserService{repo: repo}
}

func (s *UserService) SignUp(email string, password string) error {
	if len(password) < model.PasswordLength {
		return fmt.Errorf("Password must > %v character", 6)
	}
	err := s.repo.CreateUser(&model.User{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) CheckEmail(email string) (bool, error) {
	exists := s.repo.CheckEmail(email)
	return exists, nil
}
