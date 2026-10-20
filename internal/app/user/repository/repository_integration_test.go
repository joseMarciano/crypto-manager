//go:build integration

package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/joseMarciano/crypto-manager/internal/app/user/domain"
	"github.com/joseMarciano/crypto-manager/internal/app/user/repository"
	errorspkg "github.com/joseMarciano/crypto-manager/internal/errors"
	"github.com/joseMarciano/crypto-manager/internal/test/integration"
	"github.com/stretchr/testify/require"
)

const integrationUserID = "40000000-0000-4000-8000-000000000001"

func TestRepository_CreateAndFindUserByID(t *testing.T) {
	t.Parallel()

	suite, err := integration.NewSuite(t)
	require.NoError(t, err)

	repo := repository.New(suite.DB())
	want := domain.User{
		ID:             integrationUserID,
		Name:           "integration user",
		Birthday:       time.Date(1990, time.January, 2, 0, 0, 0, 0, time.UTC),
		DocumentNumber: "integration-document-001",
	}

	created, err := repo.Create(context.Background(), want)
	require.NoError(t, err)
	require.Equal(t, want, created)

	got, err := repo.FindUserByID(context.Background(), want.ID)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestRepository_ExistsByNameAndDocument(t *testing.T) {
	t.Parallel()

	suite, err := integration.NewSuite(t)
	require.NoError(t, err)

	repo := repository.New(suite.DB())
	user := domain.User{
		ID:             integrationUserID,
		Name:           "integration user",
		Birthday:       time.Date(1990, time.January, 2, 0, 0, 0, 0, time.UTC),
		DocumentNumber: "integration-document-001",
	}
	_, err = repo.Create(context.Background(), user)
	require.NoError(t, err)

	exists, err := repo.ExistsByName(context.Background(), user.Name)
	require.NoError(t, err)
	require.True(t, exists)

	exists, err = repo.ExistsByName(context.Background(), "missing user")
	require.NoError(t, err)
	require.False(t, exists)

	exists, err = repo.ExistsByDocument(context.Background(), user.DocumentNumber)
	require.NoError(t, err)
	require.True(t, exists)

	exists, err = repo.ExistsByDocument(context.Background(), "missing-document")
	require.NoError(t, err)
	require.False(t, exists)
}

func TestRepository_FindUserByID_NotFound(t *testing.T) {
	t.Parallel()

	suite, err := integration.NewSuite(t)
	require.NoError(t, err)

	repo := repository.New(suite.DB())
	got, err := repo.FindUserByID(context.Background(), "40000000-0000-4000-8000-000000000099")
	require.ErrorIs(t, err, errorspkg.ErrNotFound)
	require.Equal(t, domain.User{}, got)
}
