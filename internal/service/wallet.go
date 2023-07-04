package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"wallet/internal/model"
	"wallet/internal/repo"
)

type WalletService struct {
	WalletRepo repo.WalletRepo
	TransSrv   TransactionServiceInterface
}

func NewWalletService(walletRepo repo.WalletRepo, transSrv TransactionServiceInterface) WalletServiceInterface {
	return &WalletService{
		WalletRepo: walletRepo,
		TransSrv:   transSrv,
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
	err = s.WalletRepo.CheckWalletExist(name)
	if err == nil {
		return fmt.Errorf("Wallet existed")
	}

	newWallet := &model.Wallet{
		Name:   name,
		UserID: userUUID,
	}

	if err := s.WalletRepo.CreateWallet(newWallet); err != nil {
		logrus.Errorf("Failed to create new wallet: %s", err.Error())
		return fmt.Errorf("Internal server error")
	}

	// Airdrop
	newWalletAddress := s.GetUserWalletAddress(xuserid, name)
	if newWalletAddress == uuid.Nil {
		return fmt.Errorf("wallet not found")
	}

	err = s.TransSrv.AirDropNewWallet(newWalletAddress)
	if err != nil {
		return err
	}
	return nil
}

func (s *WalletService) GetOneWallet(userID string, name string) ([]model.Wallet, error) {
	err := s.WalletRepo.CheckWalletExist(name)
	if err != nil {
		return nil, fmt.Errorf("User dont have any wallet ")
	}
	wallet, err := s.WalletRepo.GetOneWallet(name, userID)
	if err != nil {
		return nil, err
	}
	return wallet, nil

}

func (s *WalletService) GetAllWallet(order string, name string, userID string, pageSize, page int) ([]model.Wallet, error) {
	wallet, err := s.WalletRepo.GetAllWallet(order, name, userID, pageSize, page)
	if err != nil {
		return nil, fmt.Errorf("Internal server error")
	}
	return wallet, nil
}

func (s *WalletService) DeleteWallet(userId string, name string) error {
	err := s.WalletRepo.CheckWalletExist(name)
	if err != nil {
		return fmt.Errorf("User dont have any wallet")
	}
	s.WalletRepo.DeleteWallet(userId, name)
	return nil
}

func (s *WalletService) UpdateWallet(userid string, name string, updateName string) ([]model.Wallet, error) {
	wallet, err := s.WalletRepo.Update(userid, name, updateName)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func (s *WalletService) GetUserWalletAddress(userid string, name string) uuid.UUID {
	return s.WalletRepo.GetUserWalletAddress(userid, name)
}
