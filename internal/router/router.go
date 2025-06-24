package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

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

	return router
}
