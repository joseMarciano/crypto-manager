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
	var clientErr errorspkg.ClientError
	if errors.As(err, &clientErr) {
		return status.Error(codes.InvalidArgument, clientErr.Error())
	}

	var notFoundErr errorspkg.NotFoundError
	if errors.As(err, &notFoundErr) {
		return status.Error(codes.NotFound, notFoundErr.Error())
	}

	var businessErr errorspkg.BusinessValidationError
	if errors.As(err, &businessErr) {
		return status.Error(codes.InvalidArgument, businessErr.Error())
	}

	var unexpectedErr errorspkg.UnexpectedError
	if errors.As(err, &unexpectedErr) {
		return status.Error(codes.Internal, unexpectedErr.Error())
	}

	return defaultError
}
