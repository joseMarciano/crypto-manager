package usecase

import (
	"context"
	"time"

	"github.com/joseMarciano/crypto-manager/internal/app/reports/exchangetransaction/domain"
	errorspkg "github.com/joseMarciano/crypto-manager/internal/errors"
	"github.com/joseMarciano/crypto-manager/internal/validator"
	timepkg "github.com/joseMarciano/crypto-manager/pkg/time"
)

type (
	ExchangeTransactionFinder interface {
		FindAllBetween(context.Context, time.Time, time.Time) ([]domain.ExchangeTransaction, error)
	}

	UseCase struct {
		exchangeTransactionFinder ExchangeTransactionFinder
	}
)

func New(etFinder ExchangeTransactionFinder) UseCase {
	return UseCase{exchangeTransactionFinder: etFinder}
}

func (uc UseCase) Execute(ctx context.Context, input Input) (Output, error) {
	if err := validator.BusinessValidate(ctx, input); err != nil {
		return Output{}, err
	}

	startDate, err := timepkg.ParseCanonical(input.StartDate)
	if err != nil {
		return Output{}, errorspkg.NewClientError("invalid start_date format", err)
	}

	endDate, err := timepkg.ParseCanonical(input.EndDate)
	if err != nil {
		return Output{}, errorspkg.NewClientError("invalid end_date format", err)
	}

	transactions, err := uc.exchangeTransactionFinder.FindAllBetween(ctx, startDate, endDate) //todo: add pagination if have time
	if err != nil {
		return Output{}, err
	}

	exchangeMap := mapTotalsByExchangeAndDate(transactions)
	return Output{Data: buildDatas(exchangeMap)}, nil
}

func mapTotalsByExchangeAndDate(transactions []domain.ExchangeTransaction) map[string]map[string][]Total {
	exchangeMap := make(map[string]map[string][]Total)
	for _, tx := range transactions {
		if _, ok := exchangeMap[tx.ExchangeID]; !ok {
			exchangeMap[tx.ExchangeID] = make(map[string][]Total)
		}

		exchangeMap[tx.ExchangeID][tx.ExecutedAt] = append(
			exchangeMap[tx.ExchangeID][tx.ExecutedAt],
			Total{Amount: tx.Amount, Type: tx.Type},
		)
	}
	return exchangeMap
}

func buildDatas(exchangeTotalsMap map[string]map[string][]Total) []Exchange {
	datas := make([]Exchange, 0, len(exchangeTotalsMap))
	for exchangeID, dateMap := range exchangeTotalsMap {
		dates := make([]Date, 0, len(dateMap))
		for date, totals := range dateMap {
			dates = append(dates, Date{Date: date, Totals: totals})
		}
		datas = append(datas, Exchange{ExchangeID: exchangeID, Dates: dates})
	}

	return datas
}
