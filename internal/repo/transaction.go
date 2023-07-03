package repo

import (
	"fmt"
	"gorm.io/gorm"
	"wallet/internal/model"
)

type TransactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) TransactionRepo {
	return TransactionRepo{
		db: db,
	}
}

func (r *TransactionRepo) CreateTransaction(transaction *model.Transaction) error {
	if err := r.db.Create(&transaction).Error; err != nil {
		return fmt.Errorf("Failed to save transaction: %v. ", err)
	}
	return nil
}
