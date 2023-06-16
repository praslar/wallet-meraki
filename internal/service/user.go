package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"wallet/internal/model"
	"wallet/internal/repo"
	"wallet/internal/utils"
	"wallet/pkg/jwt"
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

	//TODO: get role_id from database
	// Good
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return fmt.Errorf("lỗi trong quá trình mã hóa password: %v", err)
	}

	roleID, err := s.GetRoleID("user")
	if err != nil {
		return fmt.Errorf("lỗi khi lấy ID của vai trò: %v", err)
	}
	newUser := &model.User{
		Email:    email,
		Password: hashedPassword,
		RoleID:   roleID,
	}
	if s.userRepo.CheckEmailExist(email) {
		return fmt.Errorf("Email existed")
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

	token, err := jwt.GenJWTToken(user.ID, user.Role.Key)
	if err != nil {
		logrus.Errorf("Failed to generate token: %s", err.Error())
		return "", fmt.Errorf("Internal server error")
	}
	return token, nil
}

func (s *UserService) GetAllUser(orderBy string) ([]model.User, error) {
	users, err := s.userRepo.GetAllUser(orderBy)
	if err != nil {
		return nil, fmt.Errorf("Internal server error")
	}
	return users, nil
}

func (s *UserService) GetUserByID(id string) (*model.User, error) {
	users, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) GetRoleID(name string) (uuid.UUID, error) {
	roleID, err := s.userRepo.GetRoleID(name)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Role not found: %v", err)
	}
	return roleID, nil
}

func (s *UserService) DeleteUser(userID uuid.UUID) error {
	err := s.userRepo.DeleteUser(userID)
	if err != nil {
		return err
	}
	return nil
}
