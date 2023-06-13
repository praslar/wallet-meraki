package model

import "github.com/google/uuid"

type User struct {
	BaseModel

	ID       uuid.UUID `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Wallets  []Wallet  `json:"wallets"`

	RoleID uuid.UUID `json:"role_id"`
	Role   Role      `json:"role" gorm:"foreignKey:role_id;references:id"`
}

type Role struct {
	BaseModel
	ID    uuid.UUID `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`
	Name  string    `json:"name"`
	Value uint8     `json:"value"`
	Key   string    `json:"key"`
}

type Token struct {
	BaseModel
	Address uuid.UUID `json:"address" gorm:"primaryKey;default:uuid_generate_v4()"`
	Symbol  string    `json:"symbol"`
	Price   float64   `json:"price"`
}

type Wallet struct {
	BaseModel
	Address string    `json:"address" gorm:"primaryKey;default:uuid_generate_v4()"`
	Name    string    `json:"name"`
	UserID  uuid.UUID `json:"user_id" gorm:"column:user_id"`
	User    User      `json:"user"`
}

type Transaction struct {
	BaseModel
	ID uuid.UUID `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`

	FromAddress uuid.UUID `json:"from_address"`
	WalletFrom  Wallet    `json:"wallet_from" gorm:"foreignKey:from_address;references:address"`

	ToAddress uuid.UUID `json:"to_address"`
	WalletTo  Wallet    `json:"wallet_to" gorm:"foreignKey:to_address;references:address"`

	TokenAddress uuid.UUID `json:"token_address"`
	Token        Token     `json:"token" gorm:"foreignKey:token_address;references:address"`
	Amount       float64   `json:"amount"`
}

type TransactionRequest struct {
	TokenAddress          uuid.UUID `json:"token_address" `
	TokenSymbol           string    `json:"token_symbol"`
	TokenPrice            float64   `json:"token_price"`
	SenderWalletAddress   uuid.UUID `json:"sender_wallet_address"`
	ReceiverWalletAddress uuid.UUID `json:"receiver_wallet_address"`
	Amount                float64   `json:"amount"`
}

type WalletRequest struct {
	Address string    `json:"address"`
	Name    string    `json:"name"`
	UserID  uuid.UUID `json:"user_id"`
}

type TokenRequest struct {
	Address uuid.UUID `json:"address" `
	Symbol  string    `json:"symbol"`
	Price   float64   `json:"price"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
