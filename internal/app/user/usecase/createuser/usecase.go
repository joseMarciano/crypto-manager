package createuser

import (
	"context"
	"fmt"

	"github.com/joseMarciano/crypto-manager/internal/app/user/domain"
	errorspkg "github.com/joseMarciano/crypto-manager/internal/errors"
	"github.com/joseMarciano/crypto-manager/internal/validator"
)

//go:generate mockery --all --case=snake --with-expecter

type (
	Creator interface {
		Create(context.Context, domain.User) (domain.User, error)
	}

	Finder interface {
		ExistsByName(context.Context, string) (bool, error)
		ExistsByDocument(context.Context, string) (bool, error)
	}

	UseCase struct {
		creator Creator
		finder  Finder
	}
)

func New(creator Creator, finder Finder) UseCase {
	return UseCase{creator: creator, finder: finder}
}

func (uc UseCase) Execute(ctx context.Context, input Input) (Output, error) {
	if err := validator.BusinessValidate(ctx, input); err != nil {
		return Output{}, err
	}

	user := input.toDomain()

	exists, err := uc.finder.ExistsByName(ctx, user.Name)
	if err != nil {
		return Output{}, err
	}

	if exists {
		return Output{}, errorspkg.NewBusinessValidationError(fmt.Sprintf("user %s already exists", user.Name), nil)
	}

	if exists, err = uc.finder.ExistsByDocument(ctx, user.DocumentNumber); err != nil {
		return Output{}, err
	}

	if exists {
		return Output{}, errorspkg.NewBusinessValidationError(fmt.Sprintf("user with document %s already exists", user.DocumentNumber), nil)
	}

	savedUser, err := uc.creator.Create(ctx, user)
	if err != nil {
		return Output{}, err
	}

	return toOutput(savedUser), nil
}
