package getaccountdetails

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/ssengalanto/runic/pkg/exceptions"
	"github.com/ssengalanto/runic/pkg/interfaces"
	"github.com/ssengalanto/runic/pkg/log"
	"github.com/ssengalanto/runic/services/account/internal/data"
	"github.com/ssengalanto/runic/services/account/internal/domain"
)

// Service represents a service that interacts with the database
// and provides methods for retrieving account user information.
type Service struct {
	slog interfaces.Logger
	db   *sqlx.DB
}

// NewService creates a new instance of the Service.
func NewService(slog interfaces.Logger, db *sqlx.DB) *Service {
	return &Service{
		slog: slog,
		db:   db,
	}
}

// GetAccountUser retrieves the account user details by the specified ID.
func (s *Service) GetAccountUser(ctx context.Context, id uuid.UUID) (domain.User, error) {
	model := data.User{}
	empty := domain.User{}

	query := `
		SELECT
			* 
		FROM account.user
		WHERE account.user.id = $1;
	`

	stmt, err := s.db.PreparexContext(ctx, query)
	if err != nil {
		return empty, err
	}
	defer stmt.Close()

	row := stmt.QueryRowxContext(ctx, id)
	err = row.StructScan(&model)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.slog.Error(
				fmt.Sprintf("no record found for account.user with id of `%s`", id),
				log.Field("error", err),
			)
			return empty, fmt.Errorf("%w: account.user with id of `%s`", exceptions.ErrNotFound, id)
		}
		s.slog.Error(
			fmt.Sprintf("account.user record with id of `%s` retrieval failed", id),
			log.Field("error", err),
		)
		return empty, err
	}

	return model.ToEntity(), nil
}
