package createexchange

import (
	"github.com/joseMarciano/crypto-manager/internal/app/exchange/usecase/createexchange"
	createexchangepb "github.com/joseMarciano/crypto-manager/pkg/proto/exchange/create"
)

type (
	Input struct {
		Name                  string  `json:"name"`
		MinimumAge            int     `json:"minimum_age"`
		MaximumTransferAmount float64 `json:"maximum_transfer_amount"`
	}

	Output struct {
		ID                    string  `json:"id"`
		Name                  string  `json:"name"`
		MinimumAge            int     `json:"minimum_age"`
		MaximumTransferAmount float64 `json:"maximum_transfer_amount"`
	}
)

func toInputDTO(request *createexchangepb.CreateExchangeRequest) createexchange.Input {
	return createexchange.Input{
		Name:                  request.GetName(),
		MinimumAge:            int(request.GetMinimumAge()),
		MaximumTransferAmount: request.GetMaximumTransferAmount(),
	}
}

func toResponse(u createexchange.Output) *createexchangepb.CreateExchangeResponse {
	return &createexchangepb.CreateExchangeResponse{
		Id:                    u.ID,
		Name:                  u.Name,
		MinimumAge:            uint32(u.MinimumAge),
		MaximumTransferAmount: u.MaximumTransferAmount,
	}
}
