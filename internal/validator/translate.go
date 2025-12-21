package validator

import (
	"errors"

	validatorpkg "github.com/go-playground/validator/v10"
)

var businessErrorTagMap = map[string]func(validatorpkg.FieldError) string{
	"required": required, "gte": gte, "gt": gt, "two-decimals": twoDecimals,
}

func translateBusinessError(err error) string {
	var validationErrors validatorpkg.ValidationErrors
	if !errors.As(err, &validationErrors) {
		return err.Error()
	}

	validationError := validationErrors[0]
	if msgTemplate, exists := businessErrorTagMap[validationError.Tag()]; exists {
		return msgTemplate(validationError)
	}

	return validationError.Error()
}
