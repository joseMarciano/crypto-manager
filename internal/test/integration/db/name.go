package db

import (
	"strings"
	"uuid"
)

func databaseName(testName string, id uuid.UUID) string {
	slug := databaseNameSlug(testName)
	const prefix = "it_"
	suffix := "_" + id.String()
	maxSlugLength := 63 - len(prefix) - len(suffix)
	if len(slug) > maxSlugLength {
		slug = strings.TrimRight(slug[:maxSlugLength], "_")
	}

	return prefix + slug + suffix
}

func databaseNameSlug(testName string) string {
	var slug strings.Builder
	underscorePending := false

	for _, char := range strings.ToLower(testName) {
		switch {
		case char >= 'a' && char <= 'z', char >= '0' && char <= '9':
			if underscorePending && slug.Len() > 0 {
				slug.WriteByte('_')
			}
			slug.WriteRune(char)
			underscorePending = false
		default:
			underscorePending = true
		}
	}

	if slug.Len() == 0 {
		return "test"
	}
	return slug.String()
}
