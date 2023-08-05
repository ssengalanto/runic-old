package data

import (
	"time"

	"github.com/google/uuid"
	"github.com/ssengalanto/runic/services/account/internal/domain"
)

// UserProfile pgsql model.
type UserProfile struct {
	ID          uuid.UUID `json:"id" db:"id"`
	UserID      uuid.UUID `json:"userId" db:"user_id"`
	FirstName   string    `json:"firstName" db:"first_name"`
	LastName    string    `json:"lastName" db:"last_name"`
	DateOfBirth time.Time `json:"dateOfBirth" db:"date_of_birth"`
	Avatar      string    `json:"avatar,omitempty" db:"avatar"`
	Bio         string    `json:"bio,omitempty" db:"bio"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

// ToEntity transforms the account's user profile model to account's user profile entity.
func (u UserProfile) ToEntity() domain.UserProfile {
	return domain.UserProfile{
		ID:          u.ID,
		UserID:      u.UserID,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		DateOfBirth: u.DateOfBirth,
		Avatar:      u.Avatar,
		Bio:         u.Bio,
	}
}
