package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/joseMarciano/crypto-manager/internal/app/exchange/domain"
	errorspkg "github.com/joseMarciano/crypto-manager/internal/errors"
	slicespkg "github.com/joseMarciano/crypto-manager/pkg/slices"

	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return AccountRepository{db: db}
}

func (r AccountRepository) CreateAccount(ctx context.Context, d domain.Account) (domain.Account, error) {
	model := toAccountEntity(d)
	if err := r.db.WithContext(ctx).Model(model).Create(model).Error; err != nil {
		return domain.Account{}, errorspkg.NewUnexpectedError("unexpected error creating account", err)
	}

	return model.toDomain(), nil
}

func (r AccountRepository) ExistsAccountByUserAndExchange(ctx context.Context, userID string, exchangeID string) (bool, error) {
	var exists bool
	query := r.db.WithContext(ctx).Model(accountEntity{}).Where("user_id = ? AND exchange_id = ?", userID, exchangeID).Select("1").Scan(&exists)
	if err := query.Error; err != nil {
		return exists, errorspkg.NewUnexpectedError("unexpected error on existing account by user and exchange", err)
	}

	return exists, nil
}

func (r AccountRepository) FindAccountByID(ctx context.Context, id string) (domain.Account, error) {
	var entity accountEntity
	query := r.db.WithContext(ctx).Model(accountEntity{}).Where("id = ?", id).First(&entity)
	if err := query.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Account{}, errorspkg.NewNotFoundError(fmt.Sprintf("account %s not found", id), err)
		}
		return domain.Account{}, errorspkg.NewUnexpectedError("unexpected error finding account by id", err)
	}

	return entity.toDomain(), nil
}

func (r AccountRepository) UpdateAccount(ctx context.Context, d domain.Account) (domain.Account, error) {
	model := toAccountEntity(d)
	if err := r.db.WithContext(ctx).Model(model).Where("id = ?", model.ID).Updates(model).Error; err != nil {
		return domain.Account{}, errorspkg.NewUnexpectedError("unexpected error updating account", err)
	}

	return model.toDomain(), nil
}

func (r AccountRepository) FindAccountsByUserID(ctx context.Context, id string) ([]domain.Account, error) {
	var entities []accountEntity
	query := r.db.WithContext(ctx).Model(accountEntity{}).Where("user_id = ?", id).Find(&entities)
	if err := query.Error; err != nil {
		return nil, errorspkg.NewUnexpectedError("unexpected error finding accounts by user id", err)
	}

	return slicespkg.Map(entities, accountEntity.toDomain), nil
}
