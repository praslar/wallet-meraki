package service

import (
	"net/http"
	"wallet/internal/repo"
)

type MigrateService struct {
	repo repo.MigrateRepo
}

func NewMigrateService(repo repo.MigrateRepo) MigrateService {
	return MigrateService{
		repo: repo}
}

func (s *MigrateService) Migrate(w http.ResponseWriter, r *http.Request) {
	s.repo.Migrate(w, r)
}
