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

func (s *TokenService) CreateToken(symbol string, totalsupply uint64) error {

	newToken := &model.Token{
		Symbol:      symbol,
		TotalSupply: totalsupply,
	}
	if err := s.userRepo.CreateToken(newToken); err != nil {
		logrus.Errorf("Failed to create new user: %s", err.Error())
		return fmt.Errorf("Internal server error. ")
	}
	return nil

}

func (s *TokenService) UpdateToken(address uuid.UUID, symbol string) error {

	newToken := &model.Token{
		Address: address,
		Symbol:  symbol,
	}
	if err := s.userRepo.UpdateToken(newToken); err != nil {
		logrus.Errorf("Failed to create new user: %s", err.Error())
		return fmt.Errorf("Internal server error. ")
	}
	return nil

}

func (s *TokenService) DeleteToken(address uuid.UUID, symbol string) error {

	newToken := &model.Token{
		Address: address,
		Symbol:  symbol,
	}
	if err := s.userRepo.DeleteToken(newToken); err != nil {
		logrus.Errorf("Failed to create new user: %s", err.Error())
		return fmt.Errorf("Internal server error. ")
	}
	return nil

}

//func (s *TokenService) TransferTokenAd(senderWalletAddress uuid.UUID, receiverWalletAddress uuid.UUID, tokenID uuid.UUID, amount float64) error {
//
//	// Create a new token with the specified symbol, wallet address, and amount
//	token := model.Token{
//		TokenID:       tokenID,
//		WalletAddress: senderWalletAddress,
//		Amount:        amount,
//	}
//
//	// Create a new transaction with the sender wallet address, receiver wallet address, token ID, and amount
//	transaction := model.Transaction{
//		SenderWalletAddress:   senderWalletAddress,
//		ReceiverWalletAddress: receiverWalletAddress,
//		TokenID:               token.TokenID,
//		Amount:                amount,
//	}
//
//	if err := s.userRepo.TransferTokenAd(&token, &transaction); err != nil {
//		logrus.Errorf("Failed to create new user: %s", err.Error())
//		return fmt.Errorf("Internal server error. ")
//	}
//	return nil
//
//}
