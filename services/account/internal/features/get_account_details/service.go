package get_account_details

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

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

// GetAccountUserModel represents the database structure for retrieving account details.
type GetAccountUserModel struct {
	data.User
	ID          uuid.UUID `db:"p_id"`
	UserID      uuid.UUID `db:"p_user_id"`
	FirstName   string    `db:"p_first_name"`
	LastName    string    `db:"p_last_name"`
	DateOfBirth time.Time `db:"p_date_of_birth"`
	Avatar      string    `db:"p_avatar"`
	Bio         string    `db:"p_bio"`
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
	model := GetAccountUserModel{}
	empty := domain.User{}

	query := `
		SELECT
			au.id,
			au.email,
			au.password,
			au.active,
			au.role,
			au.last_login_at,
			aup.id as p_id,
			aup.user_id as p_user_id,
			aup.first_name as p_first_name,
			aup.last_name as p_last_name,
			aup.avatar as p_avatar,
			aup.bio as p_bio,
			aup.date_of_birth as p_date_of_birth
		FROM account.user au
		JOIN account.user_profile aup
		ON aup.user_id = au.id
		WHERE au.id = $1;
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

	return s.mapToEntity(model), nil
}

// mapToEntity converts a GetAccountUserModel to a domain.User entity.
func (s *Service) mapToEntity(model GetAccountUserModel) domain.User {
	user := model.User.ToEntity()
	profile := domain.UserProfile{
		ID:          model.ID,
		UserID:      model.UserID,
		FirstName:   model.FirstName,
		LastName:    model.LastName,
		Avatar:      model.Avatar,
		Bio:         model.Bio,
		DateOfBirth: model.DateOfBirth,
	}

	return domain.AggregateRoot(user, profile)
}
