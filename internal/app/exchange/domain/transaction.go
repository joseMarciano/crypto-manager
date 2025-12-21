package domain

import (
	"time"
)

const (
	Deposit    Type = "deposit"
	Withdrawal Type = "withdrawal"
)

type (
	Type                string
	ExchangeTransaction struct {
		ID         string
		ExchangeID string
		Type       Type
		Amount     float64
		ExecutedAt time.Time
	}
)
