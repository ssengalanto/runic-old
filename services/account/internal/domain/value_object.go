package domain

import (
	"fmt"

	"github.com/ssengalanto/runic/pkg/exceptions"
	"github.com/ssengalanto/runic/pkg/validator"
	"golang.org/x/crypto/bcrypt"
)

// Role value object.
type Role string

const (
	RoleAdmin     Role = "admin"
	RoleModerator Role = "moderator"
	RoleUser      Role = "user"
)

// IsValid checks the validity of the user role.
func (r Role) IsValid() (bool, error) {
	switch r {
	case RoleAdmin, RoleModerator, RoleUser:
		return true, nil
	default:
		return false, fmt.Errorf(
			"%w: %s %v",
			exceptions.ErrInvalid,
			"role must be one of the following:",
			[]Role{RoleAdmin, RoleModerator, RoleUser},
		)
	}
}

// Update checks the validity of the user role and updates its value.
func (r Role) Update(s string) (Role, error) {
	role := Role(s)
	if ok, err := role.IsValid(); !ok {
		return "", err
	}

	return role, nil
}

// String converts Role to type string.
func (r Role) String() string {
	return string(r)
}

// Email value object.
type Email string

// IsValid checks the validity of the email address.
func (e Email) IsValid() (bool, error) {
	err := validator.Var(e, "email,required")
	if err != nil {
		return false, err
	}

	return true, nil
}

// Update checks the validity of the email and updates its value.
func (e Email) Update(s string) (Email, error) {
	email := Email(s)
	if ok, err := email.IsValid(); !ok {
		return "", err
	}

	return email, nil
}

// String converts Email to type string.
func (e Email) String() string {
	return string(e)
}

// Password value object.
type Password string

// IsValid checks the validity of the password.
func (p Password) IsValid() (bool, error) {
	err := validator.Var(p, "min=10,required")
	if err != nil {
		return false, err
	}

	return true, nil
}

// Update checks the validity of the password and updates its value.
func (p Password) Update(s string) (Password, error) {
	password := Password(s)
	if ok, err := password.IsValid(); !ok {
		return "", err
	}

	return password, nil
}

// Hash generates a bcrypt hash for the password.
func (p Password) Hash() (Password, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return Password(hashed), nil
}

// Check compares the hashed password with the provided password.
func (p Password) Check(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(p), []byte(password))
}

// String converts Password to type string.
func (p Password) String() string {
	return string(p)
}
