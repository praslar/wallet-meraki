package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"wallet/internal/model"
	"wallet/internal/repo"
)

type WalletService struct {
	WalletRepo  repo.WalletRepo
	authService AuthService
}

func NewWalletService(WalletRepo repo.WalletRepo, authService AuthService) WalletService {
	return WalletService{
		WalletRepo:  WalletRepo,
		authService: authService,
	}
}

func (s *WalletService) CreateWallet(address string, name string, userID uuid.UUID) error {

	newWallet := &model.Wallet{
		Address: address,
		Name:    name,
		UserID:  userID,
	}
	if err := s.WalletRepo.CreateWallet(newWallet); err != nil {
		logrus.Errorf("Failed to create new wallet: %s", err.Error())
		return fmt.Errorf("Internal server error")
	}
	return nil

}

func (s *WalletService) GetAllWallet() ([]model.Wallet, error) {
	wallet, err := s.WalletRepo.GetAllWallet()
	if err != nil {
		return nil, fmt.Errorf("Internal server error")
	}
	return wallet, nil
}

func (s *WalletService) DeleteWallet(address string) error {
	exists := s.WalletRepo.CheckWalletExist(address)
	if exists {
		return fmt.Errorf("Wallet not found")
	}
	err := s.WalletRepo.DeleteWallet(address)
	if err != nil {
		return err
	}
	return nil
}
func (s *WalletService) UpdateWallet(address string, name string) (*model.Wallet, error) {
	wallet, err := s.WalletRepo.Update(address, name)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}
