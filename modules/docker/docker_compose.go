package docker

import (
	"strings"

	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/shell"
	"github.com/gruntwork-io/terratest/modules/testing"
	"github.com/stretchr/testify/require"
)

// Options are Docker options.
type Options struct {
	WorkingDir string
	EnvVars    map[string]string
	// Set a logger that should be used. See the logger package for more info.
	Logger *logger.Logger
}

// RunDockerCompose runs docker compose with the given arguments and options and return stdout/stderr.
func RunDockerCompose(t testing.TestingT, options *Options, args ...string) string {
	out, err := runDockerComposeE(t, false, options, args...)
	if err != nil {
		t.Fatal(err)
	}
	return out
}

// RunDockerComposeAndGetStdout runs docker compose with the given arguments and options and returns only stdout.
func RunDockerComposeAndGetStdOut(t testing.TestingT, options *Options, args ...string) string {
	out, err := runDockerComposeE(t, true, options, args...)
	require.NoError(t, err)
	return out
}

// RunDockerComposeE runs docker compose with the given arguments and options and return stdout/stderr.
func RunDockerComposeE(t testing.TestingT, options *Options, args ...string) (string, error) {
	return runDockerComposeE(t, false, options, args...)
}

func runDockerComposeE(t testing.TestingT, stdout bool, options *Options, args ...string) (string, error) {
	cmd := shell.Command{
		Command: "docker",
		// We append --project-name to ensure containers from multiple different tests using Docker Compose don't end
		// up in the same project and end up conflicting with each other.
		Args:       append([]string{"compose", "--project-name", strings.ToLower(t.Name())}, args...),
		WorkingDir: options.WorkingDir,
		Env:        options.EnvVars,
		Logger:     options.Logger,
	}

	if stdout {
		return shell.RunCommandAndGetStdOut(t, cmd), nil
	}
	return shell.RunCommandAndGetOutputE(t, cmd)
}
