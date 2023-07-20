package mediator

import "fmt"

var (
	// ErrRequestHandlerAlreadyExists is returned when a request handler already exists in the registry.
	ErrRequestHandlerAlreadyExists = fmt.Errorf("request handler already exists")

	// ErrNotificationHandlerAlreadyExists is returned when a notification handler already exists in the registry.
	ErrNotificationHandlerAlreadyExists = fmt.Errorf("notification handler already exists")

	// ErrPipelineBehaviourAlreadyExists is returned when a pipeline behaviour already exists in the registry.
	ErrPipelineBehaviourAlreadyExists = fmt.Errorf("pipeline behaviour already exists")

	// ErrRequestHandlerNotFound is returned when a request handler doesn't exist in the registry.
	ErrRequestHandlerNotFound = fmt.Errorf("request handler not found")

	// ErrNotificationHandlerNotFound is returned when a notification handler doesn't exist in the registry.
	ErrNotificationHandlerNotFound = fmt.Errorf("notification handler not found")
)
