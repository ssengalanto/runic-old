package fn

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ssengalanto/runic/pkg/exceptions"
	"github.com/ssengalanto/runic/pkg/http/json"
)

// ParseUUIDFromURLParam parses a UUID from a URL parameter and handles any errors.
func ParseUUIDFromURLParam(key string, w http.ResponseWriter, r *http.Request) (uuid.UUID, error) {
	v := chi.URLParam(r, key)

	id, err := uuid.Parse(v)
	if err != nil {
		json.Error(w, fmt.Errorf("%w: %s param is not a valid uuid", exceptions.ErrInvalid, key))
		return uuid.Nil, err
	}

	return id, nil
}
