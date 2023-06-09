package model

import "github.com/google/uuid"

type User struct {
	BaseModel
	ID       uuid.UUID `json:"id" gorm:"primaryKey;unique;default:uuid_generate_v4()"`
	Email    string    `json:"email"`
	Password string    `json:"password"`

	Wallets []Wallet `json:"wallets"`

	RoleID uuid.UUID `json:"role_id;default:uuid_generate_v4()"`
	Role   Role      `json:"role" gorm:"foreignKey:role_id;references:id"`
}

type Role struct {
	BaseModel
	ID    uuid.UUID `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`
	Name  string    `json:"name"`
	Value uint8     `json:"value"`
	Key   string    `json:"key"`
}

type Wallet struct {
	BaseModel
	Address uuid.UUID `json:"address" gorm:"primaryKey;default:uuid_generate_v4()"`
	Name    string    `json:"name"`

	UserID uuid.UUID `json:"user_id"`
	User   User      `json:"user" gorm:"foreignKey:user_id;references:id"`

	Tokens []Token `json:"tokens"`
}

type Token struct {
	BaseModel
	TokenID uuid.UUID `json:"token_id" gorm:"primaryKey;default:uuid_generate_v4()"`
	Symbol  string    `json:"symbol"`

	WalletAddress uuid.UUID `json:"wallet_address"`
	Wallet        Wallet    `json:"wallet" gorm:"foreignKey:wallet_address;references:address"`
}

type WalletRequest struct {
	Address string    `json:"address"`
	Name    string    `json:"name"`
	UserID  uuid.UUID `json:"user_id"`
}

type TokenRequest struct {
	Symbol        string    `json:"symbol"`
	WalletAddress uuid.UUID `json:"wallet_address"`
	TokenID       uuid.UUID `json:"token_id"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
