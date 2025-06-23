package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gin-gonic/gin"

	e "github.com/n3xtchen/gin-3at/internal/domain/entity"
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
				"name":     "Item 1",
				"price":    10.0,
				"quantity": 2,
			}, {
				"name":     "Item 2",
				"price":    20.0,
				"quantity": 1,
			},
		},
	}
)

func GetTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}
	return ctx
}

func MockJsonPost(c *gin.Context, content interface{}) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	// the request body must be an io.ReadCloser
	// the bytes buffer though doesn't implement io.Closer,
	// so you wrap it in a no-op closer
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}

type OrderServiceMock struct {
}

func (m *OrderServiceMock) CreateOrder(order *e.Order) error {
	return nil // Default behavior if no mock function is set
}

func (m *OrderServiceMock) GetOrderDetail(orderID int) (*e.Order, error) {
	order := &e.Order{
		ID:     orderID,
		Amount: 100.0,
		Address: e.Address{
			Name:    "John Doe",
			Phone:   "12345678901",
			Address: "123 Main St",
		},
		Items: []e.OrderItem{
			{
				ProductID: 1,
				Quantity:  2,
				Price:     10.0,
			},
		},
	}

	return order, nil
}
