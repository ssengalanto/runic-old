package mux_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ssengalanto/runic/pkg/http/mux"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// Create a new request with a GET method and "/" URL
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	// Create a new response recorder to capture the response
	rec := httptest.NewRecorder()

	// Create a new router using the New function
	router := mux.New()

	// Register a route for the "/" URL pattern
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Serve the request using the router
	router.ServeHTTP(rec, req)

	// Verify the response status code is 200 OK
	assert.Equal(t, http.StatusOK, rec.Code)
}
