package main

import (
	"go.uber.org/fx"
	"my-go-project/internal/bootstrap"
)

func main() {
	app := fx.New(
		bootstrap.Module, // Подключаем модуль с зависимостями
	)

	app.Run() // Запускаем приложение
}
