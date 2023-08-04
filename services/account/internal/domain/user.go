package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ssengalanto/runic/pkg/exceptions"
	"github.com/ssengalanto/runic/pkg/validator"
)

type Role string

const (
	RoleAdmin     Role = "admin"
	RoleModerator Role = "moderator"
	RoleUser      Role = "user"
)

// User - account's user entity, serves as the aggregate root.
type User struct {
	ID          uuid.UUID    `json:"id" validate:"required"`
	Email       Email        `json:"email" validate:"required,email"`
	Password    Password     `json:"password" validate:"required,min=10"`
	Role        Role         `json:"role" validate:"required"`
	Active      bool         `json:"active" validate:"required"`
	LastLoginAt time.Time    `json:"lastLoginAt,omitempty"`
	Profile     *UserProfile `json:"profile,omitempty"`
}

// NewUser creates a new account entity with the provided information.
// It performs validation on the inputs to ensure data integrity and returns an error if any validation fails.
func NewUser(email, password string, role Role) (User, error) {
	user := User{
		ID:       uuid.New(),
		Email:    Email(email),
		Password: Password(password),
		Role:     role,
		Active:   true,
	}

	err := user.IsValid()
	if err != nil {
		return User{}, err
	}

	return user, nil
}

// IsActive returns the current active state of the user entity.
// It returns true if the user is active, otherwise false.
func (u *User) IsActive() bool {
	return u.Active
}

// Activate sets the active state of the user entity to `true`.
func (u *User) Activate() {
	u.Active = true
}

// Deactivate sets the active state of the user entity to `false`.
func (u *User) Deactivate() {
	u.Active = false
}

// RecordLastLogin sets the last login time of the user entity with the current timestamp.
func (u *User) RecordLastLogin() {
	u.LastLoginAt = time.Now()
}

// UpdateEmail checks the validity of the given email address and updates the email field.
// It returns an error if the email address is invalid or if there was a problem during the update.
func (u *User) UpdateEmail(s string) error {
	email, err := u.Email.Update(s)
	if err != nil {
		return err
	}

	u.Email = email
	return nil
}

// UpdatePassword checks the validity of the given password and updates the password field.
// It returns an error if the password is invalid or if there was a problem during the update.
func (u *User) UpdatePassword(s string) error {
	password, err := u.Password.Update(s)
	if err != nil {
		return err
	}

	u.Password = password
	return nil
}

// HashPassword hashes the User's password using the Hash method of the underlying Password field.
// It returns an error if there was a problem during hashing.
func (u *User) HashPassword() error {
	hashed, err := u.Password.Hash()
	if err != nil {
		return err
	}

	u.Password = hashed
	return nil
}

// CheckPassword checks whether the provided password matches the hashed password stored in the User's Password field.
// It returns true if the passwords match, otherwise false.
func (u *User) CheckPassword(password string) bool {
	err := u.Password.Check(password)
	return err == nil
}

// IsValid performs validation on the User and its fields to check if it is a valid User instance.
// It uses a validator to validate the User's fields based on predefined rules and tags.
// It returns an error if any validation error occurs, otherwise nil.
func (u *User) IsValid() error {
	err := u.validateRole()
	if err != nil {
		return err
	}

	err = validator.Struct(u)
	if err != nil {
		return err
	}

	return err
}

func (u *User) validateRole() error {
	switch u.Role {
	case RoleAdmin, RoleModerator, RoleUser:
		return nil
	default:
		return fmt.Errorf(
			"%w: %s %v",
			exceptions.ErrInvalid,
			"role must be one of the following:",
			[]Role{RoleAdmin, RoleModerator, RoleUser},
		)
	}
}

// UpdateProfile partially updates the user profile based on the provided UpdateProfileInput.
// It validates the request and returns an error if the input is invalid.
func (u *User) UpdateProfile(in UpdateProfileInput) error {
	err := u.Profile.Update(in)
	if err != nil {
		return err
	}

	return nil
}

// AggregateRoot aggregates user and user profile entities.
func AggregateRoot(user User, profile UserProfile) User {
	root := user
	root.Profile = &profile

	return root
}
