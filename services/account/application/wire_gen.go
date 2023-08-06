// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package application

import (
	"github.com/ssengalanto/runic/services/account/internal/features/get_account_details"
)

// Injectors from wire.go:

// Init initializes the App by using the dependency injection framework, Wire.
// It builds the registry of dependencies required by the App and panics if the build fails.
func Init() (*App, func()) {
	config := configFunc()
	log := loggerFunc(config)
	db, cleanup := pgsqlEngineFunc(config, log)
	service := getaccountdetails.NewService(log, db)
	handler := getaccountdetails.NewHandler(log, service)
	mediator := registerMediatorHandlers(log, handler)
	mux := registerHTTPControllers(log, mediator, config)
	server := httpServerFunc(config, mux)
	app := New(log, config, server)
	return app, func() {
		cleanup()
	}
}