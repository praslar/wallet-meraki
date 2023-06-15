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

func (s *TokenService) CreateToken(symbol string, price float64) error {

	newToken := &model.Token{
		Symbol: symbol,
		Price:  price,
	}
	if !s.SymbolUnit(symbol) {
		logrus.Errorf("This token was duplicated. ")
		return fmt.Errorf("This token was duplicated. ")
	}
	if err := s.userRepo.CreateToken(newToken); err != nil {
		logrus.Errorf("Failed to create new token: %s", err.Error())
		return fmt.Errorf("Internal server error. ")
	}
	return nil

}

func (s *TokenService) SymbolUnit(symbol string) bool {
	if s.userRepo.SymbolUnique(symbol) {
	}
	return true
}

func (s *TokenService) DeleteToken(tokenaddress uuid.UUID) error {

	newToken := &model.Token{
		Address: tokenaddress,
	}
	if !s.ValidateTokenInUse(tokenaddress) {
		logrus.Errorf("Failed to delete token. Token InUse. ")
		return fmt.Errorf("Internal server error. ")
	}

	if err := s.userRepo.DeleteToken(newToken); err != nil {
		logrus.Errorf("Failed to delete token. : %s", err.Error())
		return fmt.Errorf("Internal server error. ")
	}
	return nil

}

func (s *TokenService) ValidateTokenInUse(tokenaddress uuid.UUID) bool {
	if s.userRepo.ValidateTokenInUse(tokenaddress) {
	}
	return true
}
func (s *TokenService) UpdateToken(address uuid.UUID) error {

	newToken := &model.Token{
		Address: address,
	}
	if err := s.userRepo.UpdateToken(newToken); err != nil {
		logrus.Errorf("Failed to create new user: %s", err.Error())
		return fmt.Errorf("Internal server error. ")
	}
	return nil

}

func (s *TokenService) SendUserToken(senderWalletAddress uuid.UUID, receiverWalletAddress uuid.UUID, tokenAddress uuid.UUID, amount float64) error {

	newtransaction := &model.Transaction{
		FromAddress:  senderWalletAddress,
		ToAddress:    receiverWalletAddress,
		Amount:       amount,
		TokenAddress: tokenAddress,
	}

	if !s.ValidateWallet(senderWalletAddress) {
		logrus.Errorf("Sender wallet not found. ")
		return fmt.Errorf("Sender wallet not found. ")
	}

	if !s.ValidateWallet(receiverWalletAddress) {
		logrus.Errorf("Receiver wallet not found. ")
		return fmt.Errorf("Receiver wallet not found. ")
	}

	if !s.ValidateToken(tokenAddress) {
		logrus.Errorf("Token not found. ")
		return fmt.Errorf("Token not found. ")
	}
	if newtransaction.Amount < 0 {
		fmt.Print("Amount must be larger than 0. ")
	}

	if err := s.userRepo.SendUserToken(newtransaction); err != nil {
		logrus.Errorf("Failed to create new user: %s", err.Error())
		return fmt.Errorf("Internal server error. ")
	}
	return nil
}

func (s *TokenService) ValidateWallet(address uuid.UUID) bool {
	if s.userRepo.ValidateWallet(address) {
	}
	return true
}

func (s *TokenService) ValidateToken(address uuid.UUID) bool {
	if s.userRepo.ValidateWallet(address) {
	}
	return true
}
