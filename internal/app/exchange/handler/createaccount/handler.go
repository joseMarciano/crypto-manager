package createaccount

import (
	"context"

	"github.com/joseMarciano/crypto-manager/internal/app/exchange/usecase/createaccount"
	createaccountpb "github.com/joseMarciano/crypto-manager/pkg/proto/exchange/createaccount"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	createaccountpb.UnimplementedCreateAccountHandlerServer
	useCase createaccount.UseCase
}

func New(useCase createaccount.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h Handler) Execute(ctx context.Context, req *createaccountpb.CreateAccountRequest) (*createaccountpb.CreateAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	output, err := h.useCase.Execute(ctx, toInputDTO(req))
	if err != nil {
		return nil, err
	}

	return toResponse(output), nil
}
