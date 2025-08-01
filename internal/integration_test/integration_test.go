package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/n3xtchen/gin-3at/internal/dao"
	"github.com/n3xtchen/gin-3at/internal/handler"
	"github.com/n3xtchen/gin-3at/internal/router"
	"github.com/n3xtchen/gin-3at/internal/seed"
	"github.com/n3xtchen/gin-3at/internal/service"
	"github.com/n3xtchen/gin-3at/internal/testutils"

	shared "github.com/n3xtchen/gin-3at/internal/infrastructure/shared"
)

// DB is the global database client
var (
	db          *gorm.DB
	r           *gin.Engine
	loginCookie []*http.Cookie
)

func loginAndGetCookie(username, password string) []*http.Cookie {
	loginData := map[string]string{
		"username": username,
		"password": password,
	}

	jsonLoginBody, _ := json.Marshal(loginData)
	reqLogin, _ := http.NewRequest("POST", "/api/v1/users/login", bytes.NewBuffer(jsonLoginBody))
	reqLogin.Header.Set("Content-Type", "application/json")
	respLogin := httptest.NewRecorder()
	r.ServeHTTP(respLogin, reqLogin)
	return respLogin.Result().Cookies() // Store the login cookie for logout test
}

func setupTestServer() *gin.Engine {

	appConf := testutils.InitTestConfig()

	db = testutils.InitTestDB(&appConf.MySQL)

	// Seed Categories
	testutils.SeedTestDB(db)

	store := cookie.NewStore([]byte(appConf.Secret))
	sessionInitor := shared.NewCookieSession("session", store)

	// dao
	orderRepository := dao.NewOrderDao(db.Debug())
	userRepository := dao.NewUserDao(db.Debug())
	categoryRepository := dao.NewCategoryDao(db.Debug())
	productRepository := dao.NewProductDao(db.Debug())

	// service
	orderService := service.NewOrderService(db, orderRepository)
	userService := service.NewUserService(userRepository)
	categoryService := service.NewCategoryService(categoryRepository)
	productService := service.NewProductService(productRepository)

	// handler
	orderHandler := handler.NewOrderHandler(orderService)
	userHandler := handler.NewUserHandler(userService)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	productHandler := handler.NewProductHandler(productService)

	r := router.SetupRouter(sessionInitor, userHandler, orderHandler, categoryHandler, productHandler)

	return r
}

func TestMain(m *testing.M) {
	r = setupTestServer()
	user := seed.UserSeed[0] // Use the first user from the seed data for login
	loginCookie = loginAndGetCookie(user.Username, user.Password)
	code := m.Run()
	testutils.TeardownDB(db)
	os.Exit(code)
}
