package createuser

import (
	"github.com/joseMarciano/crypto-manager/internal/app/user/usecase/createuser"
	createuserpb "github.com/joseMarciano/crypto-manager/pkg/proto/user/create"
	timepkg "github.com/joseMarciano/crypto-manager/pkg/time"
)

func toInputDTO(request *createuserpb.CreateUserRequest) (createuser.Input, error) {
	birthday, err := timepkg.ParseCanonical(request.GetBirthday())
	if err != nil {
		return createuser.Input{}, err
	}

	return createuser.Input{
		Name:           request.GetName(),
		Birthday:       birthday,
		DocumentNumber: request.GetDocumentNumber(),
	}, nil
}

func toResponse(u createuser.Output) *createuserpb.CreateUserResponse {
	return &createuserpb.CreateUserResponse{
		Id:             u.ID,
		Name:           u.Name,
		Birthday:       timepkg.FormatCanonical(u.Birthday),
		DocumentNumber: u.DocumentNumber,
	}
}
