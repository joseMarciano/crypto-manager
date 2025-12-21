package domain

import "github.com/joseMarciano/crypto-manager/pkg/rounder"

type Account struct {
	ID         string
	UserID     string
	ExchangeID string
	Balance    float64
}

func (a *Account) Deposit(amount float64) {
	a.Balance = rounder.TwoDecimalPlaces(a.Balance + amount)
}

func (a *Account) WithDraw(amount float64) {
	a.Balance = rounder.TwoDecimalPlaces(a.Balance - amount)
}
