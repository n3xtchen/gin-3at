package testutils

import (
	"log"

	"gorm.io/gorm"

	conf "github.com/n3xtchen/gin-3at/internal/config"
	m "github.com/n3xtchen/gin-3at/internal/model"
	"github.com/n3xtchen/gin-3at/internal/pkg"
	s "github.com/n3xtchen/gin-3at/internal/seed"
)

func InitTestConfig() *conf.Config {
	log.Println("setup config")
	testConf := conf.InitConfig()
	log.Println("Config setup completed successfully.")
	return testConf
}

func InitTestDB(dbConf *conf.MySQLConfig) *gorm.DB {
	db := pkg.InitMySQL(dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Database)

	db.AutoMigrate(&m.User{}, &m.Category{}, &m.Product{}, &m.Order{}, &m.OrderItem{}, &m.Address{})

	log.Println("Database setup completed successfully.")

	return db
}

func SeedTestDB(db *gorm.DB) {

	log.Println("Database seeded with initial data.")

	// Seed UserSeed
	for _, user := range s.UserSeed {
		if err := db.Create(&user).Error; err != nil {
			log.Fatalf("Failed to seed user: %v", err)
		}
	}

	// Seed Categories
	for _, category := range s.CategorySeed {
		if err := db.Create(&category).Error; err != nil {
			log.Fatalf("Failed to seed category: %v", err)
		}
	}

	// Seed ProductSeed
	for _, product := range s.ProductSeed {
		if err := db.Create(&product).Error; err != nil {
			log.Fatalf("Failed to seed product: %v", err)
		}
	}

	// Seed AddressSeed
	for _, address := range s.AddressSeed {
		if err := db.Create(&address).Error; err != nil {
			log.Fatalf("Failed to seed address: %v", err)
		}
	}

	// seed OrderSeed
	for _, order := range s.OrderSeed {
		if err := db.Create(&order).Error; err != nil {
			log.Fatalf("Failed to seed order: %v", err)
		}
	}

	// Seed OrderItemSeed
	for _, orderItem := range s.OrderItemSeed {
		if err := db.Create(&orderItem).Error; err != nil {
			log.Fatalf("Failed to seed order item: %v", err)
		}
	}

}

func TeardownDB(db *gorm.DB) {
	// seed data is not deleted, but you can drop tables if needed
	log.Println("teardown DB")
	if err := db.Migrator().DropTable(&m.Product{}, &m.Category{}, &m.Order{}, &m.OrderItem{}, &m.Address{}, &m.User{}); err != nil {
		log.Fatalf("Failed to drop tables: %v", err)
	}
	log.Println("Tables dropped successfully.")
}
