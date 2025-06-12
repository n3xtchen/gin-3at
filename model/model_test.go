package model

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/gorm"

	conf "github.com/n3xtchen/gin-3at/config"
	dao "github.com/n3xtchen/gin-3at/dao"
)

var (
	DB *gorm.DB
)

func setupDB() {
	log.Println("setup DB")
	if err := godotenv.Load("../test.env"); err != nil {

		log.Println(err)
		log.Println("No .env file found or failed to load.")
	}

	conf.InitConfig()
	dao.InitMySQL()

	DB = dao.NewDBClient()
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
