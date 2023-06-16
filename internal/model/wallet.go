package model

import "github.com/google/uuid"

type Wallet struct {
	BaseModel
	Address uuid.UUID `json:"address" gorm:"primaryKey;default:uuid_generate_v4()"`
	Name    string    `json:"name"`
	UserID  uuid.UUID `json:"user_id"`
	User    User      `json:"user" gorm:"foreignKey:user_id;references:id"`
}
type WalletRequest struct {
	Name   string    `json:"name"`
	UserID uuid.UUID `json:"user_id"`
}
