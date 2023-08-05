package data

import (
	"time"

	"github.com/google/uuid"
	"github.com/ssengalanto/runic/services/account/internal/domain"
)

// User pgsql model.
type User struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Email       string    `json:"email" db:"email"`
	Password    string    `json:"password" db:"password"`
	Role        string    `json:"role" db:"role"`
	Active      bool      `json:"active" db:"active"`
	LastLoginAt time.Time `json:"lastLoginAt,omitempty"  db:"last_login_at"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

// ToEntity transforms the account's user model to account's user entity.
func (u User) ToEntity() domain.User {
	return domain.User{
		ID:          u.ID,
		Email:       domain.Email(u.Email),
		Password:    domain.Password(u.Password),
		Role:        domain.Role(u.Role),
		Active:      u.Active,
		LastLoginAt: u.LastLoginAt,
	}
}
