package dao

import (
	"context"
	"testing"

	e "github.com/n3xtchen/gin-3at/internal/domain/entity"
	m "github.com/n3xtchen/gin-3at/internal/model"
	"gorm.io/gorm"
)

func TestOrderDaoCreateOrder(t *testing.T) {

	var myOrder e.Order

	err := db.Transaction(func(tx *gorm.DB) error {
		// 初始化 OrderDao
		orderDao := NewOrderDao(tx)

		// 创建地址
		address := m.Address{
			UserID:  1,
			Name:    "n3xtchen",
			Phone:   "1231231234",
			Address: "123 Test St",
		}

		// 创建订单项
		items := []m.OrderItem{
			{ProductID: 1, Quantity: 2, Price: 100.0},
			{ProductID: 2, Quantity: 1, Price: 50.0},
		}

		// 创建订单，地址
		order := m.Order{
			BuyerID: 1,
			Status:  2,
			Address: address,
			Items:   items,
		}

		orderEntity := order.ToEntity()
		ctx := context.Background()
		err := orderDao.Save(ctx, &orderEntity)
		myOrder = orderEntity
		return err
	})

	if err != nil {
		t.Errorf("Failed to create order: %v", err)
		return
	}

	t.Logf("Order Wiht ID %d retrieved successfully", myOrder.ID)
	t.Log("Order created successfully")
}

func TestOrderDaoGetOrderByID(t *testing.T) {
	orderDao := NewOrderDao(db)

	// 假设订单 ID 为 1
	orderID := int(1)

	order, err := orderDao.GetDetailByID(orderID)
	if err != nil {
		t.Errorf("Failed to get order by ID %d: %v", orderID, err)
		return
	}

	if order == nil {
		t.Errorf("Order with ID %d not found", orderID)
		return
	}

	t.Logf("Order Wiht ID %d retrieved successfully", order.ID)
	t.Logf("Retrieved %d item", len(order.Items))
}
