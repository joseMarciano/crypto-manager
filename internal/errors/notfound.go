package errors

import "errors"

var ErrNotFound = NewNotFoundError("resource not found", nil)

type NotFoundError struct {
	AppError
}

func NewNotFoundError(message string, err error) error {
	return NotFoundError{AppError: AppError{Message: message, Code: ErrNotFoundCode, Cause: err}}
}

func (e NotFoundError) Is(target error) bool {
	return errors.Is(target, ErrNotFound)
}
