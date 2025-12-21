package usecase

import (
	"context"

	exchangedomain "github.com/joseMarciano/crypto-manager/internal/app/exchange/domain"
	userdomain "github.com/joseMarciano/crypto-manager/internal/app/user/domain"
	"github.com/joseMarciano/crypto-manager/internal/validator"
	"github.com/joseMarciano/crypto-manager/pkg/rounder"
)

type (
	UserFinder interface {
		FindUserByID(context.Context, string) (userdomain.User, error)
	}

	AccountFinder interface {
		FindAccountsByUserID(context.Context, string) ([]exchangedomain.Account, error)
	}

	UseCase struct {
		userFinder    UserFinder
		accountFinder AccountFinder
	}
)

func New(uFinder UserFinder, accFinder AccountFinder) UseCase {
	return UseCase{userFinder: uFinder, accountFinder: accFinder}
}

func (uc UseCase) Execute(ctx context.Context, input Input) (Output, error) {
	if err := validator.BusinessValidate(ctx, input); err != nil {
		return Output{}, err
	}

	user, err := uc.userFinder.FindUserByID(ctx, input.UserID)
	if err != nil {
		return Output{}, err
	}

	accounts, err := uc.accountFinder.FindAccountsByUserID(ctx, user.ID)
	if err != nil {
		return Output{}, err
	}

	balances := make([]Balance, 0, len(accounts))
	totalBalance := 0.0
	for _, account := range accounts {
		if account.Balance <= 0 {
			continue
		}

		balances = append(balances, Balance{ExchangeID: account.ExchangeID, Balance: account.Balance})
		totalBalance += account.Balance
	}

	return Output{
		UserID:       user.ID,
		TotalBalance: rounder.TwoDecimalPlaces(totalBalance),
		Balances:     balances,
	}, nil
}
