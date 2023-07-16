package log

import (
	"github.com/ssengalanto/runic/pkg/constants"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// The buildProvider interface defines the methods that a build provider must implement.
type buildProvider interface {
	// env returns the environment associated with the build provider.
	env() string

	// build creates and returns a new zap.Logger instance based on the specific build provider's configuration.
	// It returns the log and an error if the initialization fails.
	build() (*zap.Logger, error)
}

// Build provider for development environment.
type development struct{}

func (d development) env() string {
	return constants.DevEnv
}
func (d development) build() (*zap.Logger, error) {
	cfg := createDevelopmentConfig()

	logger, err := cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}

	return logger, nil
}

// Build provider for staging environment.
type staging struct{}

func (t staging) env() string {
	return constants.StgEnv
}
func (t staging) build() (*zap.Logger, error) {
	cfg := createStagingConfig()

	logger, err := cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}

	return logger, nil
}

// Build provider for production environment.
type production struct{}

func (p production) env() string {
	return constants.ProdEnv
}
func (p production) build() (*zap.Logger, error) {
	cfg := createProductionConfig()

	logger, err := cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}

	return logger, nil
}

// createDevelopmentConfig creates a new zap config for development environment.
func createDevelopmentConfig() zap.Config {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.LevelKey = "level"
	cfg.EncoderConfig.CallerKey = "caller"
	cfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return cfg
}

// createProductionConfig creates a new zap config for staging environment.
func createStagingConfig() zap.Config {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	cfg.Encoding = "json"
	cfg.DisableStacktrace = true
	cfg.DisableCaller = true

	return cfg
}

// createProductionConfig creates a new zap config for production environment.
func createProductionConfig() zap.Config {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	cfg.Encoding = "json"
	cfg.DisableStacktrace = true
	cfg.DisableCaller = true

	return cfg
}

// getBuildProviders returns a slice of buildProvider.
func getBuildProviders() []buildProvider {
	return []buildProvider{
		development{}, staging{}, production{},
	}
}
