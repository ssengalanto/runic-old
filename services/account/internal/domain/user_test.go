package domain_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ssengalanto/runic/pkg/fn"
	"github.com/ssengalanto/runic/services/account/internal/domain"
	"github.com/ssengalanto/runic/services/account/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewUser(t *testing.T) {
	testCases := []struct {
		name     string
		email    string
		password string
		assert   func(t *testing.T, err error)
	}{
		{
			name:     "new instance creation success",
			email:    gofakeit.Email(),
			password: mock.ValidPassword(),
			assert: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		{
			name:     "new instance creation failed due to invalid email",
			email:    gofakeit.Word(),
			password: mock.ValidPassword(),
			assert: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
		{
			name:     "new instance creation failed due to invalid password",
			email:    gofakeit.Email(),
			password: mock.InvalidPassword(),
			assert: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := domain.NewUser(tc.email, tc.password)
			tc.assert(t, err)
		})
	}
}

func TestUser_IsActive(t *testing.T) {
	t.Run("returns the correct user account's active state", func(t *testing.T) {
		user := mock.User()

		user.Deactivate()
		assert.False(t, user.IsActive())

		user.Activate()
		assert.True(t, user.IsActive())
	})
}

func TestUser_Activate(t *testing.T) {
	t.Run("activates the user account", func(t *testing.T) {
		user := mock.User()
		user.Activate()
		assert.True(t, user.IsActive())
	})
}

func TestUser_Deactivate(t *testing.T) {
	t.Run("deactivates the user account", func(t *testing.T) {
		user := mock.User()
		user.Deactivate()
		assert.False(t, user.IsActive())
	})
}

func TestUser_RecordLastLogin(t *testing.T) {
	t.Run("sets the last login at with the current timestamp", func(t *testing.T) {
		user := mock.User()
		user.RecordLastLogin()
		assert.False(t, user.LastLoginAt.IsZero())
	})
}

func TestUser_UpdateEmail(t *testing.T) {
	testCases := []struct {
		name    string
		payload string
		assert  func(t *testing.T, expected domain.Email, actual domain.Email, err error)
	}{
		{
			name:    "updates the email successfully",
			payload: gofakeit.Email(),
			assert: func(t *testing.T, expected domain.Email, actual domain.Email, err error) {
				assert.Equal(t, expected, actual)
				require.NoError(t, err)
			},
		},
		{
			name:    "fails to update the email",
			payload: "invalid-email",
			assert: func(t *testing.T, expected domain.Email, actual domain.Email, err error) {
				assert.NotEqual(t, expected, actual)
				require.Error(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			user := mock.User()
			err := user.UpdateEmail(tc.payload)
			tc.assert(t, user.Email, domain.Email(tc.payload), err)
		})
	}
}

func TestUser_UpdatePassword(t *testing.T) {
	testCases := []struct {
		name    string
		payload string
		assert  func(t *testing.T, expected domain.Password, actual domain.Password, err error)
	}{
		{
			name:    "updates the password successfully",
			payload: mock.ValidPassword(),
			assert: func(t *testing.T, expected domain.Password, actual domain.Password, err error) {
				assert.Equal(t, expected, actual)
				require.NoError(t, err)
			},
		},
		{
			name:    "fails to update the password",
			payload: mock.InvalidPassword(),
			assert: func(t *testing.T, expected domain.Password, actual domain.Password, err error) {
				assert.NotEqual(t, expected, actual)
				require.Error(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			user := mock.User()
			err := user.UpdatePassword(tc.payload)
			tc.assert(t, user.Password, domain.Password(tc.payload), err)
		})
	}
}

func TestUser_HashPassword(t *testing.T) {
	t.Run("hash password success", func(t *testing.T) {
		user := mock.User()
		err := user.HashPassword()
		require.NoError(t, err)
	})
}

func TestUser_CheckPassword(t *testing.T) {
	t.Run("hashed and provided password matched", func(t *testing.T) {
		user := mock.User()
		pw := user.Password.String()

		err := user.HashPassword()
		require.NoError(t, err)

		match := user.CheckPassword(pw)
		assert.True(t, match)
	})
}

func TestAggregateAccount(t *testing.T) {
	t.Run("aggregates the root successfully", func(t *testing.T) {
		user := mock.User()
		profile := mock.Profile()

		expected := user
		expected.Profile = &profile

		entity := domain.AggregateRoot(user, profile)
		require.Equal(t, expected, entity)
	})
}

func TestUser_ProfileUpdate(t *testing.T) {
	update := domain.UserProfile{
		Avatar: gofakeit.URL(),
	}

	emptyStr := ""

	testCases := []struct {
		name   string
		user   domain.User
		input  domain.UpdateProfileInput
		assert func(t *testing.T, expected string, actual *string, err error)
	}{
		{
			name: "update user profile success",
			user: mock.Root(),
			input: domain.UpdateProfileInput{
				Avatar: &update.Avatar,
			},
			assert: func(t *testing.T, expected string, actual *string, err error) {
				assert.Equal(t, expected, fn.Deref(actual))
				require.NoError(t, err)
			},
		},
		{
			name: "update user profile failed",
			user: mock.Root(),
			input: domain.UpdateProfileInput{
				Avatar: &emptyStr,
			},
			assert: func(t *testing.T, expected string, actual *string, err error) {
				assert.Equal(t, expected, fn.Deref(actual))
				require.Error(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			user := tc.user
			err := user.UpdateProfile(tc.input)
			tc.assert(t, tc.user.Profile.Avatar, tc.input.Avatar, err)
		})
	}
}
