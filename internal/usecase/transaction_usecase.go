package usecase

import (
	"errors"
	"my-go-project/internal/domain"
	"my-go-project/internal/repository"
)

type TransactionUseCase struct {
	transactionRepo *repository.TransactionRepository
	walletRepo      *repository.WalletRepository
}

func NewTransactionUseCase(tr *repository.TransactionRepository, wr *repository.WalletRepository) *TransactionUseCase {
	return &TransactionUseCase{tr, wr}
}

// Отправка средств
func (uc *TransactionUseCase) SendFunds(from, to string, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	// Проверяем баланс отправителя
	fromBalance, err := uc.walletRepo.GetBalance(from)
	if err != nil {
		return errors.New("sender wallet not found")
	}

	if fromBalance < amount {
		return errors.New("insufficient funds")
	}

	// Создаём транзакцию
	tx := &domain.Transaction{
		From:   from,
		To:     to,
		Amount: amount,
	}

	// Обновляем баланс кошельков
	err = uc.walletRepo.UpdateBalance(from, -amount)
	if err != nil {
		return err
	}

	err = uc.walletRepo.UpdateBalance(to, amount)
	if err != nil {
		return err
	}

	// Сохраняем транзакцию
	return uc.transactionRepo.CreateTransaction(tx)
}

// Получение последних N транзакций
func (uc *TransactionUseCase) GetLastTransactions(count int) ([]domain.Transaction, error) {
	return uc.transactionRepo.GetLastTransactions(count)
}
