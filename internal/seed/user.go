package seed

import (
	"gorm.io/gorm"

	m "github.com/n3xtchen/gin-3at/internal/model"
)

// UserSeed is a slice of User objects used for seeding the database.
var UserSeed = []m.User{
	{
		Model:    gorm.Model{ID: 1},
		Username: "john_doe",
		Password: "hashed_password_1",
		Email:    "john_doe@mail.com",
		Phone:    "1234567890",
	},
}
