package grpc

import (
	"context"

	"github.com/joseMarciano/crypto-manager/internal/keygenerator"
	"github.com/joseMarciano/crypto-manager/internal/observability"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func Interceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		ctx = observability.SetRequestID(ctx, extractRequestID(ctx))
		res, err := handler(ctx, req)
		if err != nil {
			return nil, errorHandler(err)
		}

		return res, nil
	}
}

func extractRequestID(ctx context.Context) string {
	incomingMetadata, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return keygenerator.Generate()
	}

	for _, id := range incomingMetadata.Get(observability.RequestIDMetadataKey) {
		if id != "" {
			return id
		}
	}

	return keygenerator.Generate()
}
