package createexchange

import (
	"context"

	"github.com/joseMarciano/crypto-manager/internal/app/exchange/usecase/createexchange"
	createexchangepb "github.com/joseMarciano/crypto-manager/pkg/proto/exchange/create"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	createexchangepb.UnimplementedCreateExchangeHandlerServer
	useCase createexchange.UseCase
}

func New(useCase createexchange.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h Handler) Execute(ctx context.Context, req *createexchangepb.CreateExchangeRequest) (*createexchangepb.CreateExchangeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	output, err := h.useCase.Execute(ctx, toInputDTO(req))
	if err != nil {
		return nil, err
	}

	return toResponse(output), nil
}
