package createexchange

import (
	"context"

	"github.com/joseMarciano/crypto-manager/internal/app/exchange/domain"
	"github.com/joseMarciano/crypto-manager/internal/validator"
)

type (
	Creator interface {
		CreateExchange(context.Context, domain.Exchange) (domain.Exchange, error)
	}

	UseCase struct {
		creator Creator
	}
)

func New(creator Creator) UseCase {
	return UseCase{creator: creator}
}

func (uc UseCase) Execute(ctx context.Context, input Input) (Output, error) {
	if err := validator.BusinessValidate(ctx, input); err != nil {
		return Output{}, err
	}

	exchange := input.toDomain()
	savedExchange, err := uc.creator.CreateExchange(ctx, exchange)
	if err != nil {
		return Output{}, err
	}

	return toOutput(savedExchange), nil
}
