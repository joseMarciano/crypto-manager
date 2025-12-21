package createexchange

import (
	"github.com/joseMarciano/crypto-manager/internal/app/exchange/domain"
)

type (
	Input struct {
		Name                  string  `validate:"required"`
		MinimumAge            int     `validate:"required,gte=1"`
		MaximumTransferAmount float64 `validate:"required,gte=1"`
	}

	Output struct {
		ID                    string
		Name                  string
		MinimumAge            int
		MaximumTransferAmount float64
	}
)

func (i Input) toDomain() domain.Exchange {
	return domain.Exchange{
		ID:                    domain.GenerateID(),
		Name:                  i.Name,
		MinimumAge:            i.MinimumAge,
		MaximumTransferAmount: i.MaximumTransferAmount,
	}
}

func toOutput(d domain.Exchange) Output {
	return Output(d)
}
