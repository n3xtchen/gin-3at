package router

import (
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	docs "github.com/n3xtchen/gin-3at/docs"    // swagger docs
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	"github.com/n3xtchen/gin-3at/internal/handler"
)

func SetupRouter(store cookie.Store, orderHandler *handler.OrderHandler) *gin.Engine {

	router := gin.Default()

	// session
	router.Use(sessions.Sessions("session", store))

	{
		v1 := router.Group("api/v1")

		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		// orders
		v1.POST("/orders", orderHandler.Save)
	}

	if os.Getenv("ENV") != "production" {
		docs.SwaggerInfo.BasePath = "/api/v1"
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return router
}
