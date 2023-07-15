package pgsql

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
)

var (
	// ErrConnectionFailed represents an error that is returned when a PostgreSQL database connection attempt fails.
	ErrConnectionFailed = fmt.Errorf("database connection failed")
)

// ErrorCode returns the PostgreSQL error code of the given error.
func ErrorCode(err error) string {
	var pgErr *pgconn.PgError

	ok := errors.As(err, &pgErr)
	if !ok {
		return ""
	}

	return pgErr.Code
}
