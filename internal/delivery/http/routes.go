package http

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, walletHandler *WalletHandler, transactionHandler *TransactionHandler) {
	r.GET("/api/wallet/:address/balance", walletHandler.GetBalance)
	r.POST("/api/send", transactionHandler.Send)
	r.GET("/api/transactions", transactionHandler.GetLast)
}
