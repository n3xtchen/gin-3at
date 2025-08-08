package dao

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/go-redis/redismock/v9"
)

// test GetProducts
func TestProductDaoGetProducts(t *testing.T) {
	cache, _ := redismock.NewClientMock()
	dao := NewProductDao(db, cache)

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

// test GetProductByID
func TestProductDaoGetProductByID(t *testing.T) {
	cache, mock := redismock.NewClientMock()
	dao := NewProductDao(db, cache)

	// Get product by ID
	productID := 1
	mock.ExpectGet(strconv.Itoa(productID)).RedisNil()
	product, err := dao.FindByID(productID)
	if err != nil {
		t.Fatalf("Failed to get product by ID %d: %v", productID, err)
	}

	expectedCache, _ := json.Marshal(product)
	mock.ExpectGet(strconv.Itoa(productID)).SetVal(string(expectedCache))

	// pretty print product
	t.Logf("Retrieved product: %+v", product)

	if product.ID != productID {
		t.Fatalf("Expected product ID %d, got %d", productID, product.ID)
	}
}
