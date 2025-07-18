package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProductIntegration(t *testing.T) {
	t.Logf("Test Product Integration")

	// Test Get Products
	reqGet, _ := http.NewRequest("GET", "/api/v1/products", nil)
	respGet := httptest.NewRecorder()
	r.ServeHTTP(respGet, reqGet)
	if respGet.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", respGet.Code)
	}
	t.Logf("Get Products Response: %s", respGet.Body.String())

	// Test Get Product by ID
	reqGetByID, _ := http.NewRequest("GET", "/api/v1/products/1", nil)
	respGetByID := httptest.NewRecorder()
	r.ServeHTTP(respGetByID, reqGetByID)
	if respGetByID.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", respGetByID.Code)
	}
	t.Logf("Get Product by ID Response: %s", respGetByID.Body.String())
}
