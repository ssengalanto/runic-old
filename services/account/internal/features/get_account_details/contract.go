package getaccountdetails

import (
	"context"

	"github.com/google/uuid"
	"github.com/ssengalanto/runic/services/account/internal/domain"
)

// ServiceContract is an interface that defines the methods to be implemented by a service
// for retrieving account user information.
type ServiceContract interface {
	GetAccountUser(ctx context.Context, id uuid.UUID) (domain.User, error)
}
