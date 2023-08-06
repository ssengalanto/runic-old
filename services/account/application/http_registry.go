package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
	"github.com/ssengalanto/runic/pkg/constants"
	"github.com/ssengalanto/runic/pkg/http/mux"
	"github.com/ssengalanto/runic/pkg/interfaces"
	"github.com/ssengalanto/runic/services/account/internal/config"

	getaccountdetails "github.com/ssengalanto/runic/services/account/internal/features/get_account_details"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

var httpHandlerSet = wire.NewSet(registerHTTPControllers) //nolint:unused,gochecknoglobals //used in wire_registry.go

// registerHTTPControllers creates a new HTTP route multiplexer and registers HTTP handlers.
func registerHTTPControllers(slog interfaces.Logger, mediator interfaces.Mediator, cfg *config.Config) *chi.Mux {
	router := mux.New()

	if cfg.App.Env != constants.ProdEnv {
		// register swagger docs endpoint
		router.Mount("/swagger/docs", httpSwagger.WrapHandler)
	}

	registerHTTPHandlers(slog, router, mediator)

	return router
}

// registerHTTPHandlers registers the HTTP handlers for the provided HTTP route multiplexer.
func registerHTTPHandlers(slog interfaces.Logger, mux *chi.Mux, mediator interfaces.Mediator) {
	mux.Route("/api", func(r chi.Router) {
		// GET /api/account/{id}
		getAccountDetails := getaccountdetails.NewController(slog, mediator)
		r.Get("/account/{id}", getAccountDetails.Handle)
	})
}
