package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/n3xtchen/gin-3at/internal/dto"
	"github.com/n3xtchen/gin-3at/internal/service"
)

// ProductHandler handles HTTP requests related to products.
type ProductHandler struct {
	ProductService service.ProductService
}

// NewProductHandler creates a new instance of ProductHandler.
func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{
		ProductService: productService,
	}
}

// GetProductByID retrieves a product by its ID.
// @Summary Get Product by Is
// @Description Get a product by its IsDefault
// @Tags Product
// @Accept json
// @Produce json
// @Param product_id path int true "Product ID"
// @Success 200 {object} dto.ProductRes "Product retrieved successfully"
// @Failure 400 {object} dto.APIResponse "Invalid request parameters"
// @Failure 404 {object} dto.APIResponse "Product not found"
// @Failure 500 {object} dto.APIResponse "Internal server error"
// @Router /products/{product_id} [get]
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	productReq := &dto.ProductReq{}

	if err := c.ShouldBindUri(&productReq); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request parameters"})
		return
	}

	product, err := h.ProductService.GetProductByID(productReq.ProductID)
	if err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(200, dto.ToProductRes(product))
}

// GetAllProducts retrieves all products.
// @Summary Get All Products
// @Description Get all Products
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {array} dto.ProductRes "List of products"
// @Failure 500 {object} dto.APIResponse "Failed to retrieve products"
// @Router /products [get]
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.ProductService.GetAllProducts()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve products"})
		return
	}

	var productResList []dto.ProductRes
	for _, product := range products {
		productResList = append(productResList, dto.ToProductRes(product))
	}

	c.JSON(200, productResList)
}
