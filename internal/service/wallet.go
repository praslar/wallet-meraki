package service

import (
	"fmt"
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

func (s *WalletService) CreateWallet(email string, address string, name string) error {

	userID, err := s.WalletRepo.GetUserIDByEmail(email)
	if err != nil {
		return fmt.Errorf("Role not found")
	}

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
