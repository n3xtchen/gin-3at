package dao

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"

	conf "github.com/n3xtchen/gin-3at/config"
	m "github.com/n3xtchen/gin-3at/model"
	s "github.com/n3xtchen/gin-3at/seed"
)

func setupDB() {
	log.Println("setup DB")
	if err := godotenv.Load("../test.env"); err != nil {

		log.Println(err)
		log.Println("No .env file found or failed to load.")
	}

	conf.InitConfig()
	InitMySQL()

	DB = NewDBClient()

	DB.AutoMigrate(&m.Category{}, &m.Product{}, &m.Order{}, &m.OrderItem{}, &m.Address{})

	// Seed Categories
	categoryDao := NewCategoryDao()
	for _, category := range s.CategorySeed {
		if err := categoryDao.CreateCategory(category); err != nil {
			log.Fatalf("Failed to seed category: %v", err)
		}
	}

	// Seed ProductSeed
	productDao := NewProductDao()
	for _, product := range s.ProductSeed {
		if err := productDao.CreateProduct(product); err != nil {
			log.Fatalf("Failed to seed product: %v", err)
		}
	}

	// Seed AddressSeed
	addressDao := NewAddressDao()
	for _, address := range s.AddressSeed {
		if err := addressDao.CreateAddress(address); err != nil {
			log.Fatalf("Failed to seed address: %v", err)
		}
	}

	// seed OrderSeed
	orderDao := NewOrderDao()
	for _, order := range s.OrderSeed {
		if err := orderDao.Create(&order).Error; err != nil {
			log.Fatalf("Failed to seed order: %v", err)
		}
	}

	// Seed OrderItemSeed
	orderItemDao := NewOrderItemDao()
	for _, orderItem := range s.OrderItemSeed {
		if err := orderItemDao.CreateOrderItem(orderItem); err != nil {
			log.Fatalf("Failed to seed order item: %v", err)
		}
	}

	log.Println("Database setup completed successfully.")
}

func teardownDB() {
	// seed data is not deleted, but you can drop tables if needed
	log.Println("teardown DB")
	if err := DB.Migrator().DropTable(&m.Product{}, &m.Category{}, &m.Order{}, &m.OrderItem{}, &m.Address{}); err != nil {
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
	DB.Raw("SHOW TABLES").Scan(&tables)
	fmt.Println(tables)
}
