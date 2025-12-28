package exchangetransactionevents

import (
	"github.com/joseMarciano/crypto-manager/internal/app/exchange/domain"
	timepkg "github.com/joseMarciano/crypto-manager/pkg/time"
)

type (
	Event struct {
		Payload Payload `json:"payload"`
	}

	Payload struct {
		ID         string  `json:"id"`
		ExchangeID string  `json:"exchange_id"`
		Type       string  `json:"type"`
		Amount     float64 `json:"amount"`
		ExecutedAt string  `json:"executed_at"`
	}
)

func (e Event) toDomain() (domain.ExchangeTransaction, error) {
	executedAt, err := timepkg.ParseRFC3339(e.Payload.ExecutedAt)
	if err != nil {
		return domain.ExchangeTransaction{}, err
	}

	return domain.ExchangeTransaction{
		ID:         e.Payload.ID,
		ExchangeID: e.Payload.ExchangeID,
		Type:       domain.Type(e.Payload.Type),
		Amount:     e.Payload.Amount,
		ExecutedAt: executedAt,
	}, nil
}

func toEvent(transaction domain.ExchangeTransaction) Event {
	return Event{
		Payload: Payload{
			ID:         transaction.ID,
			ExchangeID: transaction.ExchangeID,
			Type:       string(transaction.Type),
			Amount:     transaction.Amount,
			ExecutedAt: timepkg.FormatRFC3339(transaction.ExecutedAt),
		},
	}
}
