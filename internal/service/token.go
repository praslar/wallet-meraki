package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"wallet/internal/model"
	"wallet/internal/repo"
	"wallet/internal/utils"
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
	if !s.userRepo.SymbolUnique(symbol) {
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

func (s *TokenService) UpdateToken(address uuid.UUID, symbol string, price float64) error {

	newToken := &model.Token{
		Address: address,
		Symbol:  symbol,
		Price:   price,
	}
	if err := s.userRepo.UpdateToken(newToken); err != nil {
		logrus.Errorf("Failed To Update Token: %s", err.Error())
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

	if !s.userRepo.ValidateWallet(senderWalletAddress) {
		logrus.Errorf("Sender wallet not found. ")
		return fmt.Errorf("Sender wallet not found. ")
	}

	if !s.userRepo.ValidateWallet(receiverWalletAddress) {
		logrus.Errorf("Receiver wallet not found. ")
		return fmt.Errorf("Receiver wallet not found. ")
	}

	if !s.userRepo.ValidateToken(tokenAddress) {
		logrus.Errorf("Token not found. ")
		return fmt.Errorf("Token not found. ")
	}
	if newtransaction.Amount < 0 {
		fmt.Print("Amount must be larger than 0. ")
	}

	if err := s.userRepo.SendUserToken(newtransaction); err != nil {
		logrus.Errorf("Transaction Failed. %s", err.Error())
		return fmt.Errorf("Internal server error. ")
	}
	return nil
}

func (s *TokenService) AirdropTokenNewWallet(receiverWalletAddress uuid.UUID) error {
	var amount float64 = 1000
	senderWalletAddress, _ := uuid.Parse(utils.ADMIN_WALLET_ADDRESS)
	tokenAddress, _ := uuid.Parse(utils.TOKEN_WALLET_ADDRESS)
	airdroptransaction := &model.Transaction{
		FromAddress:  senderWalletAddress,
		ToAddress:    receiverWalletAddress,
		TokenAddress: tokenAddress,
		Amount:       amount,
	}

	if !s.userRepo.ValidateWallet(senderWalletAddress) {
		logrus.Errorf("Sender wallet not found. ")
		return fmt.Errorf("Sender wallet not found. ")
	}

	if !s.userRepo.ValidateWallet(receiverWalletAddress) {
		logrus.Errorf("Receiver wallet not found. ")
		return fmt.Errorf("Receiver wallet not found. ")
	}
	if !s.userRepo.ValidateToken(tokenAddress) {
		logrus.Errorf("Token not found. ")
		return fmt.Errorf("Token not found. ")
	}
	if s.userRepo.ValidateTokenWasSend(tokenAddress, amount) {
		logrus.Errorf("AirDrop Fail. ")
		return fmt.Errorf("This wallet had been taken an airdrop. ")
	}

	if err := s.userRepo.AirdropTokenNewWallet(airdroptransaction); err != nil {
		logrus.Errorf("Failed To Airdrop Token: %s", err.Error())
		return fmt.Errorf("Internal server error. ")
	}
	return nil
}
