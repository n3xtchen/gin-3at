package dao

import (
	m "github.com/n3xtchen/gin-3at/model"
)

type OrderDao struct {
}

// CreateOrder creates a new order with the given details.
func (dao *OrderDao) CreateOrder(order m.Order, address m.Address, items []m.OrderItem) error {

	// todo: Validate the order, address, and items.

	// start session for transaction
	tx := DB.Begin()

	// if address exists, get address id, else create a new address.
	addressID, err := NewAddressDao().GetOrCreateAddressID(address)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Insert the order into the database.
	order.AddressID = addressID
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Insert each order item into the database.
	for _, item := range items {
		item.OrderID = order.ID
		if err := tx.Create(&item).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit the transaction.
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
