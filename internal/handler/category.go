package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/n3xtchen/gin-3at/internal/dto"
	"github.com/n3xtchen/gin-3at/internal/service"
)

// CategoryHandler defines the methods for handling category-related HTTP requests.
type CategoryHandler struct {
	CategoryService service.CategoryService
}

// NewCategoryHandler creates a new instance of CategoryHandler
func NewCategoryHandler(categoryService service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		CategoryService: categoryService,
	}
}

// GetCategoryByID retrieves a category by its ID.
// @Summary Get Category by Is
// @Description Get a category by its Is
// @Tags Category
// @Accept json
// @Produce json
// @Param category_id path int true "Category ID"
// @Success 200 {object} dto.CategroyRes "Category retrieved successfully"
// @Failure 400 {object} dto.APIResponse "Invalid request parameters"
// @Failure 404 {object} dto.APIResponse "Category not found"
// @Failure 500 {object} dto.APIResponse "Internal server error"
// @Router /categories/{category_id} [get]
func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	cateReq := &dto.CategoryReq{}

	if err := c.ShouldBindUri(&cateReq); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request parameters"})
		return
	}

	category, err := h.CategoryService.GetCategoryByID(cateReq.CategoryID)
	if err != nil {
		c.JSON(404, gin.H{"error": cateReq.CategoryID})
		// c.JSON(404, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(200, dto.ToCategoryRes(category))
	return
}

// GetAllCategories retrieves all categories.
// @Summary Get All GetAllCategories
// @Description Get all Categories
// @Tags Category
// @Accept json
// @Produce json
// @Success 200 {array} dto.CategroyRes "List of categories"
// @Failure 500 {object} dto.APIResponse "Failed to retrieve categories"
// @Router /categories [get]
func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	categories, err := h.CategoryService.GetAllCategories()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve categories"})
		return
	}

	var categoryResponses []dto.CategroyRes
	for _, category := range categories {
		categoryResponses = append(categoryResponses, dto.ToCategoryRes(category))
	}

	c.JSON(200, categoryResponses)
	return
}
