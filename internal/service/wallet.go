package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"wallet/internal/model"
	"wallet/internal/repo"
)

type WalletService struct {
	walletRepo repo.WalletRepo
	transSrv   TransactionServiceInterface
}

func NewWalletService(walletRepo repo.WalletRepo, transSrv TransactionServiceInterface) WalletServiceInterface {
	return &WalletService{
		walletRepo: walletRepo,
		transSrv:   transSrv,
	}
}

type WalletServiceInterface interface {
	CreateWallet(name string, xuserid string) error
	GetOneWallet(userID string, name string) ([]model.Wallet, error)
	GetAllWallet(order string, name string, userID string, pageSize, page int) ([]model.Wallet, error)
	DeleteWallet(userId string, name string) error
	UpdateWallet(userid string, name string, updateName string) ([]model.Wallet, error)
	GetUserWalletAddress(userid string, name string) uuid.UUID
}

func (s *WalletService) CreateWallet(name string, xuserid string) error {

	userUUID, err := uuid.Parse(xuserid)
	if err != nil {
		return fmt.Errorf("Invalid x-user-id", err)
	}

	// CREATE WALLET
	err = s.walletRepo.CheckWalletExist(name)
	if err == nil {
		return fmt.Errorf("Wallet existed")
	}

	newWallet := &model.Wallet{
		Name:   name,
		UserID: userUUID,
	}

	if err := s.walletRepo.CreateWallet(newWallet); err != nil {
		logrus.Errorf("Failed to create new wallet: %s", err.Error())
		return fmt.Errorf("Internal server error")
	}

	// Airdrop
	newWalletAddress := s.GetUserWalletAddress(xuserid, name)
	if newWalletAddress == uuid.Nil {
		return fmt.Errorf("wallet not found")
	}

	err = s.transSrv.AirDropNewWallet(newWalletAddress)
	if err != nil {
		return err
	}
	return nil
}

func (s *WalletService) GetOneWallet(userID string, name string) ([]model.Wallet, error) {
	err := s.walletRepo.CheckWalletExist(name)
	if err != nil {
		return nil, fmt.Errorf("User dont have any wallet ")
	}
	wallet, err := s.walletRepo.GetOneWallet(name, userID)
	if err != nil {
		return nil, err
	}
	return wallet, nil

}

func (s *WalletService) GetAllWallet(order string, name string, userID string, pageSize, page int) ([]model.Wallet, error) {
	wallet, err := s.walletRepo.GetAllWallet(order, name, userID, pageSize, page)
	if err != nil {
		return nil, fmt.Errorf("Internal server error")
	}
	return wallet, nil
}

func (s *WalletService) DeleteWallet(userId string, name string) error {
	err := s.walletRepo.CheckWalletExist(name)
	if err != nil {
		return fmt.Errorf("User dont have any wallet")
	}
	s.walletRepo.DeleteWallet(userId, name)
	return nil
}

func (s *WalletService) UpdateWallet(userid string, name string, updateName string) ([]model.Wallet, error) {
	wallet, err := s.walletRepo.Update(userid, name, updateName)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func (s *WalletService) GetUserWalletAddress(userid string, name string) uuid.UUID {
	return s.walletRepo.GetUserWalletAddress(userid, name)
}
