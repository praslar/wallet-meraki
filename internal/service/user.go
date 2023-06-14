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
		// Handle password length validation
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

func (s *UserService) GetAllUsers() ([]model.User, error) {
	users, err := s.userRepo.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}
	return users, nil
}

func (s *UserService) GetRoleID(name string) (uuid.UUID, error) {
	roleID, err := s.userRepo.GetRoleID(name)
	if err != nil {
		return uuid.Nil, fmt.Errorf("role not found: %v", err)
	}
	return roleID, nil
}

func (s *UserService) UpdateUserRole(userID int, role string) error {
	err := s.userRepo.UpdateUserRole(userID, role)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) DeleteUser(userID int) error {
	err := s.userRepo.DeleteUser(userID)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUser(userID int) (*model.User, error) {
	user, err := s.userRepo.GetUser(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
