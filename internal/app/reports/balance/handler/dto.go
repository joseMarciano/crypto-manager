package handler

import (
	"github.com/joseMarciano/crypto-manager/internal/app/reports/balance/usecase"
	balancepb "github.com/joseMarciano/crypto-manager/pkg/proto/report/balance"
	slicespkg "github.com/joseMarciano/crypto-manager/pkg/slices"
)

func toInputDTO(request *balancepb.FetchBalanceRequest) usecase.Input {
	return usecase.Input{
		UserID: request.GetUserId(),
	}
}

func toResponse(u usecase.Output) *balancepb.FetchBalanceResponse {
	return &balancepb.FetchBalanceResponse{
		UserId:       u.UserID,
		TotalBalance: u.TotalBalance,
		Balances:     slicespkg.Map(u.Balances, toBalancePB),
	}
}

func toBalancePB(b usecase.Balance) *balancepb.Balance {
	return &balancepb.Balance{
		ExchangeId: b.ExchangeID,
		Balance:    b.Balance,
	}
}
