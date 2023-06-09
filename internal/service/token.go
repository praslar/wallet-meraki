package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"wallet/internal/model"
	"wallet/internal/repo"
)

type TokenService struct {
	userRepo    repo.UserRepo
	authService AuthService
}

func NewTokenService(userRepo repo.UserRepo) TokenService {
	return TokenService{
		userRepo: userRepo,
	}
}

func (s *TokenService) CreateToken(walletAddress uuid.UUID, symbol string) error {

	newToken := &model.Token{
		WalletAddress: walletAddress,
		Symbol:        symbol,
	}
	if err := s.userRepo.CreateToken(newToken); err != nil {
		logrus.Errorf("Failed to create new user: %s", err.Error())
		return fmt.Errorf("Internal server error. ")
	}
	return nil

}

func (s *TokenService) UpdateToken(walletaddress uuid.UUID, tokenID uuid.UUID, symbol string) error {

	newToken := &model.Token{
		WalletAddress: walletaddress,
		TokenID:       tokenID,
		Symbol:        symbol,
	}
	if err := s.userRepo.UpdateToken(newToken); err != nil {
		logrus.Errorf("Failed to create new user: %s", err.Error())
		return fmt.Errorf("Internal server error. ")
	}
	return nil

}

func (s *TokenService) DeleteToken(walletaddress uuid.UUID, tokenID uuid.UUID) error {

	newToken := &model.Token{
		WalletAddress: walletaddress,
		TokenID:       tokenID,
	}
	if err := s.userRepo.DeleteToken(newToken); err != nil {
		logrus.Errorf("Failed to create new user: %s", err.Error())
		return fmt.Errorf("Internal server error. ")
	}
	return nil

}
