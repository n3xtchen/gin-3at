package seed

import (
	"gorm.io/gorm"

	m "github.com/n3xtchen/gin-3at/model"
)

// AddressSeed is a slice of Address objects used for seeding the database.
var AddressSeed = []m.Address{
	{
		Model:     gorm.Model{ID: 1},
		UserID:    1,
		Name:      "John Doe",
		Phone:     "1234567890",
		Address:   "123 Main St, Springfield, USA",
		IsDefault: true,
	},
	{
		Model:     gorm.Model{ID: 2},
		UserID:    1,
		Name:      "Jane Smith",
		Phone:     "0987654321",
		Address:   "456 Elm St, Springfield, USA",
		IsDefault: false,
	},
}
