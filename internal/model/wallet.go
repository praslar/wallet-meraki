package model

import "github.com/google/uuid"

type Wallet struct {
	BaseModel
	Address string    `json:"address" gorm:"primaryKey;default:uuid_generate_v4()"`
	Name    string    `json:"name"`
	UserID  uuid.UUID `json:"user_id" gorm:"column:user_id"`
	User    User      `json:"user"`
}

type WalletRequest struct {
	Name   string    `json:"name"`
	UserID uuid.UUID `json:"user_id"`
}
