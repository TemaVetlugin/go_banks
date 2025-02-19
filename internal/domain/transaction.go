package domain

import "time"

type Transaction struct {
	ID        uint `gorm:"primaryKey"`
	From      string
	To        string
	Amount    float64
	CreatedAt time.Time
}
