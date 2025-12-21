package createaccount

import (
	"context"
	"fmt"

	"github.com/joseMarciano/crypto-manager/internal/app/exchange/domain"
	userdomain "github.com/joseMarciano/crypto-manager/internal/app/user/domain"
	errorspkg "github.com/joseMarciano/crypto-manager/internal/errors"
	"github.com/joseMarciano/crypto-manager/internal/validator"
)

type (
	AccountCreator interface {
		CreateAccount(context.Context, domain.Account) (domain.Account, error)
	}

	AccountFinder interface {
		ExistsAccountByUserAndExchange(context.Context, string, string) (bool, error)
	}

	ExchangeFinder interface {
		FindExchangeByID(context.Context, string) (domain.Exchange, error)
	}

	UserFinder interface {
		FindUserByID(context.Context, string) (userdomain.User, error)
	}

	UseCase struct {
		accountCreator AccountCreator
		accountFinder  AccountFinder
		exchangeFinder ExchangeFinder
		userFinder     UserFinder
	}
)

func New(accCreator AccountCreator, accFinder AccountFinder, eFinder ExchangeFinder, uFinder UserFinder) UseCase {
	return UseCase{accountCreator: accCreator, accountFinder: accFinder, exchangeFinder: eFinder, userFinder: uFinder}
}

func (uc UseCase) Execute(ctx context.Context, input Input) (Output, error) {
	if err := validator.BusinessValidate(ctx, input); err != nil {
		return Output{}, err
	}

	exchange, err := uc.exchangeFinder.FindExchangeByID(ctx, input.ExchangeID)
	if err != nil {
		return Output{}, err
	}

	user, err := uc.userFinder.FindUserByID(ctx, input.UserID)
	if err != nil {
		return Output{}, err
	}

	if user.Age() < exchange.MinimumAge {
		return Output{}, errorspkg.NewBusinessValidationError(fmt.Sprintf("user %s does not meet the minimum age requirement for exchange %s", input.UserID, input.ExchangeID), nil)
	}

	existsAccount, err := uc.accountFinder.ExistsAccountByUserAndExchange(ctx, input.UserID, input.ExchangeID)
	if err != nil {
		return Output{}, err
	}

	if existsAccount {
		return Output{}, errorspkg.NewBusinessValidationError(fmt.Sprintf("user %s already have an associated account in the exchange %s", input.UserID, input.ExchangeID), nil)
	}

	account := input.toDomain()
	savedAccount, err := uc.accountCreator.CreateAccount(ctx, account)
	if err != nil {
		return Output{}, err
	}

	return toOutput(savedAccount), nil
}
