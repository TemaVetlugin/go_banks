package bootstrap

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"log"
	"my-go-project/config"
	"my-go-project/internal/delivery/http"
	"my-go-project/internal/repository"
	"my-go-project/internal/usecase"
)

// Функция инициализации кошельков
func initWallets(walletRepo *repository.WalletRepository) {
	if err := walletRepo.InitWallets(); err != nil {
		log.Fatalf("Ошибка инициализации кошельков: %v", err)
	}
}

// FX-модуль для автоматической загрузки зависимостей
var Module = fx.Options(
	// Конфигурация и БД
	fx.Provide(config.LoadConfig),
	fx.Provide(config.InitDB),

	// Репозитории
	fx.Provide(repository.NewWalletRepository),
	fx.Provide(repository.NewTransactionRepository),

	// Use Cases (бизнес-логика)
	fx.Provide(usecase.NewWalletUseCase),
	fx.Provide(usecase.NewTransactionUseCase),

	// HTTP-хендлеры
	fx.Provide(http.NewWalletHandler),
	fx.Provide(http.NewTransactionHandler),

	// Маршрутизация
	fx.Provide(NewRouter),

	// Запуск приложения
	fx.Invoke(StartServer),

	fx.Invoke(config.MigrateDB),

	// Инициализация кошельков
	fx.Invoke(initWallets),
)

// Функция инициализации маршрутов
func NewRouter(
	walletHandler *http.WalletHandler,
	transactionHandler *http.TransactionHandler,
) *gin.Engine {
	r := gin.Default()
	http.SetupRoutes(r, walletHandler, transactionHandler)
	return r
}

// Функция запуска сервера
func StartServer(lc fx.Lifecycle, r *gin.Engine) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := r.Run(":8080"); err != nil {
					panic(err) // Обработка ошибки, если сервер не запускается
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
