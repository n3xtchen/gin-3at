package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	orderMock = map[string]interface{}{
		"amount": 100.0,
		"address": map[string]interface{}{
			"name":    "John Doe",
			"phone":   "12345678901",
			"address": "123 Main St",
		},
		"order_items": []map[string]interface{}{
			{
				"product_id": 1,
				"price":      10.0,
				"quantity":   2,
			}, {
				"product_id": 2,
				"price":      20.0,
				"quantity":   1,
			},
		},
	}
)

func TestOrderIntegration(t *testing.T) {

	t.Logf("Test Order Integration")

	jsonBody, _ := json.Marshal(orderMock)

	req, _ := http.NewRequest("POST", "/api/v1/orders", bytes.NewBuffer(jsonBody))
	// Add the login cookie to the Request
	for _, cookie := range loginCookie {
		req.AddCookie(cookie)
	}

	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	t.Logf("Response: %s", resp.Body.String())
	if resp.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", resp.Code)
	}
}
