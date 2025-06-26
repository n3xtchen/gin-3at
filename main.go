package main

import (
	"github.com/gin-contrib/sessions/cookie"

	conf "github.com/n3xtchen/gin-3at/config"
	"github.com/n3xtchen/gin-3at/internal/dao"
	"github.com/n3xtchen/gin-3at/internal/handler"
	"github.com/n3xtchen/gin-3at/internal/pkg"
	"github.com/n3xtchen/gin-3at/internal/router"
	"github.com/n3xtchen/gin-3at/internal/service"
)

func main() {

	appConf := conf.InitConfig()

	// Initialize Database
	mysqlConf := appConf.MySQL
	db := pkg.InitMySQL(mysqlConf.User, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.Database)

	// dao
	orderRepository := dao.NewOrderDao(db)

	// service
	orderService := service.NewOrderService(db, orderRepository)

	// handler
	orderHandler := handler.NewOrderHandler(orderService)

	// session
	store := cookie.NewStore([]byte(appConf.Secret))

	// router
	r := router.SetupRouter(store, orderHandler)

	r.Run(":" + appConf.Port) // 监听并在 0.0.0.0:8080 上启动服务
}
