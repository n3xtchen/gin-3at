package middleware

import (
	"log"

	"github.com/gin-gonic/gin"

	shared "github.com/n3xtchen/gin-3at/internal/service/shared"
)

const (
	DefaultKey = "session-3at"
)

func Session(sessionInitor func(*gin.Context) shared.Session) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessionInitor(c)
		c.Set(DefaultKey, session)
		log.Println("Session initialized:", session)
		log.Println(c)
		c.Next()
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("AuthRequired middleware called")
		log.Println(c)
		// session := c.MustGet(DefaultKey).(shared.Session)
		// _, err := session.Get()
		// if err == nil {
		// 	c.JSON(401, gin.H{"error": "Unauthorized"})
		// 	c.Abort()
		// 	return
		// }
		c.Next()
	}
}
