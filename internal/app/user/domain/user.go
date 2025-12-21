package domain

import (
	"time"

	"github.com/joseMarciano/crypto-manager/internal/keygenerator"
)

type User struct {
	ID             string
	Name           string
	Birthday       time.Time
	DocumentNumber string
}

func (u User) Age() int {
	return int(time.Since(u.Birthday).Hours() / 24 / 365)
}

func GenerateID() string {
	return keygenerator.Generate()
}
