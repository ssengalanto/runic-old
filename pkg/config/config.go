package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// New initializes a new Config struct and reads configuration values from the specified YAML file and environment
// variables using the cleanenv package.
//
// Example usage:
//
//	cfg := &MyConfigStruct{}
//	err := New("config.yml", cfg)
//	if err != nil {
//	  // handle error
//	}
//	// use cfg values
//
// New returns a pointer to the new Config instance and an error if any occurred during initialization.
// The returned Config struct contains the default configuration values from the YAML file and any overrides
// from environment variables.
func New[T any](filePath string, cfg *T) (*T, error) {
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
