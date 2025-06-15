package dto

// OrderItem represents an item in an order.
type CreateOrderItemReq struct {
	ProductID uint    `json:"product_id" binding:"required"`     // ID of the Product
	Quantity  int     `json:"quantity" binding:"required,min=1"` // Quantity of the product in the Order
	Price     float32 `json:"price" binding:"required"`          // Price of the product at the time of order
}

// Order Create Request
// includes details about the order, Address, and OrderItems.
type CreateOrderReq struct {
	Amount     float32              `json:"amount" binding:"required"` // Total amount of the order
	Address    CreateAddressReq     `json:"address" binding:"required"`
	OrderItems []CreateOrderItemReq `json:"order_items" binding:"required"`
}

// Order Update Request
type UpdateOrderReq struct {
	OrderID uint             `json:"order_id" binding:"required"` // ID of the Order to Update
	Status  uint             `json:"status" binding:"required"`   // New status of the OrderID
	Address CreateAddressReq `json:"address" binding:"required"`  // Updated Address Details
	// Note: OrderItems are not included in the update request as they are typically not modified after creation.
}

// Order Details Request
type GetOrderByIDReq struct {
	OrderID uint `json:"order_id" binding:"required"` // ID of the Order to retrieves
}
