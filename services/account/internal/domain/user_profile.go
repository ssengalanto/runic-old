package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/ssengalanto/runic/pkg/validator"
)

// UserProfile - account's user profile entity.
type UserProfile struct {
	ID          uuid.UUID `json:"id" validate:"required"`
	UserID      uuid.UUID `json:"userId" validate:"required"`
	FirstName   string    `json:"firstName,omitempty" validate:"required"`
	LastName    string    `json:"lastName,omitempty" validate:"required"`
	DateOfBirth time.Time `json:"dateOfBirth" validate:"required"`
	Avatar      string    `json:"avatar,omitempty" validate:"omitempty,url"`
	Bio         string    `json:"bio,omitempty"`
}

// UpdateProfileInput contains the fields for updating user profile.
type UpdateProfileInput struct {
	FirstName   *string    `validate:"omitempty,nz"`
	LastName    *string    `validate:"omitempty,nz"`
	Avatar      *string    `validate:"omitempty,url"`
	Bio         *string    `validate:"omitempty,nz"`
	DateOfBirth *time.Time `validate:"omitempty,nz"`
}

// NewUserProfile creates a new user profile entity with the provided information.
// It performs validation on the inputs to ensure data integrity and returns an error if any validation fails.
func NewUserProfile(userID uuid.UUID, firstName, lastName string, dateOfBirth time.Time) (UserProfile, error) {
	profile := UserProfile{
		ID:          uuid.New(),
		UserID:      userID,
		FirstName:   firstName,
		LastName:    lastName,
		DateOfBirth: dateOfBirth,
	}

	err := profile.IsValid()
	if err != nil {
		return UserProfile{}, err
	}

	return profile, nil
}

// Update partially updates the user profile based on the provided UpdateProfileInput.
// It validates the request and returns an error if it is invalid.
func (u *UserProfile) Update(in UpdateProfileInput) error {
	err := validator.Struct(in)
	if err != nil {
		return err
	}

	if in.FirstName != nil {
		u.FirstName = *in.FirstName
	}

	if in.LastName != nil {
		u.LastName = *in.LastName
	}

	if in.Avatar != nil {
		u.Avatar = *in.Avatar
	}

	if in.Bio != nil {
		u.Bio = *in.Bio
	}

	if in.DateOfBirth != nil {
		u.DateOfBirth = *in.DateOfBirth
	}

	return nil
}

// IsValid checks the validity of the user profile entity.
func (u *UserProfile) IsValid() error {
	err := validator.Struct(u)
	if err != nil {
		return err
	}

	return err
}
