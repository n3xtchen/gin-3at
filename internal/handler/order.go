package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/n3xtchen/gin-3at/internal/dto"
	"github.com/n3xtchen/gin-3at/internal/service"
)

type OrderHandler struct {
	// You can add fields here if needed, such as a service instance
	OrderService service.OrderService
}

// NewOrderHandler creates a new OrderHandler instance.
func NewOrderHandler(orderService service.OrderService) *OrderHandler {
	return &OrderHandler{
		OrderService: orderService,
	}
}

// CreateOrder handles the creation of a new order.
func (order *OrderHandler) Save(c *gin.Context) {
	// validate it, and then call the OrderDao to create the order.
	var orderReq dto.CreateOrderReq
	if err := c.ShouldBindJSON(&orderReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order.OrderService.CreateOrder(orderReq.ToEntity())

	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully"})
}
