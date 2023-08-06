package application

import (
	"fmt"

	"github.com/google/wire"
	"github.com/ssengalanto/runic/pkg/http/server"
	"github.com/ssengalanto/runic/pkg/interfaces"
	"github.com/ssengalanto/runic/services/account/internal/config"
)

var appSet = wire.NewSet(New) //nolint:unused,gochecknoglobals //used in wire_registry.go

// App represents the main application instance that is responsible for managing and running the application.
type App struct {
	slog interfaces.Logger
	cfg  *config.Config
	svr  *server.Server
}

// New creates a new instance of the App by initializing its dependencies.
func New(slog interfaces.Logger, config *config.Config, server *server.Server) *App {
	return &App{
		slog: slog,
		cfg:  config,
		svr:  server,
	}
}

// Run starts the server by binding it to a port and listening for incoming requests.
func (a *App) Run() {
	a.slog.Info(fmt.Sprintf("🚀 server has started on port %d", a.cfg.HTTP.Port))

	err := a.svr.Start()
	if err != nil {
		a.slog.Info(fmt.Sprintf("shutting down http server on port %d", a.cfg.HTTP.Port))
	}
}
