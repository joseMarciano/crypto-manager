package errors

type UnexpectedError struct {
	AppError
}

func NewUnexpectedError(message string, err error) error {
	return UnexpectedError{AppError: AppError{Message: message, Code: ErrUnexpectedCode, Cause: err}}
}
