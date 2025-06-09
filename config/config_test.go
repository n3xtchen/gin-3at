package config

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func TestConfig(t *testing.T) {
	if err := godotenv.Load("../test.env"); err != nil {
		log.Println(err)
		log.Println("No .env file found or failed to load.")
	}
	InitConfig()
	log.Println(Conf)
}
