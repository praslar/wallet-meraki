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

func (s *WalletService) CreateWallet(address string, name string, userID uuid.UUID) error {

	newWallet := &model.Wallet{
		Address: address,
		Name:    name,
		UserID:  userID,
	}
	if err := s.WalletRepo.CreateWallet(newWallet); err != nil {
		logrus.Errorf("Failed to create new user: %s", err.Error())
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
