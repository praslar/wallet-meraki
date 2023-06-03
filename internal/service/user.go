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
	if email == "" || password == "" {
		err := fmt.Errorf("Email and password are required. ")
		return nil, err
	} else if !s.CheckEmailFormat(email) {
		logrus.Infof("Wrong Email Format. ")
		err := fmt.Errorf("Wrong Email Format. ")
		return nil, err
	} else if len(password) < utils.LenPassword {
		err := fmt.Errorf("Length Password Can't Be Shorter Than 6. ")
		return nil, err
	}
	_, err := s.repo.UserLogin(email, password)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
