package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// func TestCategoryIntegration(t *testing.T) {
// 	t.Logf("Test Category Integration")
//
// 	// Test Create Category
// 	createData := map[string]string{
// 		"name":        "Test Category",
// 		"description": "This is a test category",
// 	}
//
// 	jsonCreateBody, _ := json.Marshal(createData)
// 	reqCreate, _ := http.NewRequest("POST", "/api/v1/categories", bytes.NewBuffer(jsonCreateBody))
// 	reqCreate.Header.Set("Content-Type", "application/json")
// 	respCreate := httptest.NewRecorder()
// 	r.ServeHTTP(respCreate, reqCreate)
// 	if respCreate.Code != http.StatusCreated {
// 		t.Fatalf("expected status 201, got %d", respCreate.Code)
// 	}
// 	t.Logf("Create Category Response: %s", respCreate.Body.String())
// }

func TestCategoryHandler_GetCategories(t *testing.T) {
	// Test Get Categories
	reqGet, _ := http.NewRequest("GET", "/api/v1/categories", nil)
	respGet := httptest.NewRecorder()
	r.ServeHTTP(respGet, reqGet)
	if respGet.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", respGet.Code)
	}
	t.Logf("Get Categories Response: %s", respGet.Body.String())
}

func TestCategoryHandler_GetCategory(t *testing.T) {
	// Test Get Category by ID
	reqGetByID, _ := http.NewRequest("GET", "/api/v1/categories/1", nil)
	respGetByID := httptest.NewRecorder()
	r.ServeHTTP(respGetByID, reqGetByID)
	if respGetByID.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", respGetByID.Code)
	}
	t.Logf("Get Category by ID Response: %s", respGetByID.Body.String())
}
