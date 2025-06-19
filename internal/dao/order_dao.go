package dao

import (
	"gorm.io/gorm"

	e "github.com/n3xtchen/gin-3at/internal/domain/entity"
	repo "github.com/n3xtchen/gin-3at/internal/domain/repository"
	m "github.com/n3xtchen/gin-3at/internal/model"
)

type OrderDao struct {
	db *gorm.DB
}

// NewOrderDao creates a new instance of OrderDao.
func NewOrderDao(db *gorm.DB) repo.OrderRepository {
	return &OrderDao{
		db,
	}
}

// CreateOrder creates a new order with the given details.
func (dao *OrderDao) Save(order *e.Order) error {

	// todo: Validate the order, address, and items.
	orderModel := m.FromEntityOrder(order)

	if err := dao.db.Create(&orderModel).Error; err != nil {
		return err
	}

	return nil
}

// GetOrderByID retrieves an order by its ID.
func (dao *OrderDao) GetDetailByID(orderID int) (*e.Order, error) {
	var order m.Order
	if err := dao.db.Preload("Items").Preload("Address").Preload("Items.Product").First(&order, orderID).Error; err != nil {
		// return order.ToEntity(), err
		return nil, err
	}
	return order.ToEntity(), nil
}
