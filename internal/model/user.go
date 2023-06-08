package model

import "github.com/google/uuid"

type Role struct {
	BaseModel
	ID    uuid.UUID `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`
	Name  string    `json:"name"`
	Value uint8     `json:"value"`
	Key   string    `json:"key"`
}

type User struct {
	BaseModel
	ID       uuid.UUID `json:"id" gorm:"primaryKey;unique;default:uuid_generate_v4()"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Wallets  []Wallet  `json:"wallets" gorm:"foreignKey:user_id;references:id"`
	RoleID   uuid.UUID `json:"role_id"`
	Role     Role      `json:"role" gorm:"foreignKey:role_id;references:id"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
