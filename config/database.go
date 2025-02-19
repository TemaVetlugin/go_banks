package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"my-go-project/internal/domain"
)

func InitDB(cfg *Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return db
}

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(&domain.Wallet{}, &domain.Transaction{})
	if err != nil {
		log.Fatalf("Ошибка миграции базы данных: %v", err)
	}
}
