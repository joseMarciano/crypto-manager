package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	gormpostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	AdminDSNEnv         = "INTEGRATION_POSTGRES_ADMIN_DSN"
	TemplateDatabaseEnv = "INTEGRATION_POSTGRES_TEMPLATE_DB"
)

func Open(t *testing.T) (*gorm.DB, error) {
	t.Helper()

	adminDSN := strings.TrimSpace(os.Getenv(AdminDSNEnv))
	if adminDSN == "" {
		return nil, fmt.Errorf("%s is missing; run tests with mise run test-integration", AdminDSNEnv)
	}
	template := strings.TrimSpace(os.Getenv(TemplateDatabaseEnv))
	if template == "" {
		return nil, fmt.Errorf("%s is missing; run tests with mise run test-integration", TemplateDatabaseEnv)
	}

	adminConn, err := pgx.Connect(t.Context(), adminDSN)
	if err != nil {
		return nil, fmt.Errorf("connect to integration PostgreSQL admin database: %w", err)
	}
	defer adminConn.Close(context.Background())

	database, err := createTestDatabase(t.Context(), adminConn, t.Name(), template)
	if err != nil {
		return nil, err
	}

	var sqlPool *sql.DB
	t.Cleanup(func() {
		if sqlPool != nil {
			if err = sqlPool.Close(); err != nil {
				t.Errorf("close test database connection pool: %v", err)
			}
		}

		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		cleanupConn, cError := pgx.Connect(cleanupCtx, adminDSN)
		if cError != nil {
			t.Errorf("connect to drop test database %q: %v", database, cError)
			return
		}
		defer cleanupConn.Close(cleanupCtx)

		_, cError = cleanupConn.Exec(cleanupCtx, "DROP DATABASE "+pgx.Identifier{database}.Sanitize())
		if cError != nil {
			t.Errorf("drop test database %q: %v", database, cError)
		}
	})

	config, err := pgx.ParseConfig(adminDSN)
	if err != nil {
		return nil, fmt.Errorf("parse integration PostgreSQL connection string: %w", err)
	}
	config.Database = database
	sqlPool = stdlib.OpenDB(*config)
	sqlPool.SetMaxOpenConns(2)
	sqlPool.SetMaxIdleConns(1)

	gormDB, err := gorm.Open(gormpostgres.New(gormpostgres.Config{Conn: sqlPool}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		return nil, fmt.Errorf("open GORM connection to test database %q: %w", database, err)
	}
	return gormDB, nil
}

func createTestDatabase(ctx context.Context, adminConn *pgx.Conn, testName string, template string) (string, error) {
	database := newDatabaseName(testName)
	statement := "CREATE DATABASE " + pgx.Identifier{database}.Sanitize() + " TEMPLATE " + pgx.Identifier{template}.Sanitize()
	if _, err := adminConn.Exec(ctx, statement); err != nil {
		return "", fmt.Errorf("create test database %q from template %q: %w", database, template, err)
	}

	return database, nil
}
