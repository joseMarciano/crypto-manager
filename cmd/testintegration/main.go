package main

import (
	"context"
	"os"

	"github.com/joseMarciano/crypto-manager/internal/test/integration"
)

func main() {
	os.Exit(integration.Run(context.Background(), os.Args[1:]))
}
