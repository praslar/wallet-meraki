package model

import "github.com/google/uuid"

type User struct {
	BaseModel
	ID       uuid.UUID `json:"id" gorm:"primaryKey;unique;default:uuid_generate_v4()"`
	Email    string    `json:"email"`
	Password string    `json:"password"`

	Wallets []Wallet `json:"wallets" gorm:"foreignKey:UserID"`

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
	WalletID uuid.UUID `json:"wallet_id" gorm:"primaryKey;default:uuid_generate_v4()"`
	Address  uuid.UUID `json:"address" gorm:"default:uuid_generate_v4()"`
	Name     string    `json:"name"`

	UserID uuid.UUID `json:"user_id"`

	//Tokens []TokenUser `json:"tokens" gorm:"many2many:wallet_tokens;"`
}

//type TokenUser struct {
//	BaseModel
//	TokenID uuid.UUID `json:"token_id" gorm:"primaryKey;default:uuid_generate_v4()"`
//	Symbol  string    `json:"symbol"`
//	Amount  float64   `json:"amount"`
//}

type Token struct {
	BaseModel
	Address     uuid.UUID `json:"address" gorm:"primaryKey;default:uuid_generate_v4()"`
	Symbol      string    `json:"symbol"`
	TotalSupply uint64    `json:"total_supply"`
}

// Which creates join table: wallet_tokens
//   foreign key: wallet_id, reference: users.id
//   foreign key: token_id, reference: token.id

//type Transaction struct {
//	BaseModel
//	ID uuid.UUID `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`
//
//	SenderWalletAddress uuid.UUID `json:"sender_wallet_address"`
//	SenderWallet        Wallet    `json:"sender_wallet" gorm:"foreignKey:SenderWalletAddress"`
//
//	ReceiverWalletAddress uuid.UUID `json:"receiver_wallet_address"`
//	ReceiverWallet        Wallet    `json:"receiver_wallet" gorm:"foreignKey:ReceiverWalletAddress"`
//
//	TokenID uuid.UUID `json:"token_id"`
//	Token   Token     `json:"token" gorm:"foreignKey:token_id;references:token_id"`
//	Amount  float64   `json:"amount"`
//}

type TransactionRequest struct {
	SenderWalletAddress   uuid.UUID `json:"sender_wallet_address"`
	ReceiverWalletAddress uuid.UUID `json:"receiver_wallet_address"`
}

type WalletRequest struct {
	Address string    `json:"address"`
	Name    string    `json:"name"`
	UserID  uuid.UUID `json:"user_id"`
}

type TokenRequest struct {
	Address     uuid.UUID `json:"address"`
	Symbol      string    `json:"symbol"`
	TotalSupply uint64    `json:"total_supply"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
