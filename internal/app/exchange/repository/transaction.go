package repository

import (
	"time"

	"github.com/joseMarciano/crypto-manager/internal/app/exchange/domain"
)

type exchangeTransactionEntity struct {
	ID         string    `gorm:"column:id;primaryKey"`
	ExchangeID string    `gorm:"column:exchange_id"`
	Type       string    `gorm:"column:type"`
	Amount     float64   `gorm:"column:amount"`
	ExecutedAt time.Time `gorm:"column:executed_at"`
}

func (u exchangeTransactionEntity) toDomain() domain.ExchangeTransaction {
	return domain.ExchangeTransaction{
		ID:         u.ID,
		ExchangeID: u.ExchangeID,
		Type:       domain.Type(u.Type),
		Amount:     u.Amount,
		ExecutedAt: u.ExecutedAt,
	}
}

func (u exchangeTransactionEntity) TableName() string {
	return "exchange_transactions"
}

func toExchangeTransactionEntity(d domain.ExchangeTransaction) exchangeTransactionEntity {
	return exchangeTransactionEntity{
		ID:         d.ID,
		ExchangeID: d.ExchangeID,
		Type:       string(d.Type),
		Amount:     d.Amount,
		ExecutedAt: d.ExecutedAt,
	}
}
