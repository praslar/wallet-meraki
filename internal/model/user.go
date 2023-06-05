package model

import "github.com/google/uuid"

type User struct {
	BaseModel
	ID          uuid.UUID `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Wallets     []Wallet  `json:"wallets"`
	RoleID      uuid.UUID `json:"roleID" gorm:"default:uuid"`
	Role        Role      `json:"role"`
	DefaultRole string    `json:"defaultRole"`
}
