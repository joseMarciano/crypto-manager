package repository

import (
	"context"

	"github.com/joseMarciano/crypto-manager/internal/app/exchange/domain"
	errorspkg "github.com/joseMarciano/crypto-manager/internal/errors"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return TransactionRepository{db: db}
}

func (r TransactionRepository) CreateTransaction(ctx context.Context, d domain.ExchangeTransaction) (domain.ExchangeTransaction, error) {
	model := toExchangeTransactionEntity(d)
	if err := r.db.WithContext(ctx).Model(model).Create(model).Error; err != nil {
		return domain.ExchangeTransaction{}, errorspkg.NewUnexpectedError("unexpected error creating exchange transaction", err)
	}

	return model.toDomain(), nil
}
