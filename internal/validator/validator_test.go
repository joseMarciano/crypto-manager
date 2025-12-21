package validator_test

import (
	"context"
	"testing"

	"github.com/joseMarciano/crypto-manager/internal/validator"
	"github.com/joseMarciano/crypto-manager/internal/validator/testdata"

	"github.com/stretchr/testify/require"
)

func TestBusinessValidate(t *testing.T) {
	ctx := context.Background()
	t.Run("should not return error for a valid struct", func(t *testing.T) {
		require.NoError(t, validator.BusinessValidate(ctx, testdata.DefaultTestStruct()), "should not return error for a valid struct")
	})

	t.Run("should return error for on required validation", func(t *testing.T) {
		s := testdata.DefaultTestStruct()
		s.Field1 = ""

		require.EqualError(t, validator.BusinessValidate(ctx, s), "BUSINESS: Field1 is required field", "should return error for on required validation")
	})
}

func TestBusinessValidate_GteValidation(t *testing.T) {
	ctx := context.Background()
	t.Run("should pass when value equals minimum threshold", func(t *testing.T) {
		s := testdata.DefaultTestStruct()
		s.Field2 = 2

		require.NoError(t, validator.BusinessValidate(ctx, s), "should pass when value equals minimum threshold")
	})

	t.Run("should fail when value is less than minimum threshold", func(t *testing.T) {
		s := testdata.DefaultTestStruct()
		s.Field2 = 1

		require.EqualError(t, validator.BusinessValidate(ctx, s), "BUSINESS: Field2 should be greater than or equal to 2", "should fail when value is less than minimum threshold")
	})
}

func TestBusinessValidate_GtValidation(t *testing.T) {
	ctx := context.Background()
	t.Run("should pass when value is greater than threshold", func(t *testing.T) {
		s := testdata.DefaultTestStruct()
		s.Field3 = 2

		require.NoError(t, validator.BusinessValidate(ctx, s), "should pass when value is greater than threshold")
	})

	t.Run("should fail when value is less than threshold", func(t *testing.T) {
		s := testdata.DefaultTestStruct()
		s.Field3 = 0

		require.EqualError(t, validator.BusinessValidate(ctx, s), "BUSINESS: Field3 should be greater than 1", "should fail when value is less than threshold")
	})
}

func TestBusinessValidate_TwoDecimalsValidation(t *testing.T) {
	ctx := context.Background()
	t.Run("should pass with whole numbers", func(t *testing.T) {
		s := testdata.DefaultTestStruct()
		s.Field4 = 100.0

		require.NoError(t, validator.BusinessValidate(ctx, s), "should pass with whole numbers")
	})

	t.Run("should fail with values having too many decimals", func(t *testing.T) {
		s := testdata.DefaultTestStruct()
		s.Field4 = 50.123

		require.EqualError(t, validator.BusinessValidate(ctx, s), "BUSINESS: Field4 should have at most two decimal places", "should fail with negative values having too many decimals")
	})
}
