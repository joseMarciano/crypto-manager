package deposit

import (
	"github.com/joseMarciano/crypto-manager/internal/app/exchange/usecase/deposit"
	depositpb "github.com/joseMarciano/crypto-manager/pkg/proto/exchange/deposit"
)

func toInputDTO(request *depositpb.DepositRequest) deposit.Input {
	return deposit.Input{
		ExchangeID: request.GetExchangeId(),
		AccountID:  request.GetAccountId(),
		Amount:     request.GetAmount(),
	}
}

func toResponse(u deposit.Output) *depositpb.DepositResponse {
	return &depositpb.DepositResponse{
		AccountId: u.AccountID,
		Balance:   u.Balance,
	}
}
