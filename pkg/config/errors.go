package config

import "fmt"

var (
	// ErrReadConfigFileFailed is an error that indicates a failure to read the configuration file.
	ErrReadConfigFileFailed = fmt.Errorf("failed to read config file")

	// ErrReadEnvFailed is an error that indicates a failure to read environment variables.
	ErrReadEnvFailed = fmt.Errorf("failed to read environment variabless")
)
