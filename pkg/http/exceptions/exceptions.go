package exceptions

import (
	"errors"
	"net/http"

	"github.com/ssengalanto/runic/pkg/exceptions"
)

// HTTPError represents an HTTP error with a custom error structure.
type HTTPError struct {
	Error Err `json:"error"`
} // @name HTTPError

// Err represents the details of an error, including code, message, and cause.
type Err struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Cause   string `json:"cause"`
} // @name Error

// NewHTTPError creates a new HTTPError instance based on the given error.
func NewHTTPError(err error) *HTTPError {
	code := http.StatusInternalServerError

	switch {
	case errors.Is(err, exceptions.ErrInvalid):
		code = http.StatusBadRequest
	case errors.Is(err, exceptions.ErrUnauthorized):
		code = http.StatusUnauthorized
	case errors.Is(err, exceptions.ErrForbidden):
		code = http.StatusForbidden
	case errors.Is(err, exceptions.ErrNotFound):
		code = http.StatusNotFound
	case errors.Is(err, exceptions.ErrTimeout):
		code = http.StatusRequestTimeout
	case errors.Is(err, exceptions.ErrTemporaryDisabled):
		code = http.StatusServiceUnavailable
	}

	httpError := &HTTPError{
		Error: Err{
			Code:    code,
			Message: http.StatusText(code),
			Cause:   err.Error(),
		},
	}

	return httpError
}
