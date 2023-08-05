package server

import "fmt"

// ErrServerClosed is returned when the server failed to start.
var ErrServerClosed = fmt.Errorf("server is closed")
