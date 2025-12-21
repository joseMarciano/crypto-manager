package validator

import (
	"context"
	"sync"

	errorspkg "github.com/joseMarciano/crypto-manager/internal/errors"

	validatorpkg "github.com/go-playground/validator/v10"
)

var (
	once      sync.Once
	validator *validatorpkg.Validate
)

func BusinessValidate[T any](ctx context.Context, target T) error {
	var err error
	once.Do(func() { err = register() })

	if err != nil {
		return errorspkg.NewUnexpectedError("validation not registered", err)
	}

	if err = validator.StructCtx(ctx, &target); err != nil {
		return errorspkg.NewBusinessValidationError(translateBusinessError(err), nil)
	}

	return nil
}

func register() error {
	validator = validatorpkg.New(validatorpkg.WithRequiredStructEnabled())
	return validator.RegisterValidation("two-decimals", TwoDecimalPlaces)
}
