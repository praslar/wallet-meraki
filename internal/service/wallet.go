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
}

func NewWalletService(WalletRepo repo.WalletRepo) WalletService {
	return WalletService{
		WalletRepo: WalletRepo,
	}
}

func (s *WalletService) CreateWallet(name string, xuserid string) error {

	err := s.WalletRepo.CheckWalletExist(name)
	if err == nil {
		return fmt.Errorf("Wallet existed")
	}
	xUserIDuuid, err := uuid.Parse(xuserid)
	if err != nil {
		return fmt.Errorf("Invalid x-user-id", err)
	}
	newWallet := &model.Wallet{
		Name:   name,
		UserID: xUserIDuuid,
	}

	if err := s.WalletRepo.CreateWallet(newWallet); err != nil {
		logrus.Errorf("Failed to create new wallet: %s", err.Error())
		return fmt.Errorf("Internal server error")
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
