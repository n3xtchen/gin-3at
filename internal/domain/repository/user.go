package repository

import (
	e "github.com/n3xtchen/gin-3at/internal/domain/entity"
)

type UserRepository interface {

	// CreateUser creates a new user with the given details.
	Create(user *e.User) error

	// UpdateUser updates an existing user with the given details.
	Update(user *e.User, newInfo map[string]interface{}) error

	// ResetPassword resets the user's password.
	ResetPassword(user *e.User, newPassword string) error

	// CheckUserExists checks if a user exists by username or email.
	FindBy(usernameOrEmail string) (*e.User, error)

	// CheckUserPassword checks if the provided password matches the user's password.
	CheckPassword(user *e.User, password string) (bool, error)
}
