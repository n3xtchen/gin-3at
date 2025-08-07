package router

import (
	"os"

	"github.com/gin-gonic/gin"

	docs "github.com/n3xtchen/gin-3at/docs"    // swagger docs
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	"github.com/n3xtchen/gin-3at/internal/handler"

	"github.com/n3xtchen/gin-3at/internal/middleware"
)

func SetupRouter(middlewares []gin.HandlerFunc, userHandler *handler.UserHandler, orderHandler *handler.OrderHandler, categoryHandler *handler.CategoryHandler, productHandler *handler.ProductHandler) *gin.Engine {

	router := gin.Default()

	// session
	// router.Use(sessions.Sessions("session", store))
	router.Use(middlewares...)

	{
		v1 := router.Group("api/v1")

		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		// User
		v1.POST("/users/register", userHandler.RegisterUser)
		v1.POST("/users/login", userHandler.LoginUser)
		v1.GET("/users/logout", userHandler.LogoutUser)
		v1.GET("/user/reset_password", userHandler.ResetPassword)

		// Categories
		v1.GET("/categories/:id", categoryHandler.GetCategoryByID)
		v1.GET("/categories", categoryHandler.GetAllCategories)

		// Products
		v1.GET("/products/:id", productHandler.GetProductByID)
		v1.GET("/products", productHandler.GetAllProducts)

		authed := v1.Group("/")
		authed.Use(middleware.AuthRequired())
		{
			// orders
			authed.POST("/orders", orderHandler.Save)
		}
	}

	if os.Getenv("ENV") != "production" {
		docs.SwaggerInfo.BasePath = "/api/v1"
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return router
}
