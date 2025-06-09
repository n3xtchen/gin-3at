package dao

import (
	"fmt"
	"log"
	"testing"

	"github.com/joho/godotenv"

	conf "github.com/n3xtchen/gin-3at/config"
)

func TestDBConf(t *testing.T) {
	if err := godotenv.Load("../test.env"); err != nil {
		log.Println(err)
		log.Println("No .env file found or failed to load.")
	}

	conf.InitConfig()
	InitMySQL()

	tables := make([]string, 0)
	db := NewDBClient()
	db.Raw("SHOW TABLES").Scan(&tables)
	fmt.Println(tables)
}
