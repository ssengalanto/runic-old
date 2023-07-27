package domain_test

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/ssengalanto/runic/services/account/internal/domain"
	"github.com/ssengalanto/runic/services/account/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewUserProfile(t *testing.T) {
	type input struct {
		userID      uuid.UUID
		firstName   string
		lastName    string
		dateOfBirth time.Time
	}

	testCases := []struct {
		name   string
		input  input
		assert func(t *testing.T, err error)
	}{
		{
			name: "new instance creation success",
			input: input{
				userID:      uuid.New(),
				firstName:   gofakeit.FirstName(),
				lastName:    gofakeit.LastName(),
				dateOfBirth: gofakeit.Date(),
			},
			assert: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		{
			name: "new instance creation failed due to zeroed date of birth",
			input: input{
				userID:      uuid.New(),
				firstName:   gofakeit.FirstName(),
				lastName:    gofakeit.LastName(),
				dateOfBirth: time.Time{},
			},
			assert: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
		{
			name: "new instance creation failed due to zeroed uuid",
			input: input{
				userID:      uuid.Nil,
				firstName:   gofakeit.FirstName(),
				lastName:    gofakeit.LastName(),
				dateOfBirth: gofakeit.Date(),
			},
			assert: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := domain.NewUserProfile(
				tc.input.userID,
				tc.input.firstName,
				tc.input.lastName,
				tc.input.dateOfBirth,
			)
			tc.assert(t, err)
		})
	}
}

func TestUserProfile_Update(t *testing.T) {
	update := domain.UserProfile{
		FirstName:   gofakeit.FirstName(),
		LastName:    gofakeit.LastName(),
		DateOfBirth: gofakeit.Date(),
		Avatar:      gofakeit.URL(),
		Bio:         gofakeit.Word(),
	}

	emptyStr := ""

	testCases := []struct {
		name    string
		profile domain.UserProfile
		input   domain.UpdateProfileInput
		assert  func(t *testing.T, expected domain.UserProfile, actual domain.UserProfile, err error)
	}{
		{
			name:    "update user profile success",
			profile: mock.Profile(),
			input: domain.UpdateProfileInput{
				FirstName:   &update.FirstName,
				LastName:    &update.LastName,
				DateOfBirth: &update.DateOfBirth,
				Avatar:      &update.Avatar,
				Bio:         &update.Bio,
			},
			assert: func(t *testing.T, expected domain.UserProfile, actual domain.UserProfile, err error) {
				assert.Equal(t, expected, actual)
				require.NoError(t, err)
			},
		},
		{
			name:    "update user profile failed",
			profile: mock.Profile(),
			input: domain.UpdateProfileInput{
				FirstName:   &emptyStr,
				LastName:    &emptyStr,
				DateOfBirth: &time.Time{},
				Avatar:      &emptyStr,
				Bio:         &emptyStr,
			},
			assert: func(t *testing.T, expected domain.UserProfile, actual domain.UserProfile, err error) {
				assert.NotEqual(t, expected, actual)
				require.Error(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			profile := tc.profile
			err := profile.Update(tc.input)
			updateDetails := domain.UserProfile{
				ID:          profile.ID,
				UserID:      profile.UserID,
				FirstName:   *tc.input.FirstName,
				LastName:    *tc.input.LastName,
				DateOfBirth: *tc.input.DateOfBirth,
				Avatar:      *tc.input.Avatar,
				Bio:         *tc.input.Bio,
			}
			tc.assert(t,
				profile, updateDetails, err)
		})
	}
}
