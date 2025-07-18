package dao

import (
	"testing"
)

// test GetProducts
func TestProductDaoGetProducts(t *testing.T) {
	dao := NewProductDao(db)

	// Get all products
	products, err := dao.List()
	if err != nil {
		t.Fatalf("Failed to get products: %v", err)
	}

	// pretty print products
	t.Logf("Retrieved %d products", len(products))

	if len(products) == 0 {
		t.Fatal("Expected to get some products, but got none")
	}
}
