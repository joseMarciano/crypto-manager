package createuser_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/joseMarciano/crypto-manager/internal/app/user/domain"
	"github.com/joseMarciano/crypto-manager/internal/app/user/usecase/createuser"
	"github.com/joseMarciano/crypto-manager/internal/app/user/usecase/createuser/mocks"
	"github.com/joseMarciano/crypto-manager/internal/app/user/usecase/createuser/testdata"
	errorspkg "github.com/joseMarciano/crypto-manager/internal/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUseCase_Execute(t *testing.T) {
	ctx := context.Background()
	t.Run("given a valid input, should execute with no error", func(t *testing.T) {
		input := testdata.DefaultInput()

		finder := mocks.NewFinder(t)
		creator := mocks.NewCreator(t)

		finder.EXPECT().ExistsByDocument(ctx, input.DocumentNumber).Return(false, nil)
		finder.EXPECT().ExistsByName(ctx, input.Name).Return(false, nil)
		creator.EXPECT().Create(ctx, mock.AnythingOfType("domain.User")).Return(testdata.ExpectedCreatedUser(), nil)

		useCase := createuser.New(creator, finder)
		output, err := useCase.Execute(ctx, input)
		require.NoError(t, err, "should execute with no error")
		require.Equal(t, testdata.ExpectedOutput(), output)
	})

	t.Run("given a user with existing name, should return business validation error", func(t *testing.T) {
		input := testdata.DefaultInput()

		finder := mocks.NewFinder(t)
		creator := mocks.NewCreator(t)

		finder.EXPECT().ExistsByName(ctx, input.Name).Return(true, nil)

		useCase := createuser.New(creator, finder)
		_, err := useCase.Execute(ctx, input)

		require.EqualError(t, err, fmt.Sprintf("BUSINESS: user %s already exists", input.Name))
		require.ErrorAs(t, err, &errorspkg.BusinessValidationError{}, "error should be of type BusinessValidationError")
	})

	t.Run("given a user with existing document number, should return business validation error", func(t *testing.T) {
		input := testdata.DefaultInput()

		finder := mocks.NewFinder(t)
		creator := mocks.NewCreator(t)

		finder.EXPECT().ExistsByName(ctx, input.Name).Return(false, nil)
		finder.EXPECT().ExistsByDocument(ctx, input.DocumentNumber).Return(true, nil)

		useCase := createuser.New(creator, finder)
		_, err := useCase.Execute(ctx, input)

		require.EqualError(t, err, fmt.Sprintf("BUSINESS: user with document %s already exists", input.DocumentNumber))
		require.ErrorAs(t, err, &errorspkg.BusinessValidationError{}, "error should be of type BusinessValidationError")
	})

	t.Run("given finder error on name check, should return error", func(t *testing.T) {
		input := testdata.DefaultInput()
		finder := mocks.NewFinder(t)
		creator := mocks.NewCreator(t)

		finder.EXPECT().ExistsByName(ctx, input.Name).Return(false, assert.AnError)

		useCase := createuser.New(creator, finder)
		_, err := useCase.Execute(ctx, input)

		require.Error(t, err, "should return finder error")
		require.Equal(t, assert.AnError, err)
	})

	t.Run("given finder error on document check, should return error", func(t *testing.T) {
		input := testdata.DefaultInput()

		finder := mocks.NewFinder(t)
		creator := mocks.NewCreator(t)

		finder.EXPECT().ExistsByName(ctx, input.Name).Return(false, nil)
		finder.EXPECT().ExistsByDocument(ctx, input.DocumentNumber).Return(false, assert.AnError)

		useCase := createuser.New(creator, finder)
		_, err := useCase.Execute(ctx, input)

		require.Error(t, err, "should return finder error")
		require.Equal(t, assert.AnError, err)
	})

	t.Run("given creator error, should return error", func(t *testing.T) {
		input := testdata.DefaultInput()

		finder := mocks.NewFinder(t)
		creator := mocks.NewCreator(t)

		finder.EXPECT().ExistsByName(ctx, input.Name).Return(false, nil)
		finder.EXPECT().ExistsByDocument(ctx, input.DocumentNumber).Return(false, nil)
		creator.EXPECT().Create(ctx, mock.AnythingOfType("domain.User")).Return(domain.User{}, assert.AnError)

		useCase := createuser.New(creator, finder)
		_, err := useCase.Execute(ctx, input)
		require.Error(t, err, "should return creator error")
		require.Equal(t, assert.AnError, err)
	})
}
