package dao

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/gorm"

	conf "github.com/n3xtchen/gin-3at/config"
	m "github.com/n3xtchen/gin-3at/internal/model"
	s "github.com/n3xtchen/gin-3at/internal/seed"
)

// DB is the global database client
var db *gorm.DB

func setupDB() {
	log.Println("setup DB")
	if err := godotenv.Load("../../test.env"); err != nil {
		log.Println(err)
		log.Println("No .env file found or failed to load.")
	}

	testConf := conf.InitConfig()
	mysqlConf := testConf.MySQL
	db = InitMySQL(mysqlConf.User, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.Database)

	db.AutoMigrate(&m.Category{}, &m.Product{}, &m.Order{}, &m.OrderItem{}, &m.Address{})

	// Seed Categories
	categoryDao := NewCategoryDao(db)
	for _, category := range s.CategorySeed {
		if err := categoryDao.CreateCategory(category); err != nil {
			log.Fatalf("Failed to seed category: %v", err)
		}
	}

	// Seed ProductSeed
	productDao := NewProductDao(db)
	for _, product := range s.ProductSeed {
		if err := productDao.CreateProduct(product); err != nil {
			log.Fatalf("Failed to seed product: %v", err)
		}
	}

	// Seed AddressSeed
	addressDao := NewAddressDao(db)
	for _, address := range s.AddressSeed {
		if err := addressDao.CreateAddress(address); err != nil {
			log.Fatalf("Failed to seed address: %v", err)
		}
	}

	// seed OrderSeed
	// orderDao := NewOrderDao(db)
	for _, order := range s.OrderSeed {
		if err := db.Create(&order).Error; err != nil {
			log.Fatalf("Failed to seed order: %v", err)
		}
	}

	// Seed OrderItemSeed
	orderItemDao := NewOrderItemDao(db)
	for _, orderItem := range s.OrderItemSeed {
		if err := orderItemDao.Create(orderItem); err != nil {
			log.Fatalf("Failed to seed order item: %v", err)
		}
	}

	log.Println("Database setup completed successfully.")
}

func teardownDB() {
	// seed data is not deleted, but you can drop tables if needed
	log.Println("teardown DB")
	if err := db.Migrator().DropTable(&m.Product{}, &m.Category{}, &m.Order{}, &m.OrderItem{}, &m.Address{}); err != nil {
		log.Fatalf("Failed to drop tables: %v", err)
	}
	log.Println("Tables dropped successfully.")
}

func TestMain(m *testing.M) {
	setupDB()
	code := m.Run()
	teardownDB()
	os.Exit(code)
}

func TestDBConf(t *testing.T) {
	tables := make([]string, 0)
	db.Raw("SHOW TABLES").Scan(&tables)
	log.Println(tables)
}
