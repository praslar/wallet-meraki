package model

import "github.com/google/uuid"

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
	SenderWalletAddress   uuid.UUID `json:"sender_wallet_address"`
	ReceiverWalletAddress uuid.UUID `json:"receiver_wallet_address"`
	Amount                float64   `json:"amount"`
}

type AirdropTransactionRequest struct {
	TokenAddress          uuid.UUID `json:"token_address" `
	SenderWalletAddress   string    `json:"sender_wallet_address"`
	ReceiverWalletAddress string    `json:"receiver_wallet_address"`
	Amount                float64   `json:"amount"`
}
