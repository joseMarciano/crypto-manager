package rounder_test

import (
	"testing"

	"github.com/joseMarciano/crypto-manager/pkg/rounder"

	"github.com/stretchr/testify/require"
)

func TestTwoDecimalPlaces_Float32(t *testing.T) {
	tests := []struct {
		name     string
		input    float32
		expected float32
	}{
		{
			name:     "Basic rounding up",
			input:    3.14159,
			expected: 3.14,
		},
		{
			name:     "Rounding down",
			input:    2.718281,
			expected: 2.72,
		},
		{
			name:     "Exact two decimals",
			input:    5.25,
			expected: 5.25,
		},
		{
			name:     "Zero value",
			input:    0.0,
			expected: 0.0,
		},
		{
			name:     "Negative number rounding",
			input:    -1.999,
			expected: -2.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rounder.TwoDecimalPlaces(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestTwoDecimalPlaces_Float64(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			name:     "High precision rounding up",
			input:    3.14159265359,
			expected: 3.14,
		},
		{
			name:     "High precision rounding down",
			input:    2.71828182846,
			expected: 2.72,
		},
		{
			name:     "Exact two decimals with high precision",
			input:    5.25000000000,
			expected: 5.25,
		},
		{
			name:     "Very small positive number",
			input:    0.001,
			expected: 0.00,
		},
		{
			name:     "Large negative number",
			input:    -123456.789,
			expected: -123456.79,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rounder.TwoDecimalPlaces(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestTruncate_Float32(t *testing.T) {
	tests := []struct {
		name     string
		input    float32
		places   int
		expected float32
	}{
		{
			name:     "Truncate to 1 decimal place",
			input:    3.14159,
			places:   1,
			expected: 3.1,
		},
		{
			name:     "Truncate to 3 decimal places",
			input:    2.71828,
			places:   3,
			expected: 2.718,
		},
		{
			name:     "Truncate to 0 decimal places (whole number)",
			input:    9.87654,
			places:   0,
			expected: 10.0,
		},
		{
			name:     "Truncate negative number to 4 places",
			input:    -1.23456789,
			places:   4,
			expected: -1.2346,
		},
		{
			name:     "Truncate with more places than needed",
			input:    5.5,
			places:   5,
			expected: 5.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rounder.Truncate(tt.input, tt.places)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestTruncate_Float64(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		places   int
		expected float64
	}{
		{
			name:     "High precision truncate to 6 places",
			input:    3.141592653589793,
			places:   6,
			expected: 3.141593,
		},
		{
			name:     "Scientific notation input",
			input:    1.23456789e-3,
			places:   8,
			expected: 0.00123457,
		},
		{
			name:     "Large number truncate to 2 places",
			input:    999999.999999,
			places:   2,
			expected: 1000000.0,
		},
		{
			name:     "Very small number truncate to 10 places",
			input:    0.0000000001234567,
			places:   10,
			expected: 0.0000000001,
		},
		{
			name:     "Negative large number truncate to 1 place",
			input:    -987654.321,
			places:   1,
			expected: -987654.3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rounder.Truncate(tt.input, tt.places)
			require.Equal(t, tt.expected, result)
		})
	}
}
