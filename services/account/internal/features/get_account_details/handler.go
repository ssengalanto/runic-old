package getaccountdetails

import (
	"context"
	"fmt"

	"github.com/ssengalanto/runic/pkg/interfaces"
)

type Handler struct {
	slog interfaces.Logger
	svc  ServiceContract
}

func NewHandler(
	slog interfaces.Logger,
	svc ServiceContract,
) *Handler {
	return &Handler{
		slog: slog,
		svc:  svc,
	}
}

func (h *Handler) Name() string {
	return fmt.Sprintf("%T", &Query{})
}

func (h *Handler) Handle(
	ctx context.Context,
	request any,
) (any, error) {
	q := request.(*Query) //nolint:errcheck //intentional panic

	user, err := h.svc.GetAccountUser(ctx, q.ID)
	if err != nil {
		return nil, err
	}

	return Response{
		Data: AccountUserResponse{
			ID:          user.ID,
			Email:       user.Email.String(),
			Role:        user.Role.String(),
			Active:      user.Active,
			LastLoginAt: user.LastLoginAt,
		},
	}, nil
}
