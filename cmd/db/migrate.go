package main

import (
	"log"

	conf "github.com/n3xtchen/gin-3at/internal/config"
	m "github.com/n3xtchen/gin-3at/internal/model"
	"github.com/n3xtchen/gin-3at/internal/pkg"
)

// todo: 后面使用 golang-migrate/migrate 来做迁移
func main() {

	appConf := conf.InitConfig()

	// Initialize Database
	mysqlConf := appConf.MySQL
	db := pkg.InitMySQL(mysqlConf.User, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.Database)

	db.AutoMigrate(&m.Category{}, &m.Product{}, &m.Order{}, &m.OrderItem{}, &m.Address{})

	log.Println("Database setup completed successfully.")
}
