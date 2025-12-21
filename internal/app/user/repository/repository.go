package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/joseMarciano/crypto-manager/internal/app/user/domain"
	errorspkg "github.com/joseMarciano/crypto-manager/internal/errors"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return Repository{db: db}
}

func (r Repository) Create(ctx context.Context, d domain.User) (domain.User, error) {
	model := toEntity(d)
	if err := r.db.WithContext(ctx).Model(model).Create(model).Error; err != nil {
		return domain.User{}, errorspkg.NewUnexpectedError("unexpected error creating user", err)
	}

	return model.toDomain(), nil
}

func (r Repository) ExistsByName(ctx context.Context, name string) (bool, error) {
	var exists bool
	query := r.db.WithContext(ctx).Model(userEntity{}).Where("name = ?", name).Select("1").Scan(&exists)
	if err := query.Error; err != nil {
		return exists, errorspkg.NewUnexpectedError("unexpected error checking existing user by name", err)
	}

	return exists, nil
}

func (r Repository) ExistsByDocument(ctx context.Context, document string) (bool, error) {
	var exists bool
	query := r.db.WithContext(ctx).Model(userEntity{}).Where("document_number = ?", document).Select("1").Scan(&exists)
	if err := query.Error; err != nil {
		return exists, errorspkg.NewUnexpectedError("unexpected error checking existing user by document_number", err)
	}

	return exists, nil
}

func (r Repository) FindUserByID(ctx context.Context, id string) (domain.User, error) {
	var entity userEntity
	query := r.db.WithContext(ctx).Model(userEntity{}).Where("id = ?", id).First(&entity)
	if err := query.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, errorspkg.NewNotFoundError(fmt.Sprintf("user %s not found", id), err)
		}
		return domain.User{}, errorspkg.NewUnexpectedError("unexpected error finding user by id", err)
	}

	return entity.toDomain(), nil
}
