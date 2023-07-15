// Package config provides tools for managing application configuration.
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	// The filePath constant defines the path to the YAML configuration file.
	filePath = "/services/account/config.yml"
)

type (
	// Config represents the application configuration structure.
	Config struct {
		App
		HTTP
		PGSQL
	}
	// App represents the application-specific configuration.
	App struct {
		Name string `env-required:"true" yaml:"name" env:"APP_NAME"`
		Env  string `env-required:"true" yaml:"env" env:"APP_ENV"`
	}
	// HTTP represents the HTTP server configuration.
	HTTP struct {
		Port int `env-required:"true" yaml:"port" env:"ACCOUNT_SERVICE_HTTP_PORT"`
	}
	// PGSQL represents the PGSQL database configuration.
	PGSQL struct {
		Username    string `env-required:"true" yaml:"username" env:"PGSQL_USERNAME"`
		Password    string `env-required:"true" yaml:"password" env:"PGSQL_PASSWORD"`
		Host        string `env-required:"true" yaml:"host" env:"PGSQL_HOST"`
		Port        int    `env-required:"true" yaml:"port" env:"PGSQL_PORT"`
		DBName      string `env-required:"true" yaml:"db_name" env:"PGSQL_DB_NAME"`
		QueryParams string `env-required:"true" yaml:"query_params" env:"PGSQL_QUERY_PARAMS"`
	}
)

// New creates a new Config instance with default values loaded from a YAML file and environment variables.
//
// New initializes a new Config struct and reads configuration values from a YAML file named "config.yml" in the
// current working directory using the cleanenv package. It then reads additional configuration values from
// environment variables using the same package.
//
// Example usage:
// cfg, err := New()
// if err != nil {
// // handle error
// }
// // use cfg values
//
// New returns a pointer to a new Config instance and an error if any occurred during initialization.
// The returned Config struct contains the default configuration values from the YAML file and any overrides
// from environment variables.
func New() (*Config, error) {
	cfg := &Config{}

	// Get the current working directory.
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Load configuration values from the YAML file.
	fp := fmt.Sprintf("%s%s", dir, filePath)

	err = cleanenv.ReadConfig(fp, cfg)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrReadConfigFileFailed, err)
	}

	// Load additional configuration values from environment variables.
	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrReadEnvFailed, err)
	}

	return cfg, nil
}
