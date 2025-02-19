package domain

type Wallet struct {
	ID      string `gorm:"primaryKey"`
	Balance float64
}
