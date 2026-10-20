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

const integrationExchangeID = "20000000-0000-4000-8000-000000000001"

func TestExchangeRepository_CreateAndFindExchangeByID(t *testing.T) {
	t.Parallel()

	suite, err := integration.NewSuite(t)
	require.NoError(t, err)

	repo := repository.NewExchangeRepository(suite.DB())
	want := domain.Exchange{
		ID:                    integrationExchangeID,
		Name:                  "integration exchange",
		MinimumAge:            18,
		MaximumTransferAmount: 12500.75,
	}

	created, err := repo.CreateExchange(context.Background(), want)
	require.NoError(t, err)
	require.Equal(t, want, created)

	got, err := repo.FindExchangeByID(context.Background(), want.ID)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestExchangeRepository_FindExchangeByID_NotFound(t *testing.T) {
	t.Parallel()

	suite, err := integration.NewSuite(t)
	require.NoError(t, err)

	repo := repository.NewExchangeRepository(suite.DB())
	got, err := repo.FindExchangeByID(context.Background(), "20000000-0000-4000-8000-000000000099")
	require.ErrorIs(t, err, errorspkg.ErrNotFound)
	require.Equal(t, domain.Exchange{}, got)
}
