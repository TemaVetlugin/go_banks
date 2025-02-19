package repository

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"my-go-project/internal/domain"
	"time"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{db}
}

func (r *WalletRepository) GetBalance(walletID string) (float64, error) {
	var wallet domain.Wallet
	if err := r.db.First(&wallet, "id = ?", walletID).Error; err != nil {
		return 0, err
	}
	return wallet.Balance, nil
}

func (r *WalletRepository) UpdateBalance(walletAddress string, amount float64) error {
	result := r.db.Model(&domain.Wallet{}).Where("id = ?", walletAddress).Update("balance", gorm.Expr("balance + ?", amount))
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Генерация случайного 64-значного адреса
func generateWalletAddress() string {
	rand.Seed(time.Now().UnixNano())
	randomBytes := []byte(fmt.Sprintf("%d", rand.Int63()))
	hash := sha256.Sum256(randomBytes)
	return hex.EncodeToString(hash[:])
}

func (r *WalletRepository) InitWallets() error {
	var count int64
	r.db.Model(&domain.Wallet{}).Count(&count)

	if count > 0 {
		return nil // Если кошельки уже есть — выходим
	}

	wallets := make([]domain.Wallet, 10)
	for i := range wallets {
		wallets[i] = domain.Wallet{
			ID:      generateWalletAddress(),
			Balance: 100.0,
		}
	}

	return r.db.Create(&wallets).Error
}
