package dao

import (
	"gorm.io/gorm"

	m "github.com/n3xtchen/gin-3at/model"
)

// OrderItemDao is a data access object for managing OrderItem entities.
type OrderItemDao struct {
	db *gorm.DB
}

// NewOrderItemDao creates a new instance of OrderItemDao.
func NewOrderItemDao(db *gorm.DB) *OrderItemDao {
	return &OrderItemDao{
		db,
	}
}

// CreateOrderItem creates a new order item in the database.
func (dao *OrderItemDao) Create(item m.OrderItem) error {
	if err := dao.db.Create(&item).Error; err != nil {
		return err
	}
	return nil
}
