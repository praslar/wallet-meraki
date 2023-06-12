package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"wallet/internal/model"
	"wallet/internal/repo"
	"wallet/internal/utils"
)

type UserService struct {
	userRepo    repo.UserRepo
	authService AuthService
}

func NewUserService(userRepo repo.UserRepo) UserService {
	return UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Register(email string, password string) error {

	userRoleID, _ := uuid.Parse("91a65d49-19ef-4306-b702-2ee0a850e7b2")
	newUser := &model.User{
		Email:    email,
		Password: password,
		RoleID:   userRoleID,
	}
	if len(password) < utils.MIN_PASSWORD_LEN {
		return fmt.Errorf("Min length password: %v. ", utils.MIN_PASSWORD_LEN)
	}

	if !utils.ValidEmail(email) {
		return fmt.Errorf("Wrong email format. ")
	}

	if err := s.userRepo.CreateUser(newUser); err != nil {
		logrus.Errorf("Failed to create new user: %s", err.Error())
		return fmt.Errorf("Internal server error. ")
	}
	return nil
}

func (s *UserService) Login(email string, password string) (string, error) {

	if !utils.ValidEmail(email) {
		return "", fmt.Errorf("wrong email format")
	}

	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		logrus.Errorf("Failed to get user by email: %s", err.Error())
		return "", fmt.Errorf("user not found")
	}

	if !utils.ComparePassword(password, user.Password) {
		return "", fmt.Errorf("wrong password")
	}

	token, err := s.authService.GenJWTToken(user.ID.String(), user.Role.Key)
	if err != nil {
		logrus.Errorf("Failed to generate token: %s", err.Error())
		return "", fmt.Errorf("Internal server error. ")
	}
	return token, nil
}

func (s *UserService) GetAllUser() ([]model.User, error) {
	users, err := s.userRepo.GetAllUser()
	if err != nil {
		return nil, fmt.Errorf("Internal server error. ")
	}
	return users, nil
}
