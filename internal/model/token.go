package model

import "github.com/google/uuid"

type Token struct {
	BaseModel
	Address string  `json:"address" gorm:"primaryKey"`
	Symbol  string  `json:"symbol"`
	Price   float64 `json:"price"`
	Key     string  `json:"key"`
}

type TokenRequest struct {
	ID      string    `json:"id"`
	Address uuid.UUID `json:"address" `
	Symbol  string    `json:"symbol"`
	Price   float64   `json:"price"`
}
