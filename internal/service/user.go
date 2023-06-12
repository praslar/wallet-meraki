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

<<<<<<< HEAD
	userRoleID, _ := uuid.Parse("91a65d49-19ef-4306-b702-2ee0a850e7b2")
	newUser := &model.User{
		Email:    email,
		Password: password,
		RoleID:   userRoleID,
=======
	//TODO: get role_id from database
	//roleID, _ := uuid.Parse("f943bd28-ea93-4638-abc4-cfc3d278fd32")

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
>>>>>>> f42f72261765b586a57e931f5a776a40c861c8d0
	}
	if s.userRepo.CheckEmailExist(email) {
		return fmt.Errorf("Email existed")
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
<<<<<<< HEAD
		return "", fmt.Errorf("Internal server error. ")
=======
		return "", fmt.Errorf("Internal server error")
>>>>>>> f42f72261765b586a57e931f5a776a40c861c8d0
	}
	return token, nil
}

func (s *UserService) GetAllUser() ([]model.User, error) {
	users, err := s.userRepo.GetAllUser()
	if err != nil {
<<<<<<< HEAD
		return nil, fmt.Errorf("Internal server error. ")
	}
	return users, nil
}
=======
		return nil, fmt.Errorf("Internal server error")
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
>>>>>>> f42f72261765b586a57e931f5a776a40c861c8d0
