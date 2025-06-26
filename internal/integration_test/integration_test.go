package test

import (
	"os"
	"testing"

	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/n3xtchen/gin-3at/internal/dao"
	"github.com/n3xtchen/gin-3at/internal/handler"
	"github.com/n3xtchen/gin-3at/internal/router"
	"github.com/n3xtchen/gin-3at/internal/service"
	"github.com/n3xtchen/gin-3at/internal/testutils"
)

// DB is the global database client
var (
	db *gorm.DB
	r  *gin.Engine
)

func setupTestServer() *gin.Engine {

	appConf := testutils.InitTestConfig()

	db = testutils.InitTestDB(&appConf.MySQL)

	// Seed Categories
	testutils.SeedTestDB(db)

	// dao
	orderRepository := dao.NewOrderDao(db.Debug())

	// service
	orderService := service.NewOrderService(db, orderRepository)

	store := cookie.NewStore([]byte(appConf.Secret))
	orderHandler := handler.NewOrderHandler(orderService)

	r := router.SetupRouter(store, orderHandler)

	return r
}

func TestMain(m *testing.M) {
	r = setupTestServer()
	code := m.Run()
	testutils.TeardownDB(db)
	os.Exit(code)
}
