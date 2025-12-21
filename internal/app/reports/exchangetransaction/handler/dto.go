package handler

import (
	"github.com/joseMarciano/crypto-manager/internal/app/reports/exchangetransaction/usecase"
	exchangetransactionpb "github.com/joseMarciano/crypto-manager/pkg/proto/report/exchangetransaction"
	slicespkg "github.com/joseMarciano/crypto-manager/pkg/slices"
)

func toInputDTO(request *exchangetransactionpb.FetchExchangeTransactionRequest) usecase.Input {
	return usecase.Input{
		StartDate: request.GetStartDate(),
		EndDate:   request.GetEndDate(),
	}
}

func toResponse(u usecase.Output) *exchangetransactionpb.FetchExchangeTransactionResponse {
	return &exchangetransactionpb.FetchExchangeTransactionResponse{
		Data: slicespkg.Map(u.Data, toExchangePB),
	}
}

func toExchangePB(e usecase.Exchange) *exchangetransactionpb.Exchange {
	return &exchangetransactionpb.Exchange{
		ExchangeId: e.ExchangeID,
		Dates:      slicespkg.Map(e.Dates, toDatePB),
	}
}

func toDatePB(d usecase.Date) *exchangetransactionpb.Date {
	return &exchangetransactionpb.Date{
		Date:   d.Date,
		Totals: slicespkg.Map(d.Totals, toTotalPB),
	}
}

func toTotalPB(t usecase.Total) *exchangetransactionpb.Total {
	return &exchangetransactionpb.Total{
		Amount: t.Amount,
		Type:   t.Type,
	}
}
