package middleware

import (
	"github.com/gin-contrib/sessions"
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
		c.Next()
	}
}

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
