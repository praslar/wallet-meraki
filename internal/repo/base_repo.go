package repo

import (
	"gorm.io/gorm"
	"wallet/internal/model"
)

type Repo struct {
	Postgres *gorm.DB
}

func NewRepo(db *gorm.DB) IRepo {
	return &Repo{Postgres: db}
}

func (r *Repo) GetRepo() *gorm.DB {
	return r.Postgres
}

type IRepo interface {
	//Database
	GetRepo() *gorm.DB

	//User
	CreateUser(user *model.User) error
	GetAllUser() ([]model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUserByID(id string) (*model.User, error)

	//Token
	CreateToken(token []model.Token) error
}
