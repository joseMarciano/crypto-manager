package integration

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

type fixtureScript struct {
	path string
	sql  string
}

func (s *Suite) SetupFixtures(paths ...string) error {
	if len(paths) == 0 {
		return nil
	}

	scripts := make([]fixtureScript, 0, len(paths))
	for _, path := range paths {
		contents, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read fixture %q: %w", path, err)
		}
		scripts = append(scripts, fixtureScript{path: path, sql: string(contents)})
	}

	return s.db.Transaction(func(tx *gorm.DB) error {
		for _, script := range scripts {
			if err := tx.Exec(script.sql).Error; err != nil {
				return fmt.Errorf("execute fixture %q: %w", script.path, err)
			}
		}
		return nil
	})
}
