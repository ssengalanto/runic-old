//nolint:unused,gochecknoglobals //used internally
package application

import (
	"fmt"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/ssengalanto/runic/pkg/http/server"
	"github.com/ssengalanto/runic/pkg/interfaces"
	"github.com/ssengalanto/runic/pkg/log"
	"github.com/ssengalanto/runic/pkg/pgsql"
	"github.com/ssengalanto/runic/services/account/internal/config"
)

var registry = wire.NewSet(
	configFunc,
	loggerSet,
	pgsqlEngineFunc,
	mediatorHandlerSet,
	httpHandlerSet,
	httpServerFunc,
	appSet,
)

// configFunc creates a new instance of config.Config struct.
func configFunc() *config.Config {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	return cfg
}

// loggerFunc creates a new instance of the log.Logger struct.
func loggerFunc(cfg *config.Config) *log.Log {
	slog, err := log.New(cfg.App.Env)
	if err != nil {
		panic(err)
	}

	return slog
}

var loggerSet = wire.NewSet(loggerFunc, wire.Bind(new(interfaces.Logger), new(*log.Log)))

// pgsqlEngineFunc creates a new PostgreSQL database connection using the provided configuration.
func pgsqlEngineFunc(cfg *config.Config, slog interfaces.Logger) (*sqlx.DB, func()) {
	pg := cfg.PGSQL

	db, err := pgsql.NewConnection(
		pg.Username,
		pg.Password,
		pg.Host,
		strconv.Itoa(pg.Port),
		pg.DBName,
		pg.QueryParams,
	)
	if err != nil {
		slog.Error("failed to connect to database", log.Field("error", err))
		panic(err)
	}

	slog.Info(
		fmt.Sprintf("connected to database %s on %s with %s user account.", pg.DBName, pg.Host, pg.Username),
	)

	return db, func() { db.Close() }
}

// httpServerFunc creates a new instance of HTTP server with the given configuration and multiplexer.
func httpServerFunc(cfg *config.Config, mux *chi.Mux) *server.Server {
	return server.New(cfg.HTTP.Port, mux)
}
