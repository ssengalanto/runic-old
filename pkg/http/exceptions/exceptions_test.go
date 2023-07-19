package exceptions_test

import (
	"net/http"
	"testing"

	apperr "github.com/ssengalanto/runic/pkg/exceptions"
	"github.com/ssengalanto/runic/pkg/http/exceptions"
	"github.com/stretchr/testify/assert"
)

func TestNewHTTPError(t *testing.T) {
	testCases := []struct {
		name     string
		err      error
		expected *exceptions.HTTPError
	}{
		{
			name:     "invalid error",
			err:      apperr.ErrInvalid,
			expected: &exceptions.HTTPError{Error: exceptions.Err{Code: http.StatusBadRequest, Message: http.StatusText(http.StatusBadRequest), Cause: "invalid"}},
		},
		{
			name:     "unauthorized error",
			err:      apperr.ErrUnauthorized,
			expected: &exceptions.HTTPError{Error: exceptions.Err{Code: http.StatusUnauthorized, Message: http.StatusText(http.StatusUnauthorized), Cause: "unauthorized"}},
		},
		{
			name:     "forbidden error",
			err:      apperr.ErrForbidden,
			expected: &exceptions.HTTPError{Error: exceptions.Err{Code: http.StatusForbidden, Message: http.StatusText(http.StatusForbidden), Cause: "forbidden"}},
		},
		{
			name:     "not found error",
			err:      apperr.ErrNotFound,
			expected: &exceptions.HTTPError{Error: exceptions.Err{Code: http.StatusNotFound, Message: http.StatusText(http.StatusNotFound), Cause: "not found"}},
		},
		{
			name:     "timeout error",
			err:      apperr.ErrTimeout,
			expected: &exceptions.HTTPError{Error: exceptions.Err{Code: http.StatusRequestTimeout, Message: http.StatusText(http.StatusRequestTimeout), Cause: "timeout"}},
		},
		{
			name:     "temporary disabled error",
			err:      apperr.ErrTemporaryDisabled,
			expected: &exceptions.HTTPError{Error: exceptions.Err{Code: http.StatusServiceUnavailable, Message: http.StatusText(http.StatusServiceUnavailable), Cause: "temporary disabled"}},
		},
		{
			name:     "internal server error",
			err:      apperr.ErrInternal,
			expected: &exceptions.HTTPError{Error: exceptions.Err{Code: http.StatusInternalServerError, Message: http.StatusText(http.StatusInternalServerError), Cause: "internal server error"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			httpError := exceptions.NewHTTPError(tc.err)
			assert.Equal(t, tc.expected, httpError)
		})
	}
}
