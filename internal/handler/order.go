package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/n3xtchen/gin-3at/internal/dto"
)

// CreateOrder handles the creation of a new order.
func CreateOrder(c *gin.Context) {
	// validate it, and then call the OrderDao to create the order.
	var order dto.CreateOrderReq
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if err := dao.NewOrderDao().CreateOrder(order); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully"})
}
