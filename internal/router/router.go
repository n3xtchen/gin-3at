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

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("userID") == nil {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func SetupRouter(store cookie.Store, userHandler *handler.UserHandler, orderHandler *handler.OrderHandler) *gin.Engine {

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

		// User
		v1.POST("/users/register", userHandler.RegisterUser)
		v1.GET("/users/login", userHandler.LoginUser)
		v1.GET("/users/logout", userHandler.LogoutUser)
		v1.GET("/user/reset_password", userHandler.ResetPassword)

		authed := v1.Group("/")
		authed.Use(AuthRequired())
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
