package model

import "github.com/google/uuid"

type Token struct {
	BaseModel
	ID       uuid.UUID `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`
	WalletID uuid.UUID `json:"wallet_id"`
	Address  string    `json:"address"`
	Symbol   string    `json:"symbol"`
	Amount   float64   `json:"amount"`
}
