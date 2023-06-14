package model

type Token struct {
	BaseModel
	Address string  `json:"address" gorm:"primaryKey"`
	Symbol  string  `json:"symbol"`
	Amount  float64 `json:"amount"`
}
