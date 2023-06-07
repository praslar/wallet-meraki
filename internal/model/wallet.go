package model

import "github.com/google/uuid"

type Wallet struct {
	BaseModel
	ID           uuid.UUID  `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`
	UserID       uuid.UUID  `json:"user_id" gorm:"default:uuid_generate_v4()"`
	Address      string     `json:"address"`
	Name         string     `json:"name"`
	Tokens       []Token    `json:"token" gorm:"foreignKey:wallet_id;references:id"`
	WalletTypeID uuid.UUID  `json:"wallet_type_id" gorm:"default:uuid_generate_v4()"`
	WalletType   WalletType `json:"type" gorm:"foreignKey:wallet_type_id;references:id"`
}

type WalletType struct {
	BaseModel
	ID           uuid.UUID `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`
	WalletTypeID uuid.UUID `json:"wallet_type_id"`
	Name         int8      `json:"name"`
	Key          string    `json:"key"`
}
