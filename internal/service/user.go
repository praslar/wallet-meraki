package service

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"wallet/internal/auth"
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
func (s *UserService) Login(email string, password string) (string, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		logrus.Errorf("Failed to check email :%s", err.Error())
		return "", fmt.Errorf("internal server error")
	}
	if user == nil {
		return "", fmt.Errorf("internal server error")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("invalid email or password")
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return "", fmt.Errorf("failed to generate token")
	}
	return token, nil
}
