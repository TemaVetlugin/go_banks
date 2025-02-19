package usecase

import "my-go-project/internal/repository"

type WalletUseCase struct {
	walletRepo *repository.WalletRepository
}

func NewWalletUseCase(walletRepo *repository.WalletRepository) *WalletUseCase {
	return &WalletUseCase{walletRepo}
}

func (uc *WalletUseCase) GetWalletBalance(walletID string) (float64, error) {
	return uc.walletRepo.GetBalance(walletID)
}
