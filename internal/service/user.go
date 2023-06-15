package service

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"wallet/internal/model"
	"wallet/internal/repo"
	"wallet/internal/utils"
)

type UserService struct {
	userRepo    repo.UserRepo
	authService AuthService
}

func NewUserService(userRepo repo.UserRepo, authService AuthService) UserService {
	return UserService{
		userRepo:    userRepo,
		authService: authService,
	}
}

func (s *UserService) Register(email string, password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return fmt.Errorf("error hashing password: %v", err)
	}

	roleID, err := s.userRepo.GetRoleID("user")
	if err != nil {
		return fmt.Errorf("error getting role ID: %v", err)
	}

	newUser := &model.User{
		Email:    email,
		Password: hashedPassword,
		RoleID:   roleID,
	}

	if s.userRepo.CheckEmailExist(email) {
		return fmt.Errorf("email already exists")
	}

	if len(password) < model.PasswordLength {
	}

	if !utils.ValidEmail(email) {
		return fmt.Errorf("wrong email format")
	}

	if err := s.userRepo.CreateUser(newUser); err != nil {
		logrus.Errorf("failed to create new user: %s", err.Error())
		return fmt.Errorf("internal server error")
	}

	return nil
}

func (s *UserService) Login(email string, password string) (string, error) {
	if !utils.ValidEmail(email) {
		return "", fmt.Errorf("wrong email format")
	}

	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		logrus.Errorf("failed to get user by email: %s", err.Error())
		return "", fmt.Errorf("user not found")
	}

	if !utils.ComparePassword(password, user.Password) {
		return "", fmt.Errorf("wrong password")
	}

	token, err := s.authService.GenJWTToken(user.ID.String(), user.Role.Key)
	if err != nil {
		logrus.Errorf("failed to generate token: %s", err.Error())
		return "", fmt.Errorf("internal server error")
	}

	return token, nil
}
