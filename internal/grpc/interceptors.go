package grpc

import (
	"context"

	"google.golang.org/grpc"
)

func Interceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		res, err := handler(ctx, req)
		if err != nil {
			return nil, errorHandler(err)
		}

		return res, nil
	}
}
