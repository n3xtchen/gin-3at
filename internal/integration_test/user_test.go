package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	LoginCookie []*http.Cookie
)

func TestUserIntegration(t *testing.T) {

	t.Logf("Test User Integration")

	// Test Register User
	registerData := map[string]string{
		"username": "testuser",
		"password": "password123",
		"email":    "mail@mail.com",
	}

	jsonRegisterBody, _ := json.Marshal(registerData)
	reqRegister, _ := http.NewRequest("POST", "/api/v1/users/register", bytes.NewBuffer(jsonRegisterBody))
	reqRegister.Header.Set("Content-Type", "application/json")
	respRegister := httptest.NewRecorder()
	r.ServeHTTP(respRegister, reqRegister)
	if respRegister.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", respRegister.Code)
	}
	t.Logf("Register Response: %s", respRegister.Body.String())
}

func TestUserHandler_LoginUser(t *testing.T) {
	// Test Login User
	loginData := map[string]string{
		"username": "testuser",
		"password": "password123",
	}

	jsonLoginBody, _ := json.Marshal(loginData)
	reqLogin, _ := http.NewRequest("POST", "/api/v1/users/login", bytes.NewBuffer(jsonLoginBody))
	reqLogin.Header.Set("Content-Type", "application/json")
	respLogin := httptest.NewRecorder()
	r.ServeHTTP(respLogin, reqLogin)
	if respLogin.Code != http.StatusOK {
		t.Logf("Login Response: %s", respLogin.Body.String())
		t.Fatalf("expected status 200, got %d", respLogin.Code)
	}

	LoginCookie = respLogin.Result().Cookies() // Store the login cookie for logout test
	t.Logf("Login Response: %s", respLogin.Body.String())
}

func TestUserHandler_LogoutUser(t *testing.T) {
	// Test Logout User
	logoutReq, _ := http.NewRequest("GET", "/api/v1/users/logout", nil)
	// Add the login cookie to the Request
	for _, cookie := range LoginCookie {
		logoutReq.AddCookie(cookie)
	}
	respLogout := httptest.NewRecorder()
	r.ServeHTTP(respLogout, logoutReq)
	if respLogout.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", respLogout.Code)
	}
	if respLogout.Body.String() != `{"message":"User logged out successfully"}` {
		t.Errorf("Expected response body to be 'User logged out successfully', got '%s'", respLogout.Body.String())
	}

	t.Logf("Logout Response: %s", respLogout.Body.String())
}
