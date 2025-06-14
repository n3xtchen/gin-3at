package dao

import (
	"gorm.io/gorm"

	m "github.com/n3xtchen/gin-3at/model"
)

type OrderDao struct {
	db *gorm.DB
}

// NewOrderDao creates a new instance of OrderDao.
func NewOrderDao(db *gorm.DB) *OrderDao {
	return &OrderDao{
		db,
	}
}

// CreateOrder creates a new order with the given details.
func (dao *OrderDao) CreateOrder(order m.Order) error {

	// todo: Validate the order, address, and items.

	if err := dao.db.Create(&order).Error; err != nil {
		return err
	}

	return nil
}

// GetOrderByID retrieves an order by its ID.
func (dao *OrderDao) GetOrderByID(orderID uint) (*m.Order, error) {
	var order m.Order
	if err := dao.db.Preload("OrderItems").Preload("Address").Preload("OrderItems.Product").First(&order, orderID).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
