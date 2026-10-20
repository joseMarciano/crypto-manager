package integration

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/joseMarciano/crypto-manager/internal/logger"
	"github.com/joseMarciano/crypto-manager/internal/test/integration/db"
)

const (
	integrationCoverageProfile = "coverage-integration.out"
	integrationCoverageHTML    = "coverage-integration.html"
)

func Run(ctx context.Context, args []string) (exitCode int) {
	repoRoot, err := repositoryRoot()
	if err != nil {
		logger.Error(ctx, "failed to locate integration repository root", "error", err)
		return 1
	}

	runner, err := db.Start(ctx, repoRoot)
	if err != nil {
		logger.Error(ctx, "failed to start integration test resources", "error", err)
		return 1
	}
	defer cleanup(runner)

	commandArgs := append([]string{"test"}, buildGoTestArgs(args)...)
	command := exec.CommandContext(ctx, "go", commandArgs...)
	command.Dir = repoRoot
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Env = runnerEnvironment(runner, os.Environ())
	if err = command.Run(); err != nil {
		if exitErr, ok := errors.AsType[*exec.ExitError](err); ok {
			return exitErr.ExitCode()
		}
		logger.Error(ctx, "failed to run integration tests", "error", err)
		return 1
	}

	profilePath := coverageProfilePath(args)
	if err = generateCoverageHTML(ctx, repoRoot, profilePath); err != nil {
		logger.Error(ctx, "failed to generate integration coverage HTML", "error", err)
		return 1
	}
	logger.Info(ctx, "integration tests completed", "coverage_profile", profilePath, "coverage_html", integrationCoverageHTML)

	return 0
}

func cleanup(runner *db.Runner) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := runner.Close(ctx); err != nil {
		logger.Error(ctx, "failed to close integration test resources", "error", err)
	}
}

func buildGoTestArgs(args []string) []string {
	result := []string{
		"-tags=integration",
		"-count=1",
		"-parallel=10",
		"-coverprofile=" + integrationCoverageProfile,
	}
	if len(args) == 0 {
		return append(result, "./...")
	}
	return append(result, args...)
}

func coverageProfilePath(args []string) string {
	profile := integrationCoverageProfile
	for i, arg := range args {
		if after, ok := strings.CutPrefix(arg, "-coverprofile="); ok {
			profile = after
			continue
		}
		if arg == "-coverprofile" && i+1 < len(args) {
			profile = args[i+1]
		}
	}
	return profile
}

func generateCoverageHTML(ctx context.Context, repoRoot string, profilePath string) error {
	command := exec.CommandContext(
		ctx,
		"go",
		"tool",
		"cover",
		"-html="+profilePath,
		"-o",
		integrationCoverageHTML,
	)
	command.Dir = repoRoot
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	if err := command.Run(); err != nil {
		return fmt.Errorf("run go tool cover: %w", err)
	}
	return nil
}

func repositoryRoot() (string, error) {
	_, sourceFile, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("locate integration runner source file")
	}
	root := filepath.Clean(filepath.Join(filepath.Dir(sourceFile), "..", "..", ".."))
	if _, err := os.Stat(filepath.Join(root, "go.mod")); err != nil {
		return "", fmt.Errorf("locate repository root from %q: %w", sourceFile, err)
	}
	return root, nil
}

func runnerEnvironment(runner *db.Runner, environment []string) []string {
	adminPrefix := fmt.Sprintf("%s=", db.AdminDSNEnv)
	templatePrefix := fmt.Sprintf("%s=", db.TemplateDatabaseEnv)
	filtered := make([]string, 0, len(environment)+2)
	for _, entry := range environment {
		if strings.HasPrefix(entry, adminPrefix) || strings.HasPrefix(entry, templatePrefix) {
			continue
		}
		filtered = append(filtered, entry)
	}

	return append(
		filtered,
		adminPrefix+runner.AdminDSN(),
		templatePrefix+runner.TemplateDatabase(),
	)
}
