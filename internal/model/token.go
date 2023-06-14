package model

import "github.com/google/uuid"

type Token struct {
	BaseModel
	Address uuid.UUID `json:"address" gorm:"primaryKey;default:uuid_generate_v4()"`
	Symbol  string    `json:"symbol"`
	Price   float64   `json:"price"`
}

type TokenRequest struct {
	ID      string    `json:"id"`
	Address uuid.UUID `json:"address" `
	Symbol  string    `json:"symbol"`
	Price   float64   `json:"price"`
}
