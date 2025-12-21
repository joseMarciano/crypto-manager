package errors

type BusinessValidationError struct {
	AppError
}

func NewBusinessValidationError(message string, cause error) error {
	return BusinessValidationError{AppError: AppError{Message: message, Code: ErrBusinessValidationCode, Cause: cause}}
}
