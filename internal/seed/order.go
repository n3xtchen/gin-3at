package seed

import (
	"gorm.io/gorm"

	m "github.com/n3xtchen/gin-3at/internal/model"
)

// OrderSeed is a slice of Order objects used for seeding the database.
// It contains a list of orders with associated user IDs, address IDs, Amount, and Status.
var OrderSeed = []m.Order{
	{
		Model:     gorm.Model{ID: 1},
		OrderNum:  "ORD123456789",
		BuyerID:   1,
		AddressID: 1,
		Amount:    199.99,
		Status:    2,
	},
	{
		Model:     gorm.Model{ID: 2},
		OrderNum:  "ORD987654321",
		BuyerID:   1,
		AddressID: 2,
		Amount:    299.99,
		Status:    1,
	},
}
