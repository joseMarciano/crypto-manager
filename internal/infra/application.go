package infra

import (
	"github.com/joseMarciano/crypto-manager/internal/config"
	"github.com/joseMarciano/crypto-manager/internal/infra/database"
	"github.com/joseMarciano/crypto-manager/internal/infra/grpc"
	natspkg "github.com/joseMarciano/crypto-manager/internal/infra/nats"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"gorm.io/gorm"
)

type (
	Application struct {
		Server        *grpc.Server
		DB            *gorm.DB
		Nats          *nats.Conn
		JetStream     jetstream.JetStream
		Configuration config.Configuration
	}
)

func New() *Application {
	loader := config.NewLoader()
	configuration, err := loader.Load()
	if err != nil {
		panic(err)
	}

	gormDB, err := database.NewConnection(configuration.Database)
	if err != nil {
		panic(err)
	}

	natsConn, err := natspkg.New(configuration.Nats)
	if err != nil {
		panic(err)
	}

	streamConn, err := natspkg.NewStreamConn(natsConn)
	if err != nil {
		panic(err)
	}

	return &Application{
		Server:        grpc.New(),
		DB:            gormDB,
		Nats:          natsConn,
		JetStream:     streamConn,
		Configuration: configuration,
	}
}
