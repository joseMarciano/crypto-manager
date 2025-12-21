package domain

import (
	"github.com/joseMarciano/crypto-manager/internal/keygenerator"
)

type (
	Exchange struct {
		ID                    string
		Name                  string
		MinimumAge            int
		MaximumTransferAmount float64
	}
)

func GenerateID() string {
	return keygenerator.Generate()
}
