package repository

import (
	"github.com/joseMarciano/crypto-manager/internal/app/exchange/domain"
)

type exchangeEntity struct {
	ID                    string  `gorm:"column:id;primaryKey"`
	Name                  string  `gorm:"column:name"`
	MinimumAge            int     `gorm:"column:minimum_age"`
	MaximumTransferAmount float64 `gorm:"column:maximum_transfer_amount"`
}

func (u exchangeEntity) toDomain() domain.Exchange {
	return domain.Exchange(u)
}

func (u exchangeEntity) TableName() string {
	return "exchanges"
}

func toExchangeEntity(u domain.Exchange) exchangeEntity {
	return exchangeEntity(u)
}
