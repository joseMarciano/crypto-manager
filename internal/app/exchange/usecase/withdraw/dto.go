package withdraw

import (
	"time"

	"github.com/joseMarciano/crypto-manager/internal/app/exchange/domain"
	"github.com/joseMarciano/crypto-manager/pkg/rounder"
)

type (
	Input struct {
		ExchangeID string  `validate:"required"`
		AccountID  string  `validate:"required"`
		Amount     float64 `validate:"required,gt=0,two-decimals"`
	}

	Output struct {
		AccountID string
		Balance   float64
	}
)

func toOutput(d domain.Account) Output {
	return Output{AccountID: d.ID, Balance: d.Balance}
}

func toExchangeTransaction(exchangeID string, amount float64) domain.ExchangeTransaction {
	return domain.ExchangeTransaction{
		ID:         domain.GenerateID(),
		ExchangeID: exchangeID,
		Type:       domain.Withdrawal,
		Amount:     rounder.TwoDecimalPlaces(amount),
		ExecutedAt: time.Now(),
	}
}
