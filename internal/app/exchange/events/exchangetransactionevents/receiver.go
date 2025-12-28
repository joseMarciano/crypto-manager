package exchangetransactionevents

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/joseMarciano/crypto-manager/internal/app/exchange/domain"
	"github.com/joseMarciano/crypto-manager/internal/config"
	"github.com/nats-io/nats.go/jetstream"
)

type (
	ExchangeTransactionCreator interface {
		CreateTransaction(context.Context, domain.ExchangeTransaction) (domain.ExchangeTransaction, error)
	}

	Receiver struct {
		consumer                   jetstream.Consumer
		exchangeTransactionCreator ExchangeTransactionCreator
	}
)

func NewReceiver(etCreator ExchangeTransactionCreator, jetStream jetstream.JetStream, cfg config.Nats) (Receiver, error) {
	consumerConfig := jetstream.ConsumerConfig{
		Durable:       "crypto-manager-consumer",
		AckPolicy:     jetstream.AckExplicitPolicy,
		DeliverPolicy: jetstream.DeliverNewPolicy,
		MaxDeliver:    2,
		FilterSubject: cfg.ExchangeTransactionLifecycleSubject,
	}
	cons, err := jetStream.CreateOrUpdateConsumer(context.Background(), cfg.ExchangeTransactionLifecycleStream, consumerConfig)
	if err != nil {
		return Receiver{}, fmt.Errorf("error on creating exchange transaction events consumer: %w", err)
	}

	return Receiver{consumer: cons, exchangeTransactionCreator: etCreator}, nil
}

func (r Receiver) Start() error {
	_, err := r.consumer.Consume(r.process)
	if err != nil {
		return fmt.Errorf("error on starting exchange transaction events receiver: %w", err)
	}

	return nil
}

func (r Receiver) process(msg jetstream.Msg) {
	var event Event
	if err := json.Unmarshal(msg.Data(), &event); err != nil {
		log.Println("error on unmarshalling exchange transaction event:", err)
		msg.Nak()
		return
	}

	transaction, err := event.toDomain()
	if err != nil {
		log.Println("error on converting exchange transaction event to domain:", err)
		msg.Nak()
		return
	}

	_, err = r.exchangeTransactionCreator.CreateTransaction(context.Background(), transaction)
	if err != nil {
		log.Println("error on creating exchange transaction from event:", err)
		msg.Nak()
		return
	}

	msg.Ack()
}
