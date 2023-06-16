package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"wallet/internal/model"
	"wallet/internal/repo"
)

type TokenService struct {
	userRepo repo.UserRepo
}

func NewTokenService(userRepo repo.UserRepo) TokenService {
	return TokenService{
		userRepo: userRepo,
	}
}

func (s *TokenService) CreateToken(symbol string, price float64) error {

	newToken := &model.Token{
		Symbol: symbol,
		Price:  price,
	}
	if s.userRepo.SymbolUnique(symbol) {
		logrus.Errorf("This token was duplicated. ")
		return fmt.Errorf("This token was duplicated. ")
	}
	if err := s.userRepo.CreateToken(newToken); err != nil {
		logrus.Errorf("Failed to create new token: %s", err.Error())
		return fmt.Errorf("Internal server error. ")
	}
	return nil

}

func (s *TokenService) DeleteToken(tokenaddress uuid.UUID) error {
	newToken := &model.Token{
		Address: tokenaddress,
	}
	if !s.userRepo.ValidateTokenInUse(tokenaddress) {
		logrus.Errorf("Failed to delete token. Token InUse. ")
		return fmt.Errorf("Internal server error. ")
	}
	if err := s.userRepo.DeleteToken(newToken); err != nil {
		logrus.Errorf("Failed to delete token. : %s", err.Error())
		return fmt.Errorf("Internal server error. ")
	}
	return nil
}
