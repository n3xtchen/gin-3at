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

// CreateOrder handles the creation creationof a new order.
// @Summary Create a new order
// @Schemes http https
// @Description
// @Tags orders
// @Accept json
// @Produce json
// @Param order body dto.CreateOrderReq true "Order details"
// @Success 201 {object} dto.APIResponse "Order created successfully"
// @Failure 400 {object} dto.APIResponse "Invalid request parameters"
// @Failure 500 {object} dto.APIResponse "Inyernal server error"
// @Router /orders [post]
func (order *OrderHandler) Save(c *gin.Context) {

	// validate it, and then call the OrderDao to create the order.
	var orderReq dto.CreateOrderReq
	if err := c.ShouldBindJSON(&orderReq); err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Code: dto.CodeParamError, Message: err.Error()})
		return
	}

	err := order.OrderService.CreateOrder(orderReq.ToEntity())

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.APIResponse{Code: dto.CodeUnknownError, Message: "Failed to create order"})
	} else {
		c.JSON(http.StatusCreated, dto.APIResponse{Code: dto.CodeSuccess, Message: "Order created successfully"})
	}
}
