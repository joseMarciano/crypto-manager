package repository

import "github.com/joseMarciano/crypto-manager/internal/app/exchange/domain"

type accountEntity struct {
	ID         string  `gorm:"column:id;primaryKey"`
	UserID     string  `gorm:"column:user_id"`
	ExchangeID string  `gorm:"column:exchange_id"`
	Balance    float64 `gorm:"column:balance"`
}

func (u accountEntity) toDomain() domain.Account {
	return domain.Account(u)
}

func (u accountEntity) TableName() string {
	return "accounts"
}

func toAccountEntity(u domain.Account) accountEntity {
	return accountEntity(u)
}
