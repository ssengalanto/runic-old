// Package pgsql is a concurrency-safe connection pool for postgres database.
// Its using sqlx internally with pgx sql driver.
package pgsql

import (
	"context"
	"fmt"
	"net"

	"github.com/jmoiron/sqlx"
	"github.com/ssengalanto/runic/pkg/constants"
)

const (
	// The driver constant represents the SQL driver used for connecting to PostgreSQL databases.
	driver = "pgx"
)

// NewConnection creates a new PostgreSQL database connection and returns a pointer to the database instance.
// It takes the following parameters:
// - user: The username for the database connection.
// - pwd: The password for the database connection.
// - host: The host address of the PostgreSQL server.
// - port: The port number of the PostgreSQL server.
// - dbn: The name of the PostgreSQL database.
// - qp: The optional query parameters for the database connection.
// The function returns a pointer to sqlx.DB and an error if any.
func NewConnection(user, pwd, host, port, dbn, qp string) (*sqlx.DB, error) {
	hp := net.JoinHostPort(host, port)
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?%s", user, pwd, hp, dbn, qp)

	ctx, cancel := context.WithTimeout(context.Background(), constants.Timeout)
	defer cancel()

	db, err := sqlx.ConnectContext(ctx, driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrConnectionFailed, err)
	}

	if err = db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("%w: %s", ErrConnectionFailed, err)
	}

	return db, nil
}
