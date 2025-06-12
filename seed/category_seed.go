package seed

import (
	"gorm.io/gorm"

	m "github.com/n3xtchen/gin-3at/model"
)

// CategorySeed is a slice of Category objects used for seeding the database.
var CategorySeed = []m.Category{
	{
		Model:       gorm.Model{ID: 1},
		Name:        "Electronics",
		Description: "Devices and gadgets",
	},
	{
		Model:       gorm.Model{ID: 2},
		Name:        "Books",
		Description: "Printed and digital books",
	},
	{
		Model:       gorm.Model{ID: 3},
		Name:        "Smartphones",
		Description: "Latest smartphones and accessories",
	},
}
