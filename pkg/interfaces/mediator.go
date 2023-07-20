package interfaces

import (
	"context"

	"github.com/ssengalanto/runic/pkg/mediator"
)

type Mediator interface {
	RegisterRequestHandler(handler mediator.RequestHandler) error
	RegisterNotificationHandler(handler mediator.NotificationHandler) error
	RegisterPipelineBehaviour(behaviour mediator.PipelineBehaviour) error
	Send(ctx context.Context, request any) (any, error)
	Publish(ctx context.Context, request any) error
}
