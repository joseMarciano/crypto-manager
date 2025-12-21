package deposit

import (
	"context"

	"github.com/joseMarciano/crypto-manager/internal/app/exchange/usecase/deposit"
	depositpb "github.com/joseMarciano/crypto-manager/pkg/proto/exchange/deposit"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	depositpb.UnimplementedDepositHandlerServer
	useCase deposit.UseCase
}

func New(useCase deposit.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h Handler) Execute(ctx context.Context, req *depositpb.DepositRequest) (*depositpb.DepositResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	output, err := h.useCase.Execute(ctx, toInputDTO(req))
	if err != nil {
		return nil, err
	}

	return toResponse(output), nil
}
