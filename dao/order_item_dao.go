package dao

import (
	"gorm.io/gorm"

	m "github.com/n3xtchen/gin-3at/model"
)

// OrderItemDao is a data access object for managing OrderItem entities.
type OrderItemDao struct {
	*gorm.DB
}

// NewOrderItemDao creates a new instance of OrderItemDao.
func NewOrderItemDao() *OrderItemDao {
	return &OrderItemDao{
		DB: DB,
	}
}

// CreateOrderItem creates a new order item in the database.
func (dao *OrderItemDao) CreateOrderItem(item m.OrderItem) error {
	if err := dao.DB.Create(&item).Error; err != nil {
		return err
	}
	return nil
}
