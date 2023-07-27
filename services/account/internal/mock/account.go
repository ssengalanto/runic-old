package mock

import (
	"log"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/ssengalanto/runic/services/account/internal/domain"
)

func ValidPassword() string {
	//nolint:gomnd //unnecessary
	return gofakeit.Password(true, true, true, true, false, 10)
}

func InvalidPassword() string {
	//nolint:gomnd //unnecessary
	return gofakeit.Password(true, true, true, true, false, 5)
}

// User create a new mock account user entity.
func User() domain.User {
	user, err := domain.NewUser(
		gofakeit.Email(),
		ValidPassword(),
	)

	if err != nil {
		log.Fatal(err)
	}

	return user
}

// Profile create a new mock account user profile entity.
func Profile() domain.UserProfile {
	user, err := domain.NewUserProfile(
		uuid.New(),
		gofakeit.FirstName(),
		gofakeit.LastName(),
		gofakeit.Date(),
	)

	if err != nil {
		log.Fatal(err)
	}

	return user
}

// Root create a new account aggregate root.
func Root() domain.User {
	user := User()
	profile := Profile()
	return domain.AggregateRoot(user, profile)
}
