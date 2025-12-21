package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/joseMarciano/crypto-manager/internal/app/exchange/domain"
	errorspkg "github.com/joseMarciano/crypto-manager/internal/errors"

	"gorm.io/gorm"
)

type ExchangeRepository struct {
	db *gorm.DB
}

func NewExchangeRepository(db *gorm.DB) ExchangeRepository {
	return ExchangeRepository{db: db}
}

func (r ExchangeRepository) CreateExchange(ctx context.Context, d domain.Exchange) (domain.Exchange, error) {
	model := toExchangeEntity(d)
	if err := r.db.WithContext(ctx).Model(model).Create(model).Error; err != nil {
		return domain.Exchange{}, errorspkg.NewUnexpectedError("unexpected error creating exchange", err)
	}

	return model.toDomain(), nil
}

func (r ExchangeRepository) FindExchangeByID(ctx context.Context, id string) (domain.Exchange, error) {
	var entity exchangeEntity
	query := r.db.WithContext(ctx).Model(exchangeEntity{}).Where("id = ?", id).First(&entity)
	if err := query.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Exchange{}, errorspkg.NewNotFoundError(fmt.Sprintf("exchange %s not found", id), err)
		}
		return domain.Exchange{}, errorspkg.NewUnexpectedError("unexpected error finding exchange by id", err)
	}

	return entity.toDomain(), nil
}
