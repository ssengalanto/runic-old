package exception

import (
	"fmt"
)

var (
	// ErrInvalid is used to represent an "invalid" error.
	ErrInvalid = New("invalid")

	// ErrUnauthorized is used to represent an "unauthorized" error.
	ErrUnauthorized = New("unauthorized")

	// ErrForbidden is used to represent a "forbidden" error.
	ErrForbidden = New("forbidden")

	// ErrNotFound is used to represent a "not found" error.
	ErrNotFound = New("not found")

	// ErrInternal is used to represent an "internal server error" error.
	ErrInternal = New("internal server error")

	// ErrUnknown is used to represent an "unknown error".
	ErrUnknown = New("unknown error")

	// ErrTemporaryDisabled is used to represent a "temporary disabled" error.
	ErrTemporaryDisabled = New("temporary disabled")

	// ErrTimeout is used to represent a "timeout" error.
	ErrTimeout = New("timeout")
)

// Error represents a custom error with an associated message and optional cause.
type Error struct {
	message string
	cause   error
}

// New creates a new Error instance with the given message.
func New(message string) *Error {
	return &Error{
		message: message,
	}
}

// Wrap wraps an existing error with a new Error instance, providing additional context.
func Wrap(err error, message string) *Error {
	return &Error{
		message: message,
		cause:   err,
	}
}

// Unwrap returns the underlying cause of the error, if available.
func (e *Error) Unwrap() error {
	return e.cause
}

// Error implements the error interface for the Error type.
func (e *Error) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("%s: %v", e.message, e.cause)
	}
	return e.message
}
