package repo

import "wallet/internal/model"

func (r *Repo) CreateToken(token []model.Token) error {
	return r.Postgres.Model(model.Token{}).Create(token).Error
}
