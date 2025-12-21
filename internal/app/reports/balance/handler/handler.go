package handler

import (
	"context"

	"github.com/joseMarciano/crypto-manager/internal/app/reports/balance/usecase"
	balancepb "github.com/joseMarciano/crypto-manager/pkg/proto/report/balance"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	balancepb.UnimplementedBalanceHandlerServer
	useCase usecase.UseCase
}

func New(useCase usecase.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h Handler) Execute(ctx context.Context, req *balancepb.FetchBalanceRequest) (*balancepb.FetchBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	output, err := h.useCase.Execute(ctx, toInputDTO(req))
	if err != nil {
		return nil, err
	}

	return toResponse(output), nil
}
