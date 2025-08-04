package middleware

import (
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
