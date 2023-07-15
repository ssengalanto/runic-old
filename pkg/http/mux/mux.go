// Package mux provides an HTTP route multiplexer with sensible middlewares.
package mux

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/ssengalanto/runic/pkg/constants"
	"github.com/ssengalanto/runic/pkg/http/middlewares"
)

// New creates a new chi HTTP route multiplexer and adds several sensible middleware to it.
func New() *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Compress(constants.GzipCompressionLevel))
	mux.Use(cors.Handler(cors.Options{ // Enable CORS with more fine-grained configuration
		AllowedOrigins:   []string{"https://*", "http://*"}, // TODO: make it dynamic based on app_nv
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		MaxAge:           constants.CorsMaxAge,
		AllowCredentials: false,
	}))
	mux.Use(httprate.LimitByIP(constants.RateLimit, time.Minute))
	mux.Use(middleware.Timeout(constants.Timeout))
	mux.Use(middlewares.JSONContentType)
	mux.Use(middleware.Heartbeat("/heartbeat")) // TODO: create specific health check for each service

	return mux
}
