package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"wallet/internal/model"
	"wallet/internal/repo"
	"wallet/internal/utils"
)

type WalletService struct {
	WalletRepo repo.WalletRepo
	userRepo   repo.UserRepo
}

func NewWalletService(userRepo repo.UserRepo, WalletRepo repo.WalletRepo) WalletService {
	return WalletService{
		WalletRepo: WalletRepo,
		userRepo:   userRepo,
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
	err = s.AirdropToken(s.GetUserWalletAddress(xuserid, name))
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

func (s *WalletService) AirdropToken(recieverWalletAddress uuid.UUID) error {

	var amount float64 = 1000
	senderWalletAddress, _ := uuid.Parse(utils.ADMIN_WALLET_ADDRESS)
	tokenAddress, _ := uuid.Parse(utils.TOKEN_WALLET_ADDRESS)
	airdroptransaction := &model.Transaction{
		FromAddress:  senderWalletAddress,
		ToAddress:    recieverWalletAddress,
		TokenAddress: tokenAddress,
		Amount:       amount,
	}
	if err := s.WalletRepo.AirdropToken(airdroptransaction); err != nil {
		logrus.Errorf("Failed to create new wallet: %s", err.Error())
		return fmt.Errorf("Internal server error")
	}
	return nil
}
