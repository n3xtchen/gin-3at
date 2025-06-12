package seed

import (
	"gorm.io/gorm"

	m "github.com/n3xtchen/gin-3at/model"
)

// ProductSeed is a slice of Product objects used for seeding the database.
var ProductSeed = []m.Product{
	{
		Model:       gorm.Model{ID: 1},
		Name:        "Smartphone X",
		Description: "Latest model with advanced features",
		Price:       699.99,
		CategoryID:  3, // Sub-category of Electronics
		ImageURL:    "https://example.com/images/smartphone-x.jpg",
	},
	{
		Model:       gorm.Model{ID: 2},
		Name:        "Laptop Pro",
		Description: "High-performance laptop for professionals",
		Price:       1299.99,
		CategoryID:  1, // Top-level category Electronics
		ImageURL:    "https://example.com/images/laptop-pro.jpg",
	},
	{
		Model:       gorm.Model{ID: 3},
		Name:        "Book of Knowledge",
		Description: "A comprehensive guide to modern technology",
		Price:       29.99,
		CategoryID:  2, // Top-level category Books
		ImageURL:    "https://example.com/images/book-of-knowledge.jpg",
	},
}
