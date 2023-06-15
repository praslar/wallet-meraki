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

func (s *WalletService) CreateWallet(name string, userID uuid.UUID) error {

	if s.WalletRepo.CheckWalletExist(name) {
		return fmt.Errorf("Wallet existed")
	}

	newWallet := &model.Wallet{
		Name:   name,
		UserID: userID,
	}

	if err := s.WalletRepo.CreateWallet(newWallet); err != nil {
		logrus.Errorf("Failed to create new wallet: %s", err.Error())
		return fmt.Errorf("Internal server error")
	}
	return nil

}