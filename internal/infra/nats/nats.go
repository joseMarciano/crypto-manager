package nats

import (
	"github.com/joseMarciano/crypto-manager/internal/config"
	"github.com/nats-io/nats.go/jetstream"
)
import "github.com/nats-io/nats.go"

func New(config config.Nats) (*nats.Conn, error) {
	return nats.Connect(config.Host)
}

func NewStreamConn(conn *nats.Conn) (jetstream.JetStream, error) {
	return jetstream.New(conn)
}
