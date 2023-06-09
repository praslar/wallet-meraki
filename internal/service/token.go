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

func (s *TokenService) CreateTokenAd(walletAddress uuid.UUID, symbol string, amount float64) error {

	newToken := &model.Token{
		WalletAddress: walletAddress,
		Symbol:        symbol,
		Amount:        amount,
	}
	if err := s.userRepo.CreateTokenAd(newToken); err != nil {
		logrus.Errorf("Failed to create new user: %s", err.Error())
		return fmt.Errorf("Internal server error. ")
	}
	return nil

}

func (s *TokenService) UpdateTokenAd(walletAddress uuid.UUID, tokenID uuid.UUID, symbol string) error {

	newToken := &model.Token{
		WalletAddress: walletAddress,
		TokenID:       tokenID,
		Symbol:        symbol,
	}
	if err := s.userRepo.UpdateTokenAd(newToken); err != nil {
		logrus.Errorf("Failed to create new user: %s", err.Error())
		return fmt.Errorf("Internal server error. ")
	}
	return nil

}

func (s *TokenService) DeleteTokenAd(walletAddress uuid.UUID, tokenID uuid.UUID) error {

	newToken := &model.Token{
		WalletAddress: walletAddress,
		TokenID:       tokenID,
	}
	if err := s.userRepo.DeleteTokenAd(newToken); err != nil {
		logrus.Errorf("Failed to create new user: %s", err.Error())
		return fmt.Errorf("Internal server error. ")
	}
	return nil

}

func (s *TokenService) TransferTokenAd(senderWalletAddress uuid.UUID, receiverWalletAddress uuid.UUID, tokenID uuid.UUID, amount float64) error {

	// Create a new token with the specified symbol, wallet address, and amount
	token := model.Token{
		TokenID:       tokenID,
		WalletAddress: senderWalletAddress,
		Amount:        amount,
	}

	// Create a new transaction with the sender wallet address, receiver wallet address, token ID, and amount
	transaction := model.Transaction{
		SenderWalletAddress:   senderWalletAddress,
		ReceiverWalletAddress: receiverWalletAddress,
		TokenID:               token.TokenID,
		Amount:                amount,
	}

	if err := s.userRepo.TransferTokenAd(&token, &transaction); err != nil {
		logrus.Errorf("Failed to create new user: %s", err.Error())
		return fmt.Errorf("Internal server error. ")
	}
	return nil

}
