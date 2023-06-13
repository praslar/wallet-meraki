package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"wallet/internal/model"
	"wallet/internal/repo"
)

type WalletService struct {
<<<<<<< HEAD
	userRepo    repo.UserRepo
	authService AuthService
}

func NewWalletService(userRepo repo.UserRepo) WalletService {
	return WalletService{
		userRepo: userRepo,
=======
	WalletRepo  repo.WalletRepo
	authService AuthService
}

func NewWalletService(WalletRepo repo.WalletRepo, authService AuthService) WalletService {
	return WalletService{
		WalletRepo:  WalletRepo,
		authService: authService,
>>>>>>> f42f72261765b586a57e931f5a776a40c861c8d0
	}
}

func (s *WalletService) CreateWallet(name string, userID uuid.UUID) error {
<<<<<<< HEAD
=======

>>>>>>> f42f72261765b586a57e931f5a776a40c861c8d0
	newWallet := &model.Wallet{
		Name:   name,
		UserID: userID,
	}
<<<<<<< HEAD
	if err := s.userRepo.CreateWallet(newWallet); err != nil {
		logrus.Errorf("Failed to create new user: %s", err.Error())
		return fmt.Errorf("Internal server error. ")
=======
	if err := s.WalletRepo.CreateWallet(newWallet); err != nil {
		logrus.Errorf("Failed to create new wallet: %s", err.Error())
		return fmt.Errorf("Internal server error")
>>>>>>> f42f72261765b586a57e931f5a776a40c861c8d0
	}
	return nil

}
<<<<<<< HEAD
=======

func (s *WalletService) GetAllWallet() ([]model.Wallet, error) {
	wallet, err := s.WalletRepo.GetAllWallet()
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
>>>>>>> f42f72261765b586a57e931f5a776a40c861c8d0
