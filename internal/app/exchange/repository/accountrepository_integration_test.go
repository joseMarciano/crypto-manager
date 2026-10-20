//go:build integration

package repository_test

import (
	"context"
	"testing"

	"github.com/joseMarciano/crypto-manager/internal/app/exchange/domain"
	"github.com/joseMarciano/crypto-manager/internal/app/exchange/repository"
	errorspkg "github.com/joseMarciano/crypto-manager/internal/errors"
	"github.com/joseMarciano/crypto-manager/internal/test/integration"
	"github.com/stretchr/testify/require"
)

const (
	integrationAccountUserID       = "10000000-0000-4000-8000-000000000001"
	integrationSecondAccountUserID = "10000000-0000-4000-8000-000000000002"
	integrationAccountExchangeID   = "20000000-0000-4000-8000-000000000001"
	integrationSecondExchangeID    = "20000000-0000-4000-8000-000000000002"
)

func TestAccountRepository_CreateExistsFindUpdateAndFindByUserID(t *testing.T) {
	t.Parallel()

	suite, err := integration.NewSuite(t)
	require.NoError(t, err)
	require.NoError(t, suite.SetupFixtures("testdata/fixtures/repository_dependencies.sql"))

	repo := repository.NewAccountRepository(suite.DB())
	first := domain.Account{
		ID:         "30000000-0000-4000-8000-000000000001",
		UserID:     integrationAccountUserID,
		ExchangeID: integrationAccountExchangeID,
		Balance:    125.50,
	}
	created, err := repo.CreateAccount(context.Background(), first)
	require.NoError(t, err)
	require.Equal(t, first, created)

	exists, err := repo.ExistsAccountByUserAndExchange(context.Background(), first.UserID, first.ExchangeID)
	require.NoError(t, err)
	require.True(t, exists)

	exists, err = repo.ExistsAccountByUserAndExchange(context.Background(), integrationSecondAccountUserID, integrationSecondExchangeID)
	require.NoError(t, err)
	require.False(t, exists)

	got, err := repo.FindAccountByID(context.Background(), first.ID)
	require.NoError(t, err)
	require.Equal(t, first, got)

	updated := first
	updated.Balance = 250.75
	got, err = repo.UpdateAccount(context.Background(), updated)
	require.NoError(t, err)
	require.Equal(t, updated, got)

	persisted, err := repo.FindAccountByID(context.Background(), first.ID)
	require.NoError(t, err)
	require.Equal(t, updated, persisted)

	second := domain.Account{
		ID:         "30000000-0000-4000-8000-000000000002",
		UserID:     integrationAccountUserID,
		ExchangeID: integrationSecondExchangeID,
		Balance:    300.00,
	}
	_, err = repo.CreateAccount(context.Background(), second)
	require.NoError(t, err)

	otherUser := domain.Account{
		ID:         "30000000-0000-4000-8000-000000000003",
		UserID:     integrationSecondAccountUserID,
		ExchangeID: integrationAccountExchangeID,
		Balance:    400.00,
	}
	_, err = repo.CreateAccount(context.Background(), otherUser)
	require.NoError(t, err)

	accounts, err := repo.FindAccountsByUserID(context.Background(), integrationAccountUserID)
	require.NoError(t, err)
	require.ElementsMatch(t, []domain.Account{updated, second}, accounts)
}

func TestAccountRepository_FindAccountByID_NotFound(t *testing.T) {
	t.Parallel()

	suite, err := integration.NewSuite(t)
	require.NoError(t, err)

	repo := repository.NewAccountRepository(suite.DB())
	got, err := repo.FindAccountByID(context.Background(), "30000000-0000-4000-8000-000000000099")
	require.ErrorIs(t, err, errorspkg.ErrNotFound)
	require.Equal(t, domain.Account{}, got)
}
