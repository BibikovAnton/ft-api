package scors

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
}

func NewAccount(name, typee string, balance float64, currency string) *Account {
	return &Account{
		Name:     name,
		Type:     typee,
		Balance:  balance,
		Currency: currency,
	}
}
