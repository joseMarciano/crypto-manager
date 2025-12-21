package createuser

import (
	"context"

	"github.com/joseMarciano/crypto-manager/internal/app/user/usecase/createuser"
	createuserpb "github.com/joseMarciano/crypto-manager/pkg/proto/user/create"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	createuserpb.UnimplementedCreateUserHandlerServer
	useCase createuser.UseCase
}

func New(useCase createuser.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h Handler) Execute(ctx context.Context, req *createuserpb.CreateUserRequest) (*createuserpb.CreateUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	input, err := toInputDTO(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid birthday format")
	}

	output, err := h.useCase.Execute(ctx, input)
	if err != nil {
		return nil, err
	}

	return toResponse(output), nil
}
