package service

import "wallet/internal/repo"

type TokenService struct {
	TokenRepo repo.UserRepo
}

func NewTokenService(TokenRepo repo.UserRepo) TokenService {
	return TokenService{
		TokenRepo: TokenRepo,
	}
}
