package config

import "github.com/nats-io/nats.go"

type Nats struct {
	Host string

	ExchangeTransactionLifecycleStream  string
	ExchangeTransactionLifecycleSubject string
}

func natsLoader(cfg *Configuration, props properties) {
	host, ok := props["NATS_HOST"]
	if !ok {
		host = nats.DefaultURL
	}

	cfg.Nats = Nats{
		Host:                                host,
		ExchangeTransactionLifecycleStream:  props["EXCHANGE_TRANSACTION_LIFECYCLE_STREAM"],
		ExchangeTransactionLifecycleSubject: props["EXCHANGE_TRANSACTION_LIFECYCLE_SUBJECT"],
	}
}
