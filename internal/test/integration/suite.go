package integration

import (
	"testing"

	"github.com/joseMarciano/crypto-manager/internal/test/integration/db"
	"gorm.io/gorm"
)

type Suite struct {
	db *gorm.DB
}

func NewSuite(t *testing.T) (*Suite, error) {
	t.Helper()

	database, err := db.Open(t)
	if err != nil {
		return nil, err
	}

	return &Suite{db: database}, nil
}

func (s *Suite) DB() *gorm.DB {
	return s.db
}
