package errors

type ClientError struct {
	AppError
}

func NewClientError(message string, err error) error {
	return ClientError{AppError: AppError{Message: message, Code: ErrClientCode, Cause: err}}
}
