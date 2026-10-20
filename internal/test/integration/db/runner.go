package db

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"time"

	"uuid"

	"github.com/jackc/pgx/v5"
	mobyContainer "github.com/moby/moby/api/types/container"
	mobyNetwork "github.com/moby/moby/api/types/network"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/network"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	postgresImage       = "postgres:16.4"
	postgresNetworkName = "integration-postgres"
	templateDatabase    = "integration_template"
	postgresUsername    = "integration"
	postgresPassword    = "integration"
	liquibaseImageRepo  = "crypto-manager-liquibase"
	liquibaseImageTag   = "integration"
)

type Runner struct {
	postgres         *postgres.PostgresContainer
	network          *testcontainers.DockerNetwork
	adminDSN         string
	templateDatabase string
}

func Start(ctx context.Context, repoRoot string) (_ *Runner, returnErr error) {
	runner := &Runner{templateDatabase: templateDatabase}
	defer func() {
		cleanup(runner, returnErr)
	}()

	var err error
	runner.network, err = network.New(ctx)
	if err != nil {
		return nil, fmt.Errorf("create integration Docker network: %w", err)
	}

	if runner.network == nil {
		return nil, errors.New("create integration Docker network: no network was returned")
	}

	runner.postgres, err = postgres.Run(
		ctx,
		postgresImage,
		postgres.WithDatabase("postgres"),
		postgres.WithUsername(postgresUsername),
		postgres.WithPassword(postgresPassword),
		testcontainers.WithWaitStrategy(wait.ForSQL(
			"5432/tcp",
			"pgx",
			func(host string, port mobyNetwork.Port) string {
				return fmt.Sprintf(
					"postgres://%s:%s@%s/postgres?sslmode=disable",
					postgresUsername,
					postgresPassword,
					net.JoinHostPort(host, port.Port()),
				)
			},
		).WithStartupTimeout(2*time.Minute)),
		network.WithNetwork([]string{postgresNetworkName}, runner.network),
	)

	if err != nil {
		return nil, fmt.Errorf("start integration PostgreSQL container: %w", err)
	}

	if runner.postgres == nil {
		return nil, errors.New("start integration PostgreSQL container: no container was returned")
	}

	runner.adminDSN, err = runner.postgres.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("get integration PostgreSQL connection string: %w", err)
	}

	if err = createTemplateDatabase(ctx, runner.adminDSN, runner.templateDatabase); err != nil {
		return nil, err
	}
	if err = runLiquibase(ctx, repoRoot, runner.network.Name, runner.templateDatabase); err != nil {
		return nil, err
	}

	return runner, nil
}

func (r *Runner) AdminDSN() string {
	return r.adminDSN
}

func (r *Runner) TemplateDatabase() string {
	return r.templateDatabase
}

func (r *Runner) Close(ctx context.Context) error {
	var closeErr error
	if r.postgres != nil {
		if err := r.postgres.Terminate(ctx); err != nil {
			closeErr = errors.Join(closeErr, fmt.Errorf("terminate integration PostgreSQL container: %w", err))
		}
		r.postgres = nil
	}
	if r.network != nil {
		if err := r.network.Remove(ctx); err != nil {
			closeErr = errors.Join(closeErr, fmt.Errorf("remove integration Docker network: %w", err))
		}
		r.network = nil
	}

	return closeErr
}

func cleanup(runner *Runner, err error) {
	if err == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err = runner.Close(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "close integration test resources: %v\n", err)
	}
}

func createTemplateDatabase(ctx context.Context, adminDSN string, database string) error {
	connection, err := pgx.Connect(ctx, adminDSN)
	if err != nil {
		return fmt.Errorf("connect to integration PostgreSQL admin database: %w", err)
	}
	defer connection.Close(context.Background())

	_, err = connection.Exec(ctx, "CREATE DATABASE "+pgx.Identifier{database}.Sanitize())
	if err != nil {
		return fmt.Errorf("create integration template database: %w", err)
	}
	return nil
}

func runLiquibase(ctx context.Context, repoRoot string, networkName string, database string) (returnErr error) {
	changelogPath, err := filepath.Abs(filepath.Join(repoRoot, "db", "changelog"))
	if err != nil {
		return fmt.Errorf("resolve Liquibase changelog path: %w", err)
	}
	if _, err = os.Stat(changelogPath); err != nil {
		return fmt.Errorf("find Liquibase changelog at %q: %w", changelogPath, err)
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		Context:    repoRoot,
		Dockerfile: "docker/liquibase.Dockerfile",
		Repo:       liquibaseImageRepo,
		Tag:        liquibaseImageTag,
		KeepImage:  true,
		Env: map[string]string{
			"LIQUIBASE_COMMAND_CHANGELOG_FILE": "db.changelog-master.yaml",
			"LIQUIBASE_COMMAND_URL":            "jdbc:postgresql://" + postgresNetworkName + ":5432/" + database,
			"LIQUIBASE_COMMAND_USERNAME":       postgresUsername,
			"LIQUIBASE_COMMAND_PASSWORD":       postgresPassword,
		},
		Cmd:      []string{"update"},
		Networks: []string{networkName},
		HostConfigModifier: func(hostConfig *mobyContainer.HostConfig) {
			hostConfig.Binds = append(hostConfig.Binds, changelogPath+":/liquibase/changelog:ro")
		},
		WaitingFor: wait.ForExit().
			WithExitTimeout(5 * time.Minute).
			WithPollInterval(250 * time.Millisecond),
		Started: true,
	})
	if container == nil {
		if err != nil {
			return fmt.Errorf("start Liquibase container: %w", err)
		}
		return errors.New("start Liquibase container: container was not returned")
	}
	defer func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		returnErr = errors.Join(returnErr, container.Terminate(cleanupCtx))
	}()
	if err != nil {
		return fmt.Errorf("run Liquibase update: %w", err)
	}

	state, err := container.State(ctx)
	if err != nil {
		return fmt.Errorf("inspect Liquibase container: %w", err)
	}
	if state.ExitCode != 0 {
		logs, logErr := readContainerLogs(ctx, container)
		if logErr != nil {
			return fmt.Errorf("Liquibase update exited with code %d (also failed to read logs: %v)", state.ExitCode, logErr)
		}
		return fmt.Errorf("Liquibase update exited with code %d: %s", state.ExitCode, logs)
	}

	return nil
}

func readContainerLogs(ctx context.Context, container testcontainers.Container) (string, error) {
	logs, err := container.Logs(ctx)
	if err != nil {
		return "", err
	}
	defer logs.Close()

	output, err := io.ReadAll(logs)
	return string(output), err
}

func newDatabaseName(testName string) string {
	return databaseName(testName, uuid.New())
}
