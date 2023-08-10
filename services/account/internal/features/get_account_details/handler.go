package get_account_details

import (
	"context"
	"fmt"

	"github.com/ssengalanto/runic/pkg/interfaces"
)

// Handler represents a request handler that processes incoming queries and returns responses.
type Handler struct {
	slog interfaces.Logger
	svc  ServiceContract
}

// NewHandler creates a new instance of the Handler.
func NewHandler(
	slog interfaces.Logger,
	svc ServiceContract,
) *Handler {
	return &Handler{
		slog: slog,
		svc:  svc,
	}
}

// Name returns the name of the handler, which is derived from the type of the associated Query struct.
func (h *Handler) Name() string {
	return fmt.Sprintf("%T", &Query{})
}

// Handle processes the incoming request and returns the corresponding response.
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
			Profile: AccountUserProfileResponse{
				ID:          user.Profile.ID,
				UserID:      user.Profile.UserID,
				FirstName:   user.Profile.FirstName,
				LastName:    user.Profile.LastName,
				DateOfBirth: user.Profile.DateOfBirth,
				Avatar:      user.Profile.Avatar,
				Bio:         user.Profile.Bio,
			},
		},
	}, nil
}
