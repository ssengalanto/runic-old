package getaccountdetails

import (
	"time"

	"github.com/google/uuid"
)

type Response struct {
	Data AccountUserResponse `json:"data"`
} // @name GetAccountDetailsResponse

type AccountUserResponse struct {
	ID          uuid.UUID `json:"id"`
	Email       string    `json:"email"`
	Role        string    `json:"role"`
	Active      bool      `json:"active"`
	LastLoginAt time.Time `json:"lastLoginAt,omitempty"`
} // @name AccountUserResponse
