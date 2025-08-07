package middleware

import (
	"github.com/gin-gonic/gin"

	shared "github.com/n3xtchen/gin-3at/internal/service/shared"
)

const (
	DefaultKey = "session-3at"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := c.MustGet(DefaultKey).(*shared.SessionData)
		if session == nil {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
