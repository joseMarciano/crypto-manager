package withdraw

import (
	"context"
	"fmt"

	"github.com/joseMarciano/crypto-manager/internal/app/exchange/domain"
	errorspkg "github.com/joseMarciano/crypto-manager/internal/errors"
	"github.com/joseMarciano/crypto-manager/internal/validator"
)

type (
	AccountUpdater interface {
		UpdateAccount(context.Context, domain.Account) (domain.Account, error)
	}

	AccountFinder interface {
		FindAccountByID(context.Context, string) (domain.Account, error)
	}

	ExchangeFinder interface {
		FindExchangeByID(context.Context, string) (domain.Exchange, error)
	}

	ExchangeTransactionNotifier interface {
		Notify(ctx context.Context, transaction domain.ExchangeTransaction) error
	}

	UseCase struct {
		accountUpdater              AccountUpdater
		accountFinder               AccountFinder
		exchangeFinder              ExchangeFinder
		exchangeTransactionNotifier ExchangeTransactionNotifier
	}
)

func New(etNotifier ExchangeTransactionNotifier, accUpdater AccountUpdater, accFinder AccountFinder, eFinder ExchangeFinder) UseCase {
	return UseCase{exchangeTransactionNotifier: etNotifier, accountUpdater: accUpdater, accountFinder: accFinder, exchangeFinder: eFinder}
}

func (uc UseCase) Execute(ctx context.Context, input Input) (Output, error) {
	if err := validator.BusinessValidate(ctx, input); err != nil {
		return Output{}, err
	}

	exchange, err := uc.exchangeFinder.FindExchangeByID(ctx, input.ExchangeID)
	if err != nil {
		return Output{}, err
	}

	account, err := uc.accountFinder.FindAccountByID(ctx, input.AccountID)
	if err != nil {
		return Output{}, err
	}

	if account.ExchangeID != exchange.ID {
		return Output{}, errorspkg.NewClientError(fmt.Sprintf("account %s does not belong to exchange %s", account.ID, exchange.ID), nil)
	}

	if input.Amount > exchange.MaximumTransferAmount {
		return Output{}, errorspkg.NewBusinessValidationError(fmt.Sprintf("deposit amount %.2f exceeds the maximum transfer amount for exchange %s", input.Amount, exchange.ID), nil)
	}

	account.WithDraw(input.Amount) //TODO: put redis lock here if have time to avoid race condition when multiple deposits happen simultaneously
	if account.Balance < 0 {
		return Output{}, errorspkg.NewBusinessValidationError(fmt.Sprintf("account %s has insufficient balance for withdrawal of %.2f", account.ID, input.Amount), nil)
	}

	updatedAccount, err := uc.accountUpdater.UpdateAccount(ctx, account)
	if err != nil {
		return Output{}, err
	}

	err = uc.exchangeTransactionNotifier.Notify(ctx, toExchangeTransaction(exchange.ID, input.Amount))
	if err != nil {
		return Output{}, errorspkg.NewUnexpectedError("unexpected error notifying exchange transaction", err)
	}

	return toOutput(updatedAccount), nil
}
