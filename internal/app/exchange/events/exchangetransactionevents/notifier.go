package exchangetransactionevents

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/joseMarciano/crypto-manager/internal/app/exchange/domain"
	"github.com/joseMarciano/crypto-manager/internal/config"
	"github.com/nats-io/nats.go/jetstream"
)

type (
	Notifier struct {
		jetStream jetstream.JetStream
		subject   string
	}
)

func NewNotifier(jetStream jetstream.JetStream, cfg config.Nats) (Notifier, error) {
	_, err := jetStream.CreateOrUpdateStream(context.Background(), jetstream.StreamConfig{
		Name:        cfg.ExchangeTransactionLifecycleStream,
		Description: "Stream to manage exchange transaction life cycle",
		Subjects:    []string{cfg.ExchangeTransactionLifecycleSubject},
	})

	return Notifier{jetStream: jetStream, subject: cfg.ExchangeTransactionLifecycleSubject}, err
}

func (n Notifier) Notify(ctx context.Context, transaction domain.ExchangeTransaction) error {
	payload, err := json.Marshal(toEvent(transaction))
	if err != nil {
		return fmt.Errorf("error on marshal payload for NotifyCreated: %w", err)
	}

	_, err = n.jetStream.Publish(ctx, n.subject, payload)
	if err != nil {
		return fmt.Errorf("error on publish payload to %s stream: %w", n.subject, err)
	}

	return nil
}
