package dao

import (
	"testing"

	m "github.com/n3xtchen/gin-3at/model"
)

func TestOrderDaoCreateOrder(t *testing.T) {
	orderDao := &OrderDao{}

	// 创建订单，地址
	order := m.Order{
		UserID: 1,
		Status: m.OrderStatusUnpaid,
	}

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

	err := orderDao.CreateOrder(order, address, items)
	if err != nil {
		t.Errorf("Failed to create order: %v", err)
		return
	}
	t.Log("Order created successfully")
}
