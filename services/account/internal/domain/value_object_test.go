package domain_test

import (
	"reflect"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ssengalanto/runic/services/account/internal/domain"
	"github.com/ssengalanto/runic/services/account/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEmail_IsValid(t *testing.T) {
	testCases := []struct {
		name    string
		payload string
		assert  func(t *testing.T, result bool, err error)
	}{
		{
			name:    "valid email",
			payload: gofakeit.Email(),
			assert: func(t *testing.T, result bool, err error) {
				assert.True(t, result)
				require.NoError(t, err)
			},
		},
		{
			name:    "invalid email",
			payload: "invalid-email",
			assert: func(t *testing.T, result bool, err error) {
				assert.False(t, result)
				require.Error(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			email := domain.Email(tc.payload)
			ok, err := email.IsValid()
			tc.assert(t, ok, err)
		})
	}
}

func TestEmail_Update(t *testing.T) {
	testCases := []struct {
		name    string
		current string
		update  string
		assert  func(t *testing.T, expected domain.Email, actual domain.Email, err error)
	}{
		{
			name:    "change email successful",
			current: gofakeit.Email(),
			update:  gofakeit.Email(),
			assert: func(t *testing.T, expected domain.Email, actual domain.Email, err error) {
				assert.Equal(t, expected, actual)
				require.NoError(t, err)
			},
		},
		{
			name:    "change email failed",
			current: gofakeit.Email(),
			update:  "invalid-email",
			assert: func(t *testing.T, expected domain.Email, actual domain.Email, err error) {
				assert.NotEqual(t, expected, actual)
				require.Error(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := domain.Email(tc.current)
			ne, err := e.Update(tc.update)
			tc.assert(t, domain.Email(tc.update), ne, err)
		})
	}
}

func TestEmail_String(t *testing.T) {
	t.Run("convert email to type string", func(t *testing.T) {
		email := domain.Email(gofakeit.Email()).String()
		kind := reflect.TypeOf(email).String()
		require.Equal(t, "string", kind)
	})
}

func TestPassword_IsValid(t *testing.T) {
	testCases := []struct {
		name    string
		payload string
		assert  func(t *testing.T, result bool, err error)
	}{
		{
			name:    "valid password",
			payload: mock.ValidPassword(),
			assert: func(t *testing.T, result bool, err error) {
				assert.True(t, result)
				require.Nil(t, err)
			},
		},
		{
			name:    "invalid password",
			payload: mock.InvalidPassword(),
			assert: func(t *testing.T, result bool, err error) {
				assert.False(t, result)
				require.NotNil(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			password := domain.Password(tc.payload)
			ok, err := password.IsValid()
			tc.assert(t, ok, err)
		})
	}
}

func TestPassword_Update(t *testing.T) {
	testCases := []struct {
		name    string
		current string
		update  string
		assert  func(t *testing.T, expected domain.Password, actual domain.Password, err error)
	}{
		{
			name:    "change password successful",
			current: mock.ValidPassword(),
			update:  mock.ValidPassword(),
			assert: func(t *testing.T, expected domain.Password, actual domain.Password, err error) {
				assert.Equal(t, expected, actual)
				require.NoError(t, err)
			},
		},
		{
			name:    "change password failed",
			current: mock.ValidPassword(),
			update:  mock.InvalidPassword(),
			assert: func(t *testing.T, expected domain.Password, actual domain.Password, err error) {
				assert.NotEqual(t, expected, actual)
				require.Error(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pw := domain.Password(tc.current)
			npw, err := pw.Update(tc.update)
			tc.assert(t, domain.Password(tc.update), npw, err)
		})
	}
}

func TestPassword_Hash(t *testing.T) {
	t.Run("password hashed", func(t *testing.T) {
		pw := domain.Password(mock.ValidPassword())
		hpw, err := pw.Hash()
		assert.NotEqual(t, pw, hpw)
		require.NoError(t, err)
	})
}

func TestPassword_Check(t *testing.T) {
	t.Run("password matched", func(t *testing.T) {
		pw := domain.Password(mock.ValidPassword())
		hpw, err := pw.Hash()
		require.NoError(t, err)

		err = hpw.Check(pw.String())
		require.NoError(t, err)
	})
	t.Run("password mismatched", func(t *testing.T) {
		pw := domain.Password(mock.ValidPassword())
		hpw, err := pw.Hash()
		require.NoError(t, err)

		err = hpw.Check(pw.String())
		require.NoError(t, err)
	})
}

func TestPassword_String(t *testing.T) {
	t.Run("convert password to type string", func(t *testing.T) {
		password := domain.Password(mock.ValidPassword()).String()
		kind := reflect.TypeOf(password).String()
		require.Equal(t, "string", kind)
	})
}
