//go:build integration

package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/joseMarciano/crypto-manager/internal/app/exchange/domain"
	"github.com/joseMarciano/crypto-manager/internal/app/exchange/repository"
	"github.com/joseMarciano/crypto-manager/internal/test/integration"
	"github.com/stretchr/testify/require"
)

func TestTransactionRepository_CreateTransaction(t *testing.T) {
	t.Parallel()

	suite, err := integration.NewSuite(t)
	require.NoError(t, err)
	require.NoError(t, suite.SetupFixtures("testdata/fixtures/repository_dependencies.sql"))

	repo := repository.NewTransactionRepository(suite.DB())
	want := domain.ExchangeTransaction{
		ID:         "50000000-0000-4000-8000-000000000001",
		ExchangeID: integrationAccountExchangeID,
		Type:       domain.Deposit,
		Amount:     19.95,
		ExecutedAt: time.Date(2024, time.January, 3, 12, 30, 0, 0, time.UTC),
	}

	created, err := repo.CreateTransaction(context.Background(), want)
	require.NoError(t, err)
	require.Equal(t, want, created)

	var persisted domain.ExchangeTransaction
	err = suite.DB().Table("exchange_transactions").Where("id = ?", want.ID).Take(&persisted).Error
	require.NoError(t, err)
	require.Equal(t, want, persisted)
}

func TestTransactionRepository_CreateTransaction_InvalidExchange(t *testing.T) {
	t.Parallel()

	suite, err := integration.NewSuite(t)
	require.NoError(t, err)

	repo := repository.NewTransactionRepository(suite.DB())
	_, err = repo.CreateTransaction(context.Background(), domain.ExchangeTransaction{
		ID:         "50000000-0000-4000-8000-000000000099",
		ExchangeID: "20000000-0000-4000-8000-000000000099",
		Type:       domain.Deposit,
		Amount:     19.95,
		ExecutedAt: time.Date(2024, time.January, 3, 12, 30, 0, 0, time.UTC),
	})
	require.Error(t, err)
}
