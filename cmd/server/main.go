package main

import (
	"log"

	"github.com/joseMarciano/crypto-manager/cmd/server/modules"
)

func main() {
	app := modules.NewApp()
	log.Println("Application started on port", app.Configuration.Server.Port)
	if err := app.Server.Start(app.Configuration.Server.Port); err != nil {
		panic(err) // todo: proper error handling
	}
}
