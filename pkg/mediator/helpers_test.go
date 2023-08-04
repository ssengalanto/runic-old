package mediator_test

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/ssengalanto/runic/pkg/mediator"
)

type CommandRequest struct {
	Data string
}

type CommandResponse struct {
	Data string
}

type CommandHandler struct{}

func (c *CommandHandler) Name() string {
	return fmt.Sprintf("%T", &CommandRequest{})
}

func (c *CommandHandler) Handle(_ context.Context, request any) (any, error) {
	empty := CommandResponse{}

	req, ok := request.(*CommandRequest)
	if !ok {
		return empty, errors.New("invalid request")
	}

	fmt.Printf("%s executed", c.Name())
	return CommandResponse{Data: req.Data}, nil
}

type NotificationRequest struct {
	Data string
}

type NotificationHandler struct{}

func (n *NotificationHandler) Name() string {
	return fmt.Sprintf("%T", &NotificationRequest{})
}

func (n *NotificationHandler) Handle(_ context.Context, notification any) error {
	_, ok := notification.(*NotificationRequest)
	if !ok {
		return errors.New("invalid notification")
	}

	fmt.Printf("%s executed", n.Name())
	return nil
}

type PipelineBehaviourHandler struct{}

func (p *PipelineBehaviourHandler) Handle(
	_ context.Context,
	request any,
	next mediator.RequestHandlerFunc,
) (any, error) {
	log.Printf("request: %v", request)

	res, err := next()
	if err != nil {
		return nil, err
	}

	return res, nil
}
