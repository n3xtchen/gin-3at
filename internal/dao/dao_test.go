package dao

import (
	"log"
	"os"
	"testing"

	"gorm.io/gorm"

	"github.com/n3xtchen/gin-3at/internal/testutils"
)

// DB is the global database client
var db *gorm.DB

func TestMain(m *testing.M) {
	appConf := testutils.InitTestConfig()
	db = testutils.InitTestDB(&appConf.MySQL)
	testutils.SeedTestDB(db)
	log.Println("Database setup completed successfully.")
	if db == nil {
		log.Fatal("Failed to initialize test database")
	}
	code := m.Run()
	testutils.TeardownDB(db)
	os.Exit(code)
}

func TestDBConf(t *testing.T) {
	tables := make([]string, 0)
	db.Raw("SHOW TABLES").Scan(&tables)
	log.Println(tables)
}
