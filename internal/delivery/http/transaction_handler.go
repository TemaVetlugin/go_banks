package http

import (
	"github.com/gin-gonic/gin"
	"my-go-project/internal/usecase"
	"net/http"
	"strconv"
)

type TransactionHandler struct {
	useCase *usecase.TransactionUseCase
}

func NewTransactionHandler(useCase *usecase.TransactionUseCase) *TransactionHandler {
	return &TransactionHandler{useCase}
}

// Обработчик для отправки средств
func (h *TransactionHandler) Send(c *gin.Context) {
	var request struct {
		From   string  `json:"from"`
		To     string  `json:"to"`
		Amount float64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := h.useCase.SendFunds(request.From, request.To, request.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction successful"})
}

// Обработчик для получения последних N транзакций
func (h *TransactionHandler) GetLast(c *gin.Context) {
	countStr := c.Query("count")
	count, err := strconv.Atoi(countStr)
	if err != nil || count <= 0 {
		count = 10 // По умолчанию 10
	}

	transactions, err := h.useCase.GetLastTransactions(count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	c.JSON(http.StatusOK, transactions)
}
