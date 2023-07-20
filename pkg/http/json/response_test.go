package json_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ssengalanto/runic/pkg/exceptions"
	httpexceptions "github.com/ssengalanto/runic/pkg/http/exceptions"
	httpjson "github.com/ssengalanto/runic/pkg/http/json"
	"github.com/stretchr/testify/assert"
)

// MockPayload is a mock payload struct for testing.
type MockPayload struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func TestSuccess(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		payload    *MockPayload
		expected   MockPayload
	}{
		{
			name:       "nil paylaod",
			statusCode: http.StatusOK,
			payload:    nil,
			expected:   MockPayload{},
		},
		{
			name:       "non-nil payload",
			statusCode: http.StatusOK,
			payload:    &MockPayload{Name: "John Doe", Email: "john@example.com"},
			expected:   MockPayload{Name: "John Doe", Email: "john@example.com"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			err := httpjson.Success(w, tc.statusCode, tc.payload)
			assert.NoError(t, err)

			// Check the response status code
			assert.Equal(t, tc.statusCode, w.Code)

			if tc.payload != nil {
				// Unmarshal the response payload
				var responsePayload MockPayload
				err = json.Unmarshal(w.Body.Bytes(), &responsePayload)
				assert.NoError(t, err)

				// Check the response payload
				assert.Equal(t, tc.expected, responsePayload)
			}
		})
	}
}

func TestError(t *testing.T) {
	tests := []struct {
		name   string
		err    error
		status int
	}{
		{
			name:   "not found error",
			err:    exceptions.ErrNotFound,
			status: http.StatusNotFound,
		},
		{
			name:   "internal server error",
			err:    nil,
			status: http.StatusInternalServerError,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			err := httpjson.Error(w, tc.err)
			assert.NoError(t, err)

			// Check the response status code
			assert.Equal(t, tc.status, w.Code)

			// Unmarshal the response payload
			var httpError httpexceptions.HTTPError
			err = json.Unmarshal(w.Body.Bytes(), &httpError)
			assert.NoError(t, err)

			// Check the error message
			assert.Equal(t, http.StatusText(tc.status), httpError.Error.Message)
		})
	}
}
