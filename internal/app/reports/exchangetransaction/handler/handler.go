package handler

import (
	"context"

	"github.com/joseMarciano/crypto-manager/internal/app/reports/exchangetransaction/usecase"
	exchangetransactionpb "github.com/joseMarciano/crypto-manager/pkg/proto/report/exchangetransaction"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	exchangetransactionpb.UnimplementedExchangeTransactionHandlerServer
	useCase usecase.UseCase
}

func New(useCase usecase.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h Handler) Execute(ctx context.Context, req *exchangetransactionpb.FetchExchangeTransactionRequest) (*exchangetransactionpb.FetchExchangeTransactionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	output, err := h.useCase.Execute(ctx, toInputDTO(req))
	if err != nil {
		return nil, err
	}

	return toResponse(output), nil
}
