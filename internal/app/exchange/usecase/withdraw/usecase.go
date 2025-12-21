package withdraw

import (
	"context"
	"fmt"

	"github.com/joseMarciano/crypto-manager/internal/app/exchange/domain"
	errorspkg "github.com/joseMarciano/crypto-manager/internal/errors"
	"github.com/joseMarciano/crypto-manager/internal/validator"
)

type (
	ExchangeTransactionCreator interface {
		CreateTransaction(context.Context, domain.ExchangeTransaction) (domain.ExchangeTransaction, error)
	}

	AccountUpdater interface {
		UpdateAccount(context.Context, domain.Account) (domain.Account, error)
	}

	AccountFinder interface {
		FindAccountByID(context.Context, string) (domain.Account, error)
	}

	ExchangeFinder interface {
		FindExchangeByID(context.Context, string) (domain.Exchange, error)
	}
	UseCase struct {
		exchangeTransactionCreator ExchangeTransactionCreator
		accountUpdater             AccountUpdater
		accountFinder              AccountFinder
		exchangeFinder             ExchangeFinder
	}
)

func New(etCreator ExchangeTransactionCreator, accUpdater AccountUpdater, accFinder AccountFinder, eFinder ExchangeFinder) UseCase {
	return UseCase{exchangeTransactionCreator: etCreator, accountUpdater: accUpdater, accountFinder: accFinder, exchangeFinder: eFinder}
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

	_, err = uc.exchangeTransactionCreator.CreateTransaction(ctx, toExchangeTransaction(exchange.ID, input.Amount)) //TODO: put this as async notification if have time
	return toOutput(updatedAccount), nil
}
