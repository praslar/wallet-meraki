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

	newWallet := &model.Wallet{
		Name:   name,
		UserID: userID,
	}

	if s.WalletRepo.CheckWalletExist(name) {
		return fmt.Errorf("Wallet existed")
	}

	if err := s.WalletRepo.CreateWallet(newWallet); err != nil {
		logrus.Errorf("Failed to create new wallet: %s", err.Error())
		return fmt.Errorf("Internal server error")
	}
	return nil

}
func (s *WalletService) GetOneWallet(name string) ([]model.Wallet, error) {
	exists := s.WalletRepo.CheckWalletExist(name)
	if !exists {
		return nil, fmt.Errorf("Wallet not found")
	}
	wallet, err := s.WalletRepo.GetOneWallet(name)
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

func (s *WalletService) DeleteWallet(name string) error {
	exists := s.WalletRepo.CheckWalletExist(name)
	if !exists {
		return fmt.Errorf("Wallet not found")
	}
	err := s.WalletRepo.DeleteWallet(name)
	if err != nil {
		return err
	}
	return nil
}
func (s *WalletService) UpdateWallet(userid uuid.UUID, name string) (*model.Wallet, error) {

	wallet, err := s.WalletRepo.Update(userid, name)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}
