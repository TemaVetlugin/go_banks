package http

import (
	"github.com/gin-gonic/gin"
	"my-go-project/internal/usecase"
	"net/http"
)

type WalletHandler struct {
	useCase *usecase.WalletUseCase
}

func NewWalletHandler(useCase *usecase.WalletUseCase) *WalletHandler {
	return &WalletHandler{useCase}
}

func (h *WalletHandler) GetBalance(c *gin.Context) {
	walletID := c.Param("address")
	balance, err := h.useCase.GetWalletBalance(walletID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": balance})
}
