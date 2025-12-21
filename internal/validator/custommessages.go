package validator

import (
	"fmt"

	validatorpkg "github.com/go-playground/validator/v10"
)

func required(field validatorpkg.FieldError) string {
	return fmt.Sprintf("%s is required field", field.Field())
}

func gte(field validatorpkg.FieldError) string {
	return fmt.Sprintf("%s should be greater than or equal to %s", field.Field(), field.Param())
}

func gt(field validatorpkg.FieldError) string {
	return fmt.Sprintf("%s should be greater than %s", field.Field(), field.Param())

}

func twoDecimals(field validatorpkg.FieldError) string {
	return fmt.Sprintf("%s should have at most two decimal places", field.Field())
}
