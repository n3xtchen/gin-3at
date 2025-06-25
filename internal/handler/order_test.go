package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n3xtchen/gin-3at/internal/dto"
)

func TestOrderHandler_Save(t *testing.T) {

	orderHandler := NewOrderHandler(&OrderServiceMock{}) // Pass a mock or real service as needed

	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)

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
