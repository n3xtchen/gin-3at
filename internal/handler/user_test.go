package handler

import (
	"net/http/httptest"
	"testing"

	e "github.com/n3xtchen/gin-3at/internal/domain/entity"
	"github.com/n3xtchen/gin-3at/internal/service"
)

// MockUserService is a mock implementation of the UserService interface for testing purposes.
type MockUserService struct {
}

func (m *MockUserService) Register(user *e.User) error {
	// Mock implementation for testing
	if user.Username == "testuser" && user.Password == "password123" && user.Email == "testUser@mail.com" {
		return nil // Simulate successful registration
	}
	return &service.BizError{Code: 10003, Message: "User already exists"} // Simulate user already exists Error
}

func (m *MockUserService) ResetPassword(username, newPassword string) error {
	// Mock implementation for testing
	if username == "testuser" && newPassword == "newpassword123" {
		return nil // Simulate successful password reset
	}
	return &service.BizError{Code: 10005, Message: "User reset password failed"} // Simulate reset password Error
}

func (m *MockUserService) Login(username, password string) (*e.User, error) {
	// Mock implementation for testing
	if username == "testuser" && password == "password123" {
		return &e.User{
			Username: username,
			Password: password,
			Email:    "mail@mail.com",
		}, nil // Simulate successful Login
	} else {
		return nil, &service.BizError{Code: 10002, Message: "Invalid password"} // Simulate invalid password Error
	}
}

func (m *MockUserService) Logout(userID int) error {
	// Mock implementation for testing
	if userID > 0 {
		return nil // Simulate successful logout
	}
	return &service.BizError{Code: 10001, Message: "User not found"} // Simulate user not found Error
}

// Test RegisterUser tests the RegisterUser handler.
func TestUserHandler_RegisterUser(t *testing.T) {

	userHandler := NewUserHandler(&MockUserService{}) // Mock the service for testing

	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)

	data := map[string]interface{}{
		"username": "testuser",
		"password": "password123",
		"email":    "testUser@mail.com",
	}

	MockJsonPost(ctx, data)

	userHandler.RegisterUser(ctx)

	if w.Code != 201 {
		t.Errorf("Expected status code 201, got %d", w.Code)
	}
	if w.Body.String() != `{"message":"User registered successfully"}` {
		t.Errorf("Expected response body to be 'User registered successfully', got '%s'", w.Body.String())
	}
	t.Log("Response Body:", w.Body.String())
}

func TestUserHandler_LoginUser(t *testing.T) {
	userHandler := NewUserHandler(&MockUserService{}) // Mock the service for testing

	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)

	data := map[string]interface{}{
		"username": "testuser",
		"password": "password123",
	}

	MockJsonPost(ctx, data)

	userHandler.LoginUser(ctx)

	if w.Code != 200 {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	if w.Body.String() != `{"message":"User logged in successfully","user":{"id":0,"username":"testuser","password":"password123","email":"mail@mail.com"}}` {
		t.Errorf("Expected response body to be 'User logged in successfully', got '%s'", w.Body.String())
	}
	t.Log("Response Body:", w.Body.String())
}

func TestUserHandler_LogoutUser(t *testing.T) {
	userHandler := NewUserHandler(&MockUserService{}) // Mock the service for testing

	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)

	// Simulate setting userID in context
	ctx.Set("userID", 1)

	userHandler.LogoutUser(ctx)

	if w.Code != 200 {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}
	if w.Body.String() != `{"message":"User logged out successfully"}` {
		t.Errorf("Expected response body to be 'User logged out successfully', got '%s'", w.Body.String())
	}
	t.Log("Response Body:", w.Body.String())
}
