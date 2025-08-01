package main

import (
	"github.com/gin-contrib/sessions/cookie"

	conf "github.com/n3xtchen/gin-3at/internal/config"
	"github.com/n3xtchen/gin-3at/internal/dao"
	"github.com/n3xtchen/gin-3at/internal/handler"
	"github.com/n3xtchen/gin-3at/internal/pkg"
	"github.com/n3xtchen/gin-3at/internal/router"
	"github.com/n3xtchen/gin-3at/internal/service"

	shared "github.com/n3xtchen/gin-3at/internal/infrastructure/shared"
)

// @title           Gin 3at
// @version         1.0
// @description     This is a server for Gin 3at API.
// @termsOfService  http://swagger.io/terms/

// @contact.name   n3xtchen
// @contact.url    http://n3xtchen.github.io/n3xtchen
// @contact.email  echenwen@gmail.com

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {

	appConf := conf.InitConfig()

	// Initialize Database
	mysqlConf := appConf.MySQL
	db := pkg.InitMySQL(mysqlConf.User, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.Database)

	// dao
	userRepository := dao.NewUserDao(db)
	orderRepository := dao.NewOrderDao(db)
	categoryRepository := dao.NewCategoryDao(db)
	productRepository := dao.NewProductDao(db)

	// service
	userService := service.NewUserService(userRepository)
	orderService := service.NewOrderService(db, orderRepository)
	categoryService := service.NewCategoryService(categoryRepository)
	productService := service.NewProductService(productRepository)

	// handler
	userHandler := handler.NewUserHandler(userService)
	orderHandler := handler.NewOrderHandler(orderService)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	productHandler := handler.NewProductHandler(productService)

	// session
	store := cookie.NewStore([]byte(appConf.Secret))
	sessionInitor := shared.NewCookieSession("session", store)

	// router
	r := router.SetupRouter(sessionInitor, userHandler, orderHandler, categoryHandler, productHandler)

	r.Run(":" + appConf.Port) // 监听并在 0.0.0.0:8080 上启动服务
}
