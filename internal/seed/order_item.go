package seed

import (
	"gorm.io/gorm"

	m "github.com/n3xtchen/gin-3at/internal/model"
)

// OrderItemSeed is a slice of OrderItem objects used for seeding the database.
var OrderItemSeed = []m.OrderItem{
	{
		Model:     gorm.Model{ID: 1},
		OrderID:   1,
		ProductID: 1,
		Quantity:  2,
		Price:     99.99,
	},
	{
		Model:     gorm.Model{ID: 2},
		OrderID:   1,
		ProductID: 2,
		Quantity:  1,
		Price:     199.99,
	},
	{
		Model:     gorm.Model{ID: 3},
		OrderID:   2,
		ProductID: 3,
		Quantity:  3,
		Price:     49.99,
	},
}
