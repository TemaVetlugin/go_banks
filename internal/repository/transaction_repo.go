package repository

import (
	"gorm.io/gorm"
	"my-go-project/internal/domain"
	"time"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db}
}

// Создание новой транзакции
func (r *TransactionRepository) CreateTransaction(tx *domain.Transaction) error {
	tx.CreatedAt = time.Now()
	return r.db.Create(tx).Error
}

// Получение последних N транзакций
func (r *TransactionRepository) GetLastTransactions(count int) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	err := r.db.Order("created_at desc").Limit(count).Find(&transactions).Error
	return transactions, err
}
