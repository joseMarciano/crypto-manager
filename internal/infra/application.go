package infra

import (
	"github.com/joseMarciano/crypto-manager/internal/config"
	"github.com/joseMarciano/crypto-manager/internal/infra/database"
	"github.com/joseMarciano/crypto-manager/internal/infra/grpc"
	"gorm.io/gorm"
)

type (
	Application struct {
		Server        *grpc.Server
		DB            *gorm.DB
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

	return &Application{
		Server:        grpc.New(),
		DB:            gormDB,
		Configuration: configuration,
	}
}
