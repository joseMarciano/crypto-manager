package validator

import (
	"fmt"
	"strings"

	validatorpkg "github.com/go-playground/validator/v10"
)

func TwoDecimalPlaces(fl validatorpkg.FieldLevel) bool {
	value := fl.Field().Float()
	str := fmt.Sprintf("%.10f", value)
	parts := strings.Split(str, ".")
	if len(parts) != 2 {
		return false
	}

	decimals := strings.TrimRight(parts[1], "0")
	return len(decimals) <= 2
}
