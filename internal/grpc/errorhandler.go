package grpc

import (
	"errors"

	errorspkg "github.com/joseMarciano/crypto-manager/internal/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var defaultError = status.Error(codes.Internal, "internal server error")

// ErrorHandler is a middleware that handles errors set in the request context and sends appropriate GRPC responses.
func errorHandler(err error) error {
	if clientErr, ok := errors.AsType[errorspkg.ClientError](err); ok {
		return status.Error(codes.InvalidArgument, clientErr.Error())
	}

	if notFoundErr, ok := errors.AsType[errorspkg.NotFoundError](err); ok {
		return status.Error(codes.NotFound, notFoundErr.Error())
	}

	if businessErr, ok := errors.AsType[errorspkg.BusinessValidationError](err); ok {
		return status.Error(codes.InvalidArgument, businessErr.Error())
	}

	if unexpectedErr, ok := errors.AsType[errorspkg.UnexpectedError](err); ok {
		return status.Error(codes.Internal, unexpectedErr.Error())
	}

	return defaultError
}
