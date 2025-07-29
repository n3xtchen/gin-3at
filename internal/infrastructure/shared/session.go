package shared

import (
	"context"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	e "github.com/n3xtchen/gin-3at/internal/domain/entity"
)

type CookieSessionContext struct {
	// Add any fields needed for session context, e.g., cookie data
	sessionID string
}

func (s *CookieSessionContext) UserID() string {
	return "exampleUserID" // Replace with actual logic to retrieve user ID
}

func (s *CookieSessionContext) save(ctx *gin.Context, user *e.User) error {
	// Implement logic to save user session, e.g., set cookie
	session := sessions.Default(ctx)
	session.Set("userID", user.ID)
	if err := session.Save(); err != nil {
		return err
	}
	return nil // Replace with actual implementation
}

func (s *CookieSessionContext) clear(ctx context.Context) error {
	// Implement logic to clear user session, e.g., delete cookie
	return nil // Replace with actual implementation
}
