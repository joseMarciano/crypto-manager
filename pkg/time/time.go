package time

import (
	"fmt"
	"time"
)

func ParseCanonical(value string) (time.Time, error) {
	return parse(value, time.DateOnly)
}

func ParseRFC3339(value string) (time.Time, error) {
	return parse(value, time.RFC3339)
}
func FormatCanonical(t time.Time) string {
	return format(t, time.DateOnly)
}

func FormatRFC3339(t time.Time) string {
	return format(t, time.RFC3339)
}

func parse(value, layout string) (time.Time, error) {
	if value == "" || value == "null" {
		return time.Time{}, nil
	}

	t, err := time.Parse(layout, value)
	if err != nil {
		return time.Time{}, fmt.Errorf("error on parse %s - valid format is %s", value, layout)
	}
	return t, nil
}

func format(t time.Time, layout string) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(layout)
}
