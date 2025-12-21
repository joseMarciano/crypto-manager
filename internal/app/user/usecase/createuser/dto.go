package createuser

import (
	"time"

	"github.com/joseMarciano/crypto-manager/internal/app/user/domain"
)

type (
	Input struct {
		Name           string    `validate:"required"`
		Birthday       time.Time `validate:"required"`
		DocumentNumber string    `validate:"required"`
	}

	Output struct {
		ID             string
		Name           string
		Birthday       time.Time
		DocumentNumber string
	}
)

func (i Input) toDomain() domain.User {
	return domain.User{
		ID:             domain.GenerateID(),
		Name:           i.Name,
		Birthday:       i.Birthday,
		DocumentNumber: i.DocumentNumber,
	}
}

func toOutput(d domain.User) Output {
	return Output(d)
}
