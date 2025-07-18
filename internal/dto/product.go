package dto

import (
	e "github.com/n3xtchen/gin-3at/internal/domain/entity"
)

type ProductReq struct {
	ProductID int `uri:"id" binding:"required"` // ID of the Product
}

type ProductRes struct {
	ProductID           int     `json:"product_id"`            // ID of the ProductID
	ProductName         string  `json:"product_name"`          // Name of the ProductID
	ProductDesc         string  `json:"product_desc"`          // Description of the ProductID
	ProductPrice        float64 `json:"product_price"`         // Price of the ProductID
	ProductImage        string  `json:"product_image"`         // Image URL of the ProductID
	ProductCategoryID   int     `json:"product_category_id"`   // Category ID of the ProductID
	ProductCategoryName string  `json:"product_category_name"` // Category Name of the ProductID
}

func ToProductRes(product *e.Product) ProductRes {
	return ProductRes{
		ProductID:           product.ID,
		ProductName:         product.Name,
		ProductDesc:         product.Description,
		ProductPrice:        product.Price,
		ProductImage:        product.ImageURL,
		ProductCategoryID:   product.CategoryID,
		ProductCategoryName: product.Category.Name,
	}
}
