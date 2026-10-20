//go:build integration

package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/joseMarciano/crypto-manager/internal/app/reports/exchangetransaction/domain"
	"github.com/joseMarciano/crypto-manager/internal/app/reports/exchangetransaction/repository"
	"github.com/joseMarciano/crypto-manager/internal/test/integration"
	"github.com/stretchr/testify/require"
)

const reportFixtureExchangeID = "aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaaa"

func TestTransactionRepository_FindAllBetween(t *testing.T) {
	t.Parallel()

	suite, err := integration.NewSuite(t)
	require.NoError(t, err)
	require.NoError(t, suite.SetupFixtures("testdata/fixtures/report_range.sql"))

	repo := repository.NewTransactionRepository(suite.DB())
	got, err := repo.FindAllBetween(
		context.Background(),
		time.Date(2024, time.January, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2024, time.January, 4, 0, 0, 0, 0, time.UTC),
	)
	require.NoError(t, err)

	want := []domain.ExchangeTransaction{
		{ExchangeID: reportFixtureExchangeID, ExecutedAt: "2024-01-02T00:00:00Z", Type: "TRANSFER", Amount: 11},
		{ExchangeID: reportFixtureExchangeID, ExecutedAt: "2024-01-03T00:00:00Z", Type: "DEPOSIT", Amount: 25},
		{ExchangeID: reportFixtureExchangeID, ExecutedAt: "2024-01-03T00:00:00Z", Type: "WITHDRAWAL", Amount: 7},
		{ExchangeID: reportFixtureExchangeID, ExecutedAt: "2024-01-04T00:00:00Z", Type: "DEPOSIT", Amount: 20},
	}
	require.ElementsMatch(t, want, got)
}

func TestTransactionRepository_FindAllBetween_Empty(t *testing.T) {
	t.Parallel()

	suite, err := integration.NewSuite(t)
	require.NoError(t, err)

	repo := repository.NewTransactionRepository(suite.DB())
	got, err := repo.FindAllBetween(
		context.Background(),
		time.Date(2024, time.January, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2024, time.January, 4, 0, 0, 0, 0, time.UTC),
	)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Empty(t, got)
}
