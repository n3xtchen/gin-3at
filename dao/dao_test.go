package dao

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"

	conf "github.com/n3xtchen/gin-3at/config"
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
}

func teardownDB() {
	log.Println("teardown DB")
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
