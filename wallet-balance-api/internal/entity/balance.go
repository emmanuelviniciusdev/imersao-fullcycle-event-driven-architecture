package entity

import (
	"time"
)

type Balance struct {
	ID        int
	AccountID string
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewBalance(accountID string, balance float64) *Balance {
	return &Balance{
		AccountID: accountID,
		Balance:   balance,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
