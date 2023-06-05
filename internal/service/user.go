package service

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/mail"
	"wallet/internal/model"
	"wallet/internal/repo"
	"wallet/utils"
)

type UserService struct {
	repo repo.UserRepo
}

func NewUserService(repo repo.UserRepo) UserService {
	return UserService{repo: repo}
}

// ---------------------- CheckFormEmail ----------------------

func (s *UserService) CheckEmailFormat(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		fmt.Println("Wrong Email Format. ")
		return false
	}
	return true
}

func (s *UserService) UserLogin(email string, password string) (*model.User, error) {
	user := &model.User{}

	if email == "" || password == "" {
		logrus.Infof("Email and password are required. ")
		err := fmt.Errorf("Email and password are required. ")
		return nil, err
	} else if !s.CheckEmailFormat(email) {
		logrus.Infof("Wrong Email Format. ")
		err := fmt.Errorf("Wrong Email Format. ")
		return nil, err
	} else if len(password) < utils.LenPassword {
		logrus.Infof("Length Password Can't Be Shorter Than 6. ")
		err := fmt.Errorf("Length Password Can't Be Shorter Than 6. ")
		return nil, err
	}
	_, _, err := s.repo.UserLogin(email, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) CreateToken() string {
	return s.repo.CreateToken()

}
