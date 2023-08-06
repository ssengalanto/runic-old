package getaccountdetails

import "github.com/google/uuid"

// Query represents a query for retrieving account details.
type Query struct {
	ID uuid.UUID
}

// NewQuery creates a new Query instance with the provided ID.
func NewQuery(id uuid.UUID) *Query {
	return &Query{ID: id}
}
