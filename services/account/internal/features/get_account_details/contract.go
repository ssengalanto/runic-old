package getaccountdetails

import (
	"context"

	"github.com/google/uuid"
	"github.com/ssengalanto/runic/services/account/internal/domain"
)

type ServiceContract interface {
	GetAccountUser(ctx context.Context, id uuid.UUID) (domain.User, error)
}
