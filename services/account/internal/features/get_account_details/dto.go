package get_account_details

import (
	"time"

	"github.com/google/uuid"
)

// Response is a struct representing the response format for the GetAccountDetails API.
type Response struct {
	Data AccountUserResponse `json:"data"`
}

// AccountUserResponse is a struct representing the user details in the response for GetAccountDetails API.
type AccountUserResponse struct {
	ID          uuid.UUID                  `json:"id"`
	Email       string                     `json:"email"`
	Role        string                     `json:"role"`
	Active      bool                       `json:"active"`
	LastLoginAt time.Time                  `json:"lastLoginAt,omitempty"`
	Profile     AccountUserProfileResponse `json:"profile"`
} // @name AccountUserResponse

// AccountUserProfileResponse is a struct representing the user profile details
// in the response for GetAccountDetails API.
type AccountUserProfileResponse struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"userId"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Avatar      string    `json:"avatar,omitempty"`
	Bio         string    `json:"bio,omitempty"`
} // @name AccountUserProfileResponse
