package get_account_details

import (
	"github.com/google/uuid"
)

// Query represents a request data structure used for querying account details.
type Query struct {
	ID uuid.UUID
}

// NewQuery creates a new instance of the Query.
func NewQuery(id uuid.UUID) *Query {
	return &Query{ID: id}
}
