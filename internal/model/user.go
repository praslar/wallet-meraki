package model

import "github.com/google/uuid"

type User struct {
	BaseModel
	ID       uuid.UUID `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Wallets  []Wallet  `json:"wallets"`
	Role     Role      `json:"role" gorm:"foreignKey:role_id"`
	UserRole int       `json:"role_id"`
}
