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

func (s *WalletService) CreateWallet(userID uuid.UUID, name string) error {

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
