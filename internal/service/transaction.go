package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"wallet/internal/model"
	"wallet/internal/repo"
	"wallet/internal/utils"
)

type TransactionService struct {
	transactionRepo repo.TransactionRepo
}

func NewTransactionService(transactionRepo repo.TransactionRepo) TransactionServiceInterface {
	return &TransactionService{
		transactionRepo: transactionRepo,
	}

}

type TransactionServiceInterface interface {
	CreateTransaction(fromWalletID uuid.UUID, toWalletID uuid.UUID, tokenID uuid.UUID, amount float64) error
	AirDropNewWallet(toWalletID uuid.UUID) error
}

func (s *TransactionService) CreateTransaction(fromWalletID uuid.UUID, toWalletID uuid.UUID, tokenID uuid.UUID, amount float64) error {
	newtransaction := &model.Transaction{
		FromAddress:  fromWalletID,
		ToAddress:    toWalletID,
		Amount:       amount,
		TokenAddress: tokenID,
	}

	if newtransaction.Amount < 0 {
		fmt.Print("Amount must be larger than 0. ")
	}

	if err := s.transactionRepo.CreateTransaction(newtransaction); err != nil {
		logrus.Errorf("Failed to create new user: %s", err.Error())
		return fmt.Errorf("Internal server error. ")
	}
	return nil
}

func (s *TransactionService) AirDropNewWallet(toWalletID uuid.UUID) error {

	var amount float64 = 1000

	senderWalletAddress, _ := uuid.Parse(utils.ADMIN_WALLET_ADDRESS)
	tokenAddress, _ := uuid.Parse(utils.MERAKI_TOKEN_ADDRESS)

	newtransaction := &model.Transaction{
		FromAddress:  senderWalletAddress,
		ToAddress:    toWalletID,
		Amount:       amount,
		TokenAddress: tokenAddress,
	}

	if newtransaction.Amount < 0 {
		fmt.Print("Amount must be larger than 0. ")
	}

	if err := s.transactionRepo.CreateTransaction(newtransaction); err != nil {
		logrus.Errorf("Failed to create new user: %s", err.Error())
		return fmt.Errorf("Internal server error. ")
	}
	return nil
}
