package service

import (
	"wallet/internal/model"
	"wallet/internal/repo"
)

type CoingeckoService struct {
	userRepo repo.UserRepo
}

func NewCoingeckoService(userRepo repo.UserRepo) CoingeckoService {
	return CoingeckoService{
		userRepo: userRepo,
	}
}

func (s *CoingeckoService) GetCoinInfo(coin *model.Token) error {
	err := s.userRepo.GetCoinInfo(coin)
	if err != nil {
		return err
	}
	return nil
}
