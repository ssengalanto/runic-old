package mediator

import (
	"context"
	"fmt"
	"reflect"
)

// RequestHandler defines the interface for handling request messages.
type RequestHandler interface {
	// Name returns the unique name of the request handler.
	Name() string
	// Handle processes the incoming request and returns the response or an error.
	Handle(ctx context.Context, request any) (any, error)
}

// NotificationHandler defines the interface for handling notification messages.
type NotificationHandler interface {
	// Name returns the unique name of the notification handler.
	Name() string
	// Handle processes the incoming notification.
	Handle(ctx context.Context, notification any) error
}

// PipelineBehaviour defines the interface for extending the request handling pipeline.
type PipelineBehaviour interface {
	// Handle processes the request with additional behavior and forwards it to the next handler.
	Handle(ctx context.Context, request any, next RequestHandlerFunc) (any, error)
}

// RequestHandlerFunc represents the function signature of a request handler.
type RequestHandlerFunc func() (any, error)

// Mediator is a mediator that registers and dispatches requests and notifications.
type Mediator struct {
	requestHandlerRegistry      map[string]RequestHandler
	notificationHandlerRegistry map[string]NotificationHandler
	pipelineBehaviourRegistry   []PipelineBehaviour
}

// New creates a new Mediator instance.
func New() *Mediator {
	return &Mediator{
		requestHandlerRegistry:      make(map[string]RequestHandler),
		notificationHandlerRegistry: make(map[string]NotificationHandler),
		pipelineBehaviourRegistry:   []PipelineBehaviour{},
	}
}

// RegisterRequestHandler registers a RequestHandler in the mediator's registry.
func (m *Mediator) RegisterRequestHandler(handler RequestHandler) error {
	hn := handler.Name()

	if _, ok := m.requestHandlerRegistry[hn]; ok {
		return fmt.Errorf("%w: %s", ErrRequestHandlerAlreadyExists, hn)
	}

	m.requestHandlerRegistry[hn] = handler
	return nil
}

// RegisterNotificationHandler registers a NotificationHandler in the mediator's registry.
func (m *Mediator) RegisterNotificationHandler(handler NotificationHandler) error {
	hn := handler.Name()

	if _, ok := m.notificationHandlerRegistry[hn]; ok {
		return fmt.Errorf("%w: %s", ErrNotificationHandlerAlreadyExists, hn)
	}

	m.notificationHandlerRegistry[hn] = handler
	return nil
}

// RegisterPipelineBehaviour registers a PipelineBehaviour in the mediator's registry.
func (m *Mediator) RegisterPipelineBehaviour(behaviour PipelineBehaviour) error {
	bt := reflect.TypeOf(behaviour)

	if m.existsPipeType(bt) {
		return fmt.Errorf("%w: %s", ErrPipelineBehaviourAlreadyExists, bt)
	}

	m.pipelineBehaviourRegistry = append(m.pipelineBehaviourRegistry, behaviour)
	return nil
}

// Send sends the request to its corresponding RequestHandler, processing it through any registered pipeline behaviors.
func (m *Mediator) Send(ctx context.Context, request any) (any, error) {
	rt := reflect.TypeOf(request).String()

	handler, ok := m.requestHandlerRegistry[rt]
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrRequestHandlerNotFound, rt)
	}

	if len(m.pipelineBehaviourRegistry) > 0 {
		var lastHandler RequestHandlerFunc = func() (any, error) {
			return handler.Handle(ctx, request)
		}

		var aggregateResult = lastHandler
		for _, pipe := range m.pipelineBehaviourRegistry {
			pipeValue := pipe
			nextValue := aggregateResult

			aggregateResult = func() (any, error) {
				return pipeValue.Handle(ctx, request, nextValue)
			}
		}

		response, err := aggregateResult()
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	response, err := handler.Handle(ctx, request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Publish publishes the notification event to its corresponding NotificationHandler.
func (m *Mediator) Publish(ctx context.Context, notification any) error {
	rt := reflect.TypeOf(notification).String()

	handler, ok := m.notificationHandlerRegistry[rt]
	if !ok {
		return fmt.Errorf("%w: %s", ErrNotificationHandlerNotFound, rt)
	}

	err := handler.Handle(ctx, notification)
	if err != nil {
		return err
	}

	return nil
}

// existsPipeType checks if a pipeline behavior exists in the mediator's registry.
func (m *Mediator) existsPipeType(p reflect.Type) bool {
	for _, pipe := range m.pipelineBehaviourRegistry {
		if reflect.TypeOf(pipe) == p {
			return true
		}
	}
	return false
}
