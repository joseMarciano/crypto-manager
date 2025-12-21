package keygenerator

import (
	"github.com/google/uuid"
)

func Generate() string {
	u, _ := uuid.NewV7()
	return u.String()
}
