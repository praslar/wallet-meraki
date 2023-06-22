package repo

import (
	"wallet/internal/model"
)

func (r *Repo) CreateUser(user *model.User) error {
	return r.Postgres.Create(user).Error
}

func (r *Repo) GetAllUser() ([]model.User, error) {
	rs := []model.User{}
	if err := r.Postgres.Find(&rs).Error; err != nil {
		return nil, err
	}
	return rs, nil
}

func (r *Repo) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.Postgres.Model(&model.User{}).Where("email = ?", email).Take(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repo) GetUserByID(id string) (*model.User, error) {
	var user model.User
	if err := r.Postgres.Model(&model.User{}).Where("id = ?", id).Take(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
