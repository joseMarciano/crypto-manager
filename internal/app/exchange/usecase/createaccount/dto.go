package createaccount

import (
	"github.com/joseMarciano/crypto-manager/internal/app/exchange/domain"
)

type (
	Input struct {
		UserID     string `validate:"required"`
		ExchangeID string `validate:"required"`
	}

	Output struct {
		ID         string
		UserID     string
		ExchangeID string
		Balance    float64
	}
)

func (i Input) toDomain() domain.Account {
	return domain.Account{
		ID:         domain.GenerateID(),
		UserID:     i.UserID,
		ExchangeID: i.ExchangeID,
		Balance:    0,
	}
}

func toOutput(d domain.Account) Output {
	return Output(d)
}
