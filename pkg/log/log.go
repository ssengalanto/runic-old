package log

import (
	"fmt"
	"strings"

	"github.com/ssengalanto/runic/pkg/constants"
	"go.uber.org/zap"
)

// Log is a log implementation that uses zap.Logger internally.
type Log struct {
	zap *zap.Logger
}

// KeyValue represents a key-value pair for log fields.
type KeyValue struct {
	Key   string // The key or name of the log field.
	Value any    // The value associated with the log field.
}

// Field creates a new log field with the given key and value.
func Field(key string, value any) KeyValue {
	return KeyValue{
		Key:   key,
		Value: value,
	}
}

// New creates a new Logger instance with the specified environment.
// It returns a pointer to the Logger and an error if the initialization fails.
func New(env string) (*Log, error) {
	log, err := buildZapLogger(env)
	if err != nil {
		return nil, err
	}

	return &Log{
		zap: log,
	}, nil
}

// Info logs an informational message with additional fields.
func (l *Log) Info(msg string, fields ...KeyValue) {
	l.zap.Info(msg, toZapFields(fields)...)
}

// Error logs an error message with additional fields.
func (l *Log) Error(msg string, fields ...KeyValue) {
	l.zap.Error(msg, toZapFields(fields)...)
}

// Debug logs a debug message with additional fields.
func (l *Log) Debug(msg string, fields ...KeyValue) {
	l.zap.Debug(msg, toZapFields(fields)...)
}

// Warn logs a warning message with additional fields.
func (l *Log) Warn(msg string, fields ...KeyValue) {
	l.zap.Warn(msg, toZapFields(fields)...)
}

// Fatal logs a fatal message with additional fields and causes a panic.
func (l *Log) Fatal(msg string, fields ...KeyValue) {
	l.zap.Fatal(msg, toZapFields(fields)...)
}

// Panic logs a message with additional fields and causes a panic.
func (l *Log) Panic(msg string, fields ...KeyValue) {
	l.zap.Panic(msg, toZapFields(fields)...)
}

// toZapFields maps the log fields to zap fields.
func toZapFields(fields []KeyValue) []zap.Field {
	zapFields := make([]zap.Field, len(fields))
	for i, f := range fields {
		zapFields[i] = zap.Any(f.Key, f.Value)
	}
	return zapFields
}

// buildZapLogger builds a new zap.Logger for the specified environment with predefined configuration.
// It returns the log and an error if the initialization fails or the environment is invalid.
func buildZapLogger(env string) (*zap.Logger, error) {
	env = strings.ToLower(env)
	buildProviders := getBuildProviders()

	for i, provider := range buildProviders {
		if strings.ToLower(provider.env()) == env {
			log, err := provider.build()
			if err != nil {
				return nil, fmt.Errorf("%w: %s", ErrInitializationFailed, err)
			}
			return log, nil
		}

		if i == len(buildProviders)-1 {
			return nil, fmt.Errorf(
				"%w: invalid env with value of `%s`, must be one of the following: %v",
				ErrInitializationFailed,
				env,
				[]string{constants.DevEnv, constants.StgEnv, constants.ProdEnv},
			)
		}
	}

	return nil, fmt.Errorf(
		"%w: invalid env with value of `%s`, must be one of the following: %v",
		ErrInitializationFailed,
		env,
		[]string{constants.DevEnv, constants.StgEnv, constants.ProdEnv},
	)
}
