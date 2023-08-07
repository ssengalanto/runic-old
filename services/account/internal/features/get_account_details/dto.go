package getaccountdetails

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
	ID          uuid.UUID `json:"id"`
	Email       string    `json:"email"`
	Role        string    `json:"role"`
	Active      bool      `json:"active"`
	LastLoginAt time.Time `json:"lastLoginAt,omitempty"`
} // @name AccountUserResponse
