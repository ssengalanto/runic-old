package interfaces

import "github.com/ssengalanto/runic/pkg/log"

// Logger is an interface consisting of the core log methods.
type Logger interface {
	// Info logs a message at the info level with optional fields.
	Info(msg string, fields ...log.KeyValue)

	// Error logs a message at the error level with optional fields.
	Error(msg string, fields ...log.KeyValue)

	// Debug logs a message at the debug level with optional fields.
	Debug(msg string, fields ...log.KeyValue)

	// Warn logs a message at the warn level with optional fields.
	Warn(msg string, fields ...log.KeyValue)

	// Fatal logs a message at the fatal level with optional fields.
	Fatal(msg string, fields ...log.KeyValue)

	// Panic logs a message at the panic level with optional fields.
	Panic(msg string, fields ...log.KeyValue)
}
