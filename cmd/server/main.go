package main

import (
	"context"

	"github.com/joseMarciano/crypto-manager/cmd/server/modules"
	"github.com/joseMarciano/crypto-manager/internal/logger"
)

func main() {
	app := modules.NewApp()
	logger.Info(context.Background(), "Application started", "port", app.Configuration.Server.Port)
	if err := app.Server.Start(app.Configuration.Server.Port); err != nil {
		panic(err) // todo: proper error handling
	}
}
