package model

import "github.com/google/uuid"

type User struct {
	BaseModel
	ID       uuid.UUID `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Roles    uint8     `json:"roles"  gorm:""`
	Wallets  []Wallet  `json:"wallets"`
}

type Role struct {
	BaseModel
	Name  string `json:"name"`
	Value uint8  `json:"value" gorm:"primaryKey"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
