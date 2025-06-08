package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// session
	store := cookie.NewStore([]byte("something-very-secret"))
	router.Use(sessions.Sessions("session", store))

	{
		v1 := router.Group("api/v1")

		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
