package config

import (
	"flag"
	"fmt"

	"github.com/caarlos0/env"
)

type Config struct {
	AppName       string `env:"APP_NAME" envDefault:"shortener"`
	BaseURL       string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	Environment   string `env:"ENVIRONMENT" envDefault:"development"`
	LogLevel      string `env:"LOG_LEVEL" envDefault:"info"`
	ServerAddress string `env:"SERVER_ADDRESS" envDefault:"8080"`
	Verbose       bool   `env:"VERBOSE" envDefault:"false"`
	Version       string
}

// New creates a new Config instance and parses environment variables and command-line flags.
// Command-line flags have precedence over environment variables.
// It returns the Config instance or an error if the environment variables or command-line flags are invalid.
func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse environment variables: %w", err)
	}

	// Define command-line flags
	serverAddressFlag := flag.String("a", "", "Server address")
	baseURLFlag := flag.String("b", "", "Base URL")
	verboseFlag := flag.Bool("v", false, "Verbose output")

	// Parse command-line flags
	flag.Parse()

	if *serverAddressFlag != "" {
		cfg.ServerAddress = *serverAddressFlag
	}
	if *baseURLFlag != "" {
		cfg.BaseURL = *baseURLFlag
	}
	if *verboseFlag {
		cfg.Verbose = *verboseFlag
	}

	return cfg, nil
}
