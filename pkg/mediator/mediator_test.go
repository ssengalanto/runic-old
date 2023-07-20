package mediator_test

import (
	"context"
	"testing"

	"github.com/ssengalanto/runic/pkg/mediator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Run("creates a new mediator instance", func(t *testing.T) {
		m := mediator.New()
		assert.NotNil(t, m)
	})
}

func TestMediator_RegisterRequestHandler(t *testing.T) {
	t.Run("registers all request handlers successfully", func(t *testing.T) {
		m := mediator.New()

		hdlr := &CommandHandler{}
		err := m.RegisterRequestHandler(hdlr)
		require.NoError(t, err)
	})
	t.Run("return an error when trying to register an already existing request handler", func(t *testing.T) {
		m := mediator.New()

		hdlr1 := &CommandHandler{}
		err := m.RegisterRequestHandler(hdlr1)
		require.NoError(t, err)

		hdlr2 := &CommandHandler{}
		err = m.RegisterRequestHandler(hdlr2)
		require.Error(t, err)
	})
}

func TestMediator_RegisterNotificationHandler(t *testing.T) {
	t.Run("registers the notification handler successfully", func(t *testing.T) {
		m := mediator.New()

		hdlr := &NotificationHandler{}
		err := m.RegisterNotificationHandler(hdlr)
		require.NoError(t, err)
	})
	t.Run("return an error when trying to register an already existing notification handler", func(t *testing.T) {
		m := mediator.New()

		hdlr1 := &NotificationHandler{}
		err := m.RegisterNotificationHandler(hdlr1)
		require.NoError(t, err)

		hdlr2 := &NotificationHandler{}
		err = m.RegisterNotificationHandler(hdlr2)
		require.Error(t, err)
	})
}

func TestMediator_RegisterPipelineBehaviour(t *testing.T) {
	t.Run("registers the pipeline behaviour successfully", func(t *testing.T) {
		m := mediator.New()

		pb := &PipelineBehaviourHandler{}
		err := m.RegisterPipelineBehaviour(pb)
		require.NoError(t, err)
	})
	t.Run("return an error when trying to register an already existing pipeline behaviour",
		func(t *testing.T) {
			m := mediator.New()

			pb1 := &PipelineBehaviourHandler{}
			err := m.RegisterPipelineBehaviour(pb1)
			require.NoError(t, err)

			pb2 := &PipelineBehaviourHandler{}
			err = m.RegisterPipelineBehaviour(pb2)
			require.Error(t, err)
		})
}

func TestMediator_Send(t *testing.T) {
	t.Run("executes the request handler",
		func(t *testing.T) {
			m := mediator.New()

			hdlr := &CommandHandler{}
			err := m.RegisterRequestHandler(hdlr)
			require.NoError(t, err)

			_, err = m.Send(context.Background(), &CommandRequest{})
			require.NoError(t, err)
		})
	t.Run("executes the pipeline behaviours",
		func(t *testing.T) {
			m := mediator.New()

			pb := &PipelineBehaviourHandler{}
			err := m.RegisterPipelineBehaviour(pb)
			require.NoError(t, err)

			hdlr := &CommandHandler{}
			err = m.RegisterRequestHandler(hdlr)
			require.NoError(t, err)

			_, err = m.Send(context.Background(), &CommandRequest{})
			require.NoError(t, err)
		})
	t.Run("return an error if there are no registered request handlers in the registry",
		func(t *testing.T) {
			m := mediator.New()

			_, err := m.Send(context.Background(), &CommandRequest{})
			require.Error(t, err)
		})
}

func TestMediator_Publish(t *testing.T) {
	t.Run("executes the notification handler",
		func(t *testing.T) {
			m := mediator.New()

			hdlr := &NotificationHandler{}
			err := m.RegisterNotificationHandler(hdlr)
			require.NoError(t, err)

			err = m.Publish(context.Background(), &NotificationRequest{})
			require.NoError(t, err)
		})
	t.Run("return an error if there are no registered notification handlers in the registry",
		func(t *testing.T) {
			m := mediator.New()

			err := m.Publish(context.Background(), &NotificationRequest{})
			require.Error(t, err)
		})
}
