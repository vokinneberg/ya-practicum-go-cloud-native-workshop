// config_test.go
package config

import (
	"flag"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNew(t *testing.T) {
	// Save original command-line arguments and defer restoration
	origArgs := os.Args
	defer func() { os.Args = origArgs }()

	// Test cases
	tests := []struct {
		name           string
		envVars        map[string]string
		args           []string
		expectedConfig Config
	}{
		{
			name: "default_values",
			envVars: map[string]string{
				"SERVER_ADDRESS": "",
				"BASE_URL":       "",
				"VERBOSE":        "",
			},
			args: []string{},
			expectedConfig: Config{
				ServerAddress: "8080",
				BaseURL:       "http://localhost:8080",
				Verbose:       false,
			},
		},
		{
			name: "environment_variables",
			envVars: map[string]string{
				"SERVER_ADDRESS": ":9090",
				"BASE_URL":       "http://example.com",
				"VERBOSE":        "true",
			},
			args: []string{},
			expectedConfig: Config{
				ServerAddress: ":9090",
				BaseURL:       "http://example.com",
				Verbose:       true,
			},
		},
		{
			name: "command_line_flags",
			envVars: map[string]string{
				"SERVER_ADDRESS": "",
				"BASE_URL":       "",
				"VERBOSE":        "",
			},
			args: []string{"-a", ":9090", "-b", "http://example.com", "-v"},
			expectedConfig: Config{
				ServerAddress: ":9090",
				BaseURL:       "http://example.com",
				Verbose:       true,
			},
		},
		{
			name: "command_line_flags_override_environment_variables",
			envVars: map[string]string{
				"SERVER_ADDRESS": ":8080",
				"BASE_URL":       "http://localhost:8080",
				"VERBOSE":        "false",
			},
			args: []string{"-a", ":9090", "-b", "http://example.com", "-v"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment variables
			for key, value := range tt.envVars {
				if err := os.Setenv(key, value); err != nil {
					t.Fatalf("failed to set environment variable %s: %v", key, err)
				}
			}

			// Set command-line arguments
			os.Args = append([]string{"cmd"}, tt.args...)

			// Reset flags to avoid conflicts between tests
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

			// Create new config
			cfg, err := New()
			if err != nil {
				t.Fatalf("failed to create config: %v", err)
			}

			// Compare expected and actual config
			if diff := cmp.Diff(cfg, &tt.expectedConfig); diff != "" {
				t.Errorf("config mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
