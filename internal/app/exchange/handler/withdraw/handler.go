package withdraw

import (
	"context"

	"github.com/joseMarciano/crypto-manager/internal/app/exchange/usecase/withdraw"
	withdrawpb "github.com/joseMarciano/crypto-manager/pkg/proto/exchange/withdraw"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	withdrawpb.UnimplementedWithdrawHandlerServer
	useCase withdraw.UseCase
}

func New(useCase withdraw.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h Handler) Execute(ctx context.Context, req *withdrawpb.WithdrawRequest) (*withdrawpb.WithdrawResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	output, err := h.useCase.Execute(ctx, toInputDTO(req))
	if err != nil {
		return nil, err
	}

	return toResponse(output), nil
}
