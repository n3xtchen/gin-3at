package dao

import (
	"testing"
)

// test GetProducts
func TestProductDaoGetProducts(t *testing.T) {
	dao := NewProductDao()

	// Get all products
	products, err := dao.GetProducts(10, 0)
	if err != nil {
		t.Fatalf("Failed to get products: %v", err)
	}

	// pretty print products
	t.Logf("Retrieved %d products", len(products))
	for _, product := range products {
		t.Logf("Product: %+v", product.Name)
		t.Logf("Product Category: %+v", product.Category.Name)
	}

	if len(products) == 0 {
		t.Fatal("Expected to get some products, but got none")
	}

	for _, product := range products {
		t.Logf("Product: %+v", product)
	}
}
