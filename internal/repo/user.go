package repo

import (
	"fmt"
	"gorm.io/gorm"
	"wallet/internal/model"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return UserRepo{db: db}
}

func (rp *UserRepo) UserLogin(email string, password string) (*model.User, error) {
	user := &model.User{}
	rp.db.Where("email = ?", email).Find(&user)
	rp.db.Where("password = ?", password).Find(&user)
	if password != user.Password {
		err := fmt.Errorf("Wrong Password - Try Again. ")
		return nil, err
	}
	fmt.Println("Login successful")
	fmt.Printf("\n User found: \n- Email: %s \n- Role: %v \n", user.Email, user.Roles)
	return user, nil
}
