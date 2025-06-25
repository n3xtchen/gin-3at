package test

import (
	"log"
	"os"
	"testing"

	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	conf "github.com/n3xtchen/gin-3at/config"
	"github.com/n3xtchen/gin-3at/internal/dao"
	"github.com/n3xtchen/gin-3at/internal/handler"
	m "github.com/n3xtchen/gin-3at/internal/model"
	"github.com/n3xtchen/gin-3at/internal/router"
	s "github.com/n3xtchen/gin-3at/internal/seed"
	"github.com/n3xtchen/gin-3at/internal/service"
)

// DB is the global database client
var (
	db *gorm.DB
	r  *gin.Engine
)

func setupTestServer() *gin.Engine {
	if err := godotenv.Load("../../test.env"); err != nil {
		log.Println(err)
		log.Println("No .env file found or failed to load.")
	}

	testConf := conf.InitConfig()
	mysqlConf := testConf.MySQL
	db = dao.InitMySQL(mysqlConf.User, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.Database)

	db.AutoMigrate(&m.Category{}, &m.Product{}, &m.Order{}, &m.OrderItem{}, &m.Address{})

	// Seed Categories
	categoryDao := dao.NewCategoryDao(db)
	for _, category := range s.CategorySeed {
		if err := categoryDao.CreateCategory(category); err != nil {
			log.Fatalf("Failed to seed category: %v", err)
		}
	}

	// Seed ProductSeed
	productDao := dao.NewProductDao(db)
	for _, product := range s.ProductSeed {
		if err := productDao.CreateProduct(product); err != nil {
			log.Fatalf("Failed to seed product: %v", err)
		}
	}

	// dao
	orderRepository := dao.NewOrderDao(db.Debug())

	// service
	orderService := service.NewOrderService(db, orderRepository)

	store := cookie.NewStore([]byte(testConf.Secret))
	orderHandler := handler.NewOrderHandler(orderService)

	r := router.SetupRouter(store, orderHandler)

	return r
}

func tearDownTestServer() {
	log.Println("Tearing down test server...")
	// Here you can add any cleanup logic if needed, like closing database connections.
	if err := db.Migrator().DropTable(&m.Product{}, &m.Category{}, &m.Order{}, &m.OrderItem{}, &m.Address{}); err != nil {
		log.Fatalf("Failed to drop tables: %v", err)
	}
	log.Println("Tables dropped successfully.")
}

func TestMain(m *testing.M) {
	r = setupTestServer()
	code := m.Run()
	tearDownTestServer()
	os.Exit(code)
}
