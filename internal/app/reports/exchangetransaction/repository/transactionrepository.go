package repository

import (
	"context"
	"time"

	"github.com/joseMarciano/crypto-manager/internal/app/reports/exchangetransaction/domain"
	errorspkg "github.com/joseMarciano/crypto-manager/internal/errors"
	slicespkg "github.com/joseMarciano/crypto-manager/pkg/slices"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return TransactionRepository{db: db}
}

func (r TransactionRepository) FindAllBetween(ctx context.Context, startDate time.Time, endDate time.Time) ([]domain.ExchangeTransaction, error) {
	var models []exchangeTransactionView
	query := r.db.WithContext(ctx).Model(exchangeTransactionView{}).
		Where("executed_at BETWEEN ? AND ?", startDate, endDate).
		Find(&models)
	if err := query.Error; err != nil {
		return nil, errorspkg.NewUnexpectedError("unexpected error creating exchange transaction", err)
	}

	return slicespkg.Map(models, exchangeTransactionView.toDomain), nil
}
