package middlewares_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ssengalanto/runic/pkg/http/middlewares"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJSONContentType(t *testing.T) {
	// Create a sample request
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
	require.NoError(t, err)

	// Create a response recorder to record the middleware's behavior
	rr := httptest.NewRecorder()

	// Create a handler that will be wrapped by the middleware
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the "Content-Type" header is set correctly
		contentType := w.Header().Get("Content-Type")
		assert.Equal(t, "application/json", contentType)

		_, err = w.Write([]byte("Hello, world!"))
		require.NoError(t, err)
	})

	// Create an instance of the middleware, wrapping the handler
	middleware := middlewares.JSONContentType(handler)

	// Serve the request through the middleware
	middleware.ServeHTTP(rr, req)

	// Check if the response body is correct
	assert.Equal(t, "Hello, world!", rr.Body.String())
}
