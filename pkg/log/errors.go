package log

import "fmt"

// Errors used by the zap package.
var (
	// ErrInitializationFailed is returned when log instance initialization failed.
	ErrInitializationFailed = fmt.Errorf("initialization failed")
)
