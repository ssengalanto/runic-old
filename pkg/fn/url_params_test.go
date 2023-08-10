package fn_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ssengalanto/runic/pkg/fn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockHTTPWriter is a mock implementation of http.ResponseWriter for testing.
type MockHTTPWriter struct {
	mock.Mock
}

func (m *MockHTTPWriter) Header() http.Header {
	args := m.Called()
	return args.Get(0).(http.Header)
}

func (m *MockHTTPWriter) Write(data []byte) (int, error) {
	args := m.Called(data)
	return args.Int(0), args.Error(1)
}

func (m *MockHTTPWriter) WriteHeader(code int) {
	m.Called(code)
}

func TestParseUUIDFromURLParam(t *testing.T) {
	ctx := context.Background()
	r := chi.NewRouter()
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := fn.ParseUUIDFromURLParam("id", w, r)
		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, id)
	})

	url := fmt.Sprintf("/%s", uuid.New())
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)
}

func TestParseUUIDFromURLParam_InvalidUUID(t *testing.T) {
	ctx := context.Background()
	r := chi.NewRouter()
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		fn.ParseUUIDFromURLParam("id", w, r) //nolint:errcheck //intentional
	})

	url := fmt.Sprintf("/%s", "invalid-uuid")
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)
}
