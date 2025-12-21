package errors

import (
	"fmt"
)

const (
	ErrUnexpectedCode         ErrorCode = "UNEXPECTED"
	ErrNotFoundCode           ErrorCode = "NOT_FOUND"
	ErrBusinessValidationCode ErrorCode = "BUSINESS"
	ErrClientCode             ErrorCode = "CLIENT"
)

type (
	ErrorCode string

	AppError struct {
		Code    ErrorCode
		Message string
		Cause   error
	}
)

func (e AppError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %s | cause: %v", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}
