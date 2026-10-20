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

	"github.com/joseMarciano/crypto-manager/internal/test/integration/db"
)

const (
	integrationCoverageProfile = "coverage-integration.out"
	integrationCoverageHTML    = "coverage-integration.html"
)

func Run(ctx context.Context, args []string) (exitCode int) {
	repoRoot, err := repositoryRoot()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	runner, err := db.Start(ctx, repoRoot)
	if err != nil {
		fmt.Fprintf(os.Stderr, "start integration test resources: %v\n", err)
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
		fmt.Fprintf(os.Stderr, "run go test for integration tests: %v\n", err)
		return 1
	}

	profilePath := coverageProfilePath(args)
	if err = generateCoverageHTML(ctx, repoRoot, profilePath); err != nil {
		fmt.Fprintf(os.Stderr, "generate integration coverage HTML: %v\n", err)
		return 1
	}
	fmt.Fprintf(os.Stdout, "Integration coverage profile: %s\nIntegration coverage HTML: %s\n", profilePath, integrationCoverageHTML)

	return 0
}

func cleanup(runner *db.Runner) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := runner.Close(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "close integration test resources: %v\n", err)
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
