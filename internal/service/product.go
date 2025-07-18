package service

import (
	e "github.com/n3xtchen/gin-3at/internal/domain/entity"
	repo "github.com/n3xtchen/gin-3at/internal/domain/repository"
)

// ProductService defines the methods for product-related operations.
type ProductService interface {
	// GetProductByID retrieves a product by its ID.
	GetProductByID(productID int) (*e.Product, error)
	// GetAllProducts retrieves all products.
	GetAllProducts() ([]*e.Product, error)
	// CreateProduct creates a new product.
	CreateProduct(product *e.Product) error
	// UpdateProduct updates an existing product.
	UpdateProduct(product *e.Product) error
	// DeleteProduct deletes a product by its ID.
	DeleteProduct(productID int) error
}

// ProductServiceImp implements the ProductService interface.
type ProductServiceImp struct {
	ProductRepo repo.ProductRepository
}

// NewProductService creates a new instance of ProductService.
func NewProductService(productRepo repo.ProductRepository) ProductService {
	return &ProductServiceImp{
		ProductRepo: productRepo,
	}
}

// GetProductByID retrieves a product by its ID.
func (s *ProductServiceImp) GetProductByID(productID int) (*e.Product, error) {
	return s.ProductRepo.FindByID(productID)
}

// GetAllProducts retrieves all products.
func (s *ProductServiceImp) GetAllProducts() ([]*e.Product, error) {
	return s.ProductRepo.List()
}

// CreateProduct creates a new product.
func (s *ProductServiceImp) CreateProduct(product *e.Product) error {
	return s.ProductRepo.Save(product)
}

// UpdateProduct updates an existing product.
func (s *ProductServiceImp) UpdateProduct(product *e.Product) error {
	return s.ProductRepo.Update(product)
}

// DeleteProduct deletes a product by its ID.
func (s *ProductServiceImp) DeleteProduct(productID int) error {
	return s.ProductRepo.Delete(productID)
}
