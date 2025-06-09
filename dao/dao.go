package dao

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	conf "github.com/n3xtchen/gin-3at/config"
)

var (
	_db *gorm.DB
)

func InitMySQL() {
	mysqlConf := conf.Conf.MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8mb4",
		mysqlConf.User, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	_db = db
}

func NewDBClient() *gorm.DB {
	db := _db
	return db
}
