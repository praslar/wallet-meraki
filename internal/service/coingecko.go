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

func (s *CoingeckoService) GetCoinInfo(symbol string, price float64) error {

	newcoin := &model.Token{
		Symbol: symbol,
		Price:  price,
	}
	err := s.userRepo.GetCoinInfo(newcoin)
	if err != nil {
		return err
	}
	return nil
}
