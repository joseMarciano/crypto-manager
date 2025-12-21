package time_test

import (
	"testing"
	"time"

	timepkg "github.com/joseMarciano/crypto-manager/pkg/time"

	"github.com/stretchr/testify/require"
)

func TestParseCanonical(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expected      time.Time
		expectedError string
	}{
		{
			name:          "valid date",
			input:         "2023-12-25",
			expected:      time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC),
			expectedError: "",
		},
		{
			name:          "empty string",
			input:         "",
			expected:      time.Time{},
			expectedError: "",
		},
		{
			name:          "null string",
			input:         "null",
			expected:      time.Time{},
			expectedError: "",
		},
		{
			name:          "invalid format - datetime",
			input:         "2023-12-25 10:30:00",
			expected:      time.Time{},
			expectedError: "error on parse 2023-12-25 10:30:00 - valid format is 2006-01-02",
		},
		{
			name:          "invalid format - wrong separator",
			input:         "2023/12/25",
			expected:      time.Time{},
			expectedError: "error on parse 2023/12/25 - valid format is 2006-01-02",
		},
		{
			name:          "invalid date",
			input:         "2023-13-25",
			expected:      time.Time{},
			expectedError: "error on parse 2023-13-25 - valid format is 2006-01-02",
		},
		{
			name:          "invalid day",
			input:         "2023-02-30",
			expected:      time.Time{},
			expectedError: "error on parse 2023-02-30 - valid format is 2006-01-02",
		},
		{
			name:          "leap year valid",
			input:         "2024-02-29",
			expected:      time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC),
			expectedError: "",
		},
		{
			name:          "non-leap year invalid",
			input:         "2023-02-29",
			expected:      time.Time{},
			expectedError: "error on parse 2023-02-29 - valid format is 2006-01-02",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := timepkg.ParseCanonical(tt.input)

			if tt.expectedError != "" {
				require.EqualError(t, err, tt.expectedError, "expected an error but got none")
				return
			}

			require.NoError(t, err, "expected no error")
			require.Equal(t, tt.expected, result, "expected result does not match")
		})
	}
}

func TestFormatCanonical(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected string
	}{
		{
			name:     "valid date",
			input:    time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC),
			expected: "2023-12-25",
		},
		{
			name:     "zero time",
			input:    time.Time{},
			expected: "",
		},
		{
			name:     "date with time components ignored",
			input:    time.Date(2023, 12, 25, 15, 30, 45, 123456789, time.UTC),
			expected: "2023-12-25",
		},
		{
			name:     "leap year date",
			input:    time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC),
			expected: "2024-02-29",
		},
		{
			name:     "first day of year",
			input:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: "2023-01-01",
		},
		{
			name:     "last day of year",
			input:    time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
			expected: "2023-12-31",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := timepkg.FormatCanonical(tt.input)
			require.Equal(t, tt.expected, result, "expected result does not match")
		})
	}
}
