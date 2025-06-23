package dto

import (
	e "github.com/n3xtchen/gin-3at/internal/domain/entity"
)

// OrderItem represents an item in an order.
type CreateOrderItemReq struct {
	ProductID int     `json:"product_id" binding:"required"`     // ID of the Product
	Quantity  int     `json:"quantity" binding:"required,min=1"` // Quantity of the product in the Order
	Price     float64 `json:"price" binding:"required"`          // Price of the product at the time of order
}

func (c *CreateOrderItemReq) ToEntity() e.OrderItem {
	return e.OrderItem{
		ProductID: c.ProductID,
		Quantity:  c.Quantity,
		Price:     c.Price,
	}
}

// Order Create Request
// includes details about the order, Address, and OrderItems.
type CreateOrderReq struct {
	Amount     float64              `json:"amount" binding:"required"` // Total amount of the order
	Address    CreateAddressReq     `json:"address" binding:"required"`
	OrderItems []CreateOrderItemReq `json:"order_items" binding:"required"`
}

func (c CreateOrderReq) ToEntity() *e.Order {
	items := make([]e.OrderItem, len(c.OrderItems))
	for i, item := range c.OrderItems {
		items[i] = item.ToEntity()
	}

	order := &e.Order{
		Amount:  c.Amount,
		Address: c.Address.ToEntity(),
		Items:   items,
	}

	return order
}

// Order Update Request
type UpdateOrderReq struct {
	OrderID int              `json:"order_id" binding:"required"` // ID of the Order to Update
	Status  int              `json:"status" binding:"required"`   // New status of the OrderID
	Address CreateAddressReq `json:"address" binding:"required"`  // Updated Address Details
	// Note: OrderItems are not included in the update request as they are typically not modified after creation.
}

// Order Details Request
type GetOrderByIDReq struct {
	OrderID uint `json:"order_id" binding:"required"` // ID of the Order to retrieves
}
