package createaccount

import (
	"github.com/joseMarciano/crypto-manager/internal/app/exchange/usecase/createaccount"
	createaccountpb "github.com/joseMarciano/crypto-manager/pkg/proto/exchange/createaccount"
)

func toInputDTO(request *createaccountpb.CreateAccountRequest) createaccount.Input {
	return createaccount.Input{
		UserID:     request.GetUserId(),
		ExchangeID: request.GetExchangeId(),
	}
}

func toResponse(u createaccount.Output) *createaccountpb.CreateAccountResponse {
	return &createaccountpb.CreateAccountResponse{
		Id:         u.ID,
		UserId:     u.UserID,
		ExchangeId: u.ExchangeID,
		Balance:    u.Balance,
	}
}
