package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-contrib/sessions"

	e "github.com/n3xtchen/gin-3at/internal/domain/entity"
	"github.com/n3xtchen/gin-3at/internal/dto"
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

func TestOrderHandler_Save(t *testing.T) {

	orderHandler := NewOrderHandler(&OrderServiceMock{}) // Pass a mock or real service as needed

	w := httptest.NewRecorder()
	ctx := GetTestGinContextWithSession(w)

	// Simulate setting userID in context
	session := sessions.Default(ctx)
	session.Set("userID", 1)
	session.Save()

	t.Log("Session before logout:", session.Get("userID"))

	data := orderMock

	MockJsonPost(ctx, data)

	orderHandler.Save(ctx)

	// Check the response status code
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Additional checks can be added here for response body, etc.
	t.Log("Response Body:", w.Body.String())
}

// dto 层一般不做单元测试，因为它们只是数据传输对象，没有业务逻辑。
// 都放在 handler_test.go 中进行集成测试。
func TestOrderDTO(t *testing.T) {

	var orderReq dto.CreateOrderReq

	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)

	data := orderMock

	MockJsonPost(ctx, data)

	if err := ctx.ShouldBindJSON(&orderReq); err != nil {
		t.Fatalf("Failed to bind JSON: %v", err)
	}
}
