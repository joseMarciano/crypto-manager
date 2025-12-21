package withdraw

import (
	"github.com/joseMarciano/crypto-manager/internal/app/exchange/usecase/withdraw"
	withdrawpb "github.com/joseMarciano/crypto-manager/pkg/proto/exchange/withdraw"
)

func toInputDTO(request *withdrawpb.WithdrawRequest) withdraw.Input {
	return withdraw.Input{
		ExchangeID: request.GetExchangeId(),
		AccountID:  request.GetAccountId(),
		Amount:     request.GetAmount(),
	}
}

func toResponse(u withdraw.Output) *withdrawpb.WithdrawResponse {
	return &withdrawpb.WithdrawResponse{
		AccountId: u.AccountID,
		Balance:   u.Balance,
	}
}
