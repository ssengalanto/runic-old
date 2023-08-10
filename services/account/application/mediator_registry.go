//nolint:unused,gochecknoglobals //used internally
package application

import (
	"github.com/google/wire"
	"github.com/ssengalanto/runic/pkg/behaviours"
	"github.com/ssengalanto/runic/pkg/interfaces"
	"github.com/ssengalanto/runic/pkg/log"
	"github.com/ssengalanto/runic/pkg/mediator"

	"github.com/ssengalanto/runic/services/account/internal/features/get_account_details"
)

var (
	// mediator-registry set.
	mediatorHandlerSet = wire.NewSet(
		registerMediatorHandlers,
		getAccountDetailsSet,
	)

	// mediator handler sets.
	getAccountDetailsSet = wire.NewSet(
		get_account_details.NewService,
		wire.Bind(new(get_account_details.ServiceContract), new(*get_account_details.Service)),
		get_account_details.NewHandler,
	)
)

// registerRequestHandlers registers all request handlers in the registry.
func registerMediatorHandlers(
	slog interfaces.Logger,
	getAccountDetailsHandler *get_account_details.Handler,
) interfaces.Mediator {
	m := mediator.New()

	err := m.RegisterRequestHandler(getAccountDetailsHandler)
	if err != nil {
		slog.Error("get_account_details.Handler registry failed", log.Field("error", err))
		panic(err)
	}

	// register pipeline behaviours
	registerPipelineBehaviours(slog, m)

	return m
}

// registerPipelineBehaviours registers all pipeline behaviour in the registry.
func registerPipelineBehaviours(slog interfaces.Logger, m interfaces.Mediator) {
	loggerBehaviour := behaviours.NewLoggerBehaviour(slog)
	err := m.RegisterPipelineBehaviour(loggerBehaviour)
	if err != nil {
		slog.Error("logger pipeline behaviour registry failed", log.Field("error", err))
		panic(err)
	}
}
