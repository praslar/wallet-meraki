package repo

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"time"
	"wallet/internal/model"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return UserRepo{db: db}
}

func (rp *UserRepo) CreateToken() string {
	var user *model.User
	var role *model.Role
	err := rp.db.Table("users").First(&user).Error
	if err != nil {
		return ""
	}
	err = rp.db.Table("roles").First(&role).Error
	if err != nil {
		return ""
	}
	// Create a new token object with claims
	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"role":  role.Name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign the token with the secret
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return ""
	}
	return tokenString
}

func (rp *UserRepo) UserLogin(email string, password string) (*model.User, *model.Role, error) {
	var user *model.User
	var role *model.Role
	err := rp.db.Table("users").Where(" email = ?", email).Find(&user).Error
	if err != nil {
		return nil, nil, err
	}
	err = rp.db.Table("roles").Find(&role).Error
	if err != nil {
		return nil, nil, err
	}
	if password != user.Password {
		err := fmt.Errorf("Wrong Password - Try Again. ")
		return nil, nil, err
	}

	if user.UserRole == role.Value {
		fmt.Printf("- Email: %s\n- Role: %s\n", user.Email, role.Name)
	}
	rp.CreateToken()
	return user, role, nil
}
