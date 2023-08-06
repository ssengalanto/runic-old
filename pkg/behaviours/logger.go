package behaviours

import (
	"context"
	"fmt"

	"github.com/ssengalanto/runic/pkg/interfaces"
	"github.com/ssengalanto/runic/pkg/log"
	"github.com/ssengalanto/runic/pkg/mediator"
)

// LoggerBehaviour represents a struct that encapsulates a logger interface, allowing logging capabilities
// for requests handling.
type LoggerBehaviour struct {
	slog interfaces.Logger
}

// NewLoggerBehaviour constructs a new instance of LoggerBehaviour and returns a pointer to it.
func NewLoggerBehaviour(slog interfaces.Logger) *LoggerBehaviour {
	return &LoggerBehaviour{
		slog: slog,
	}
}

// Handle function executes a middleware function to handle an incoming request, logs the request information
// using the logger interface and returns the response and an error, if any.
func (l *LoggerBehaviour) Handle(
	_ context.Context,
	request any,
	next mediator.RequestHandlerFunc,
) (any, error) {
	l.slog.Info(fmt.Sprintf("executing %T", request), log.Field("request", request))

	res, err := next()
	if err != nil {
		return nil, err
	}

	return res, nil
}
