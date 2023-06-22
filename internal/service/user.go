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
	Repo        repo.IRepo
	authService AuthService
}

func NewUserService(repo repo.IRepo, authService AuthService) UserService {
	return UserService{
		Repo:        repo,
		authService: authService,
	}
}

func (s *UserService) Register(email string, password string) error {

	//TODO: get role_id from database
	// Good
	userRoleID, _ := uuid.Parse("5c042680-2227-457d-b4fd-cccd5b09c658")
	newUser := &model.User{
		Email:    email,
		Password: password,
		RoleID:   userRoleID,
	}
	if len(password) < utils.MIN_PASSWORD_LEN {
		return fmt.Errorf("Min length password: %v", utils.MIN_PASSWORD_LEN)
	}

	if !utils.ValidEmail(email) {
		return fmt.Errorf("Wrong email format")
	}

	if err := s.Repo.CreateUser(newUser); err != nil {
		logrus.Errorf("Failed to create new user: %s", err.Error())
		return fmt.Errorf("Internal server error")
	}
	return nil
}

func (s *UserService) Login(email string, password string) (string, error) {

	if !utils.ValidEmail(email) {
		return "", fmt.Errorf("wrong email format")
	}

	user, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		logrus.Errorf("Failed to get user by email: %s", err.Error())
		return "", fmt.Errorf("user not found")
	}

	if !utils.ComparePassword(password, user.Password) {
		return "", fmt.Errorf("wrong password")
	}

	token, err := s.authService.GenJWTToken(user.ID.String())
	if err != nil {
		logrus.Errorf("Failed to generate token: %s", err.Error())
		return "", fmt.Errorf("Internal server error")
	}
	return token, nil
}

func (s *UserService) GetAllUser() ([]model.User, error) {
	users, err := s.Repo.GetAllUser()
	if err != nil {
		return nil, fmt.Errorf("Internal server error")
	}
	return users, nil
}
