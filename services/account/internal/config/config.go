package config

import (
	"github.com/ssengalanto/runic/pkg/config"
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

// New creates a new *Config instance with default values loaded from a YAML file and environment variables.
func New() (*Config, error) {
	c := &Config{}

	res, err := config.New(filePath, c)
	if err != nil {
		return nil, err
	}

	cfg := res.(*Config) //nolint:errcheck //intentional panic

	return cfg, nil
}
