package repository

import (
	e "github.com/n3xtchen/gin-3at/internal/domain/entity"
)

// ProductRepository defines the methods for interacting with product data.
type ProductRepository interface {
	// GetProductByID retrieves a product by its ID.
	FindByID(productID int) (*e.Product, error)
	// GetAllProducts retrieves all products.
	List() ([]*e.Product, error)
	// CreateProduct creates a new product.
	Save(product *e.Product) error
	// UpdateProduct updates an existing product.
	Update(product *e.Product) error
	// DeleteProduct deletes a product by its ID.
	Delete(productID int) error
}
