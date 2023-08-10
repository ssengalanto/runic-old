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
	ID          uuid.UUID                  `json:"id" example:"0b6ecded-fa9d-4b39-a309-9ef501de15f4"`
	Email       string                     `json:"email" example:"johndoe@example.com"`
	Role        string                     `json:"role" example:"admin"`
	Active      bool                       `json:"active" example:"true"`
	LastLoginAt time.Time                  `json:"lastLoginAt,omitempty" example:"2000-11-12T13:14:15Z"`
	Profile     AccountUserProfileResponse `json:"profile"`
} // @name AccountUserResponse

// AccountUserProfileResponse is a struct representing the user profile details
// in the response for GetAccountDetails API.
type AccountUserProfileResponse struct {
	ID          uuid.UUID `json:"id" example:"0b6ecded-fa9d-4b39-a309-9ef501de15f4"`
	UserID      uuid.UUID `json:"userId" example:"0b6ecded-fa9d-4b39-a309-9ef501de15f4"`
	FirstName   string    `json:"firstName" example:"John"`
	LastName    string    `json:"lastName" example:"Doe"`
	DateOfBirth time.Time `json:"dateOfBirth" example:"2000-11-12T13:14:15Z"`
	Avatar      string    `json:"avatar,omitempty" example:"https://avatar.com"`
	Bio         string    `json:"bio,omitempty" example:"Hi, I'm John Doe."`
} // @name AccountUserProfileResponse
