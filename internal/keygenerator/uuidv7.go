package keygenerator

import (
	"uuid"
)

func Generate() string {
	u := uuid.NewV7()
	return u.String()
}
