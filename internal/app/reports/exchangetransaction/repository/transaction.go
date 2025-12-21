package repository

import (
	"github.com/joseMarciano/crypto-manager/internal/app/reports/exchangetransaction/domain"
)

type exchangeTransactionView struct {
	ExchangeID string  `gorm:"column:exchange_id"`
	Type       string  `gorm:"column:type"`
	Amount     float64 `gorm:"column:amount"`
	ExecutedAt string  `gorm:"column:executed_at"`
}

func (u exchangeTransactionView) toDomain() domain.ExchangeTransaction {
	return domain.ExchangeTransaction{
		ExchangeID: u.ExchangeID,
		Type:       u.Type,
		Amount:     u.Amount,
		ExecutedAt: u.ExecutedAt,
	}
}

func (u exchangeTransactionView) TableName() string {
	return "exchange_transactions_view"
}
