package dao

import (
	"context"
	"log"

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
func (dao *OrderDao) Save(ctx context.Context, order *e.Order) error {

	db := GetDBFromContext(ctx, dao.db)

	// todo: Validate the order, address, and items.

	// Set the address and items to the order Model
	addresModel := m.FromEntityAddress(order.Address)
	if err := db.Create(&addresModel).Error; err != nil {
		return err
	}

	log.Printf("Address created successfully with ID: %d", addresModel.ID)

	orderModel := m.FromEntityOrder(order)
	orderModel.AddressID = addresModel.ID

	if err := db.Create(&orderModel).Error; err != nil {
		return err
	}

	// Create order Items
	for _, item := range order.Items {
		orderItemModel := m.FromEntityOrderItem(item)
		orderItemModel.OrderID = orderModel.ID
		if err := db.Create(&orderItemModel).Error; err != nil {
			return err
		}
		log.Printf("Order item created successfully with ID: %d", orderItemModel.ID)
	}

	// 新的值更新 order 实体的值
	*order = orderModel.ToEntity()

	log.Printf("Order created successfully with ID: %d", order.ID)

	return nil
}

// GetOrderByID retrieves an order by its ID.
func (dao *OrderDao) GetDetailByID(orderID int) (*e.Order, error) {
	var order m.Order
	if err := dao.db.Preload("Items").Preload("Address").Preload("Items.Product").First(&order, orderID).Error; err != nil {
		// return order.ToEntity(), err
		return nil, err
	}

	orderEntity := order.ToEntity()
	return &orderEntity, nil
}
