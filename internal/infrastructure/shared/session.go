package shared

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	shared "github.com/n3xtchen/gin-3at/internal/service/shared"
)

type CookieSession struct {
	context *gin.Context
}

type CookieSessionStarter struct {
	name  string
	store cookie.Store
}

func (s *CookieSessionStarter) Start(ctx *gin.Context) shared.Session {
	return NewCookieSession(ctx, s.name, s.store)
}

// 从 context.Context 获取 gin.Context（需在 handler 里注入）
// NewCookieSessionContext creates a new CookieSessionStore instance.
func NewCookieSession(ctx *gin.Context, name string, store cookie.Store) shared.Session {
	sessions.Sessions(name, store)(ctx)
	return &CookieSession{
		ctx,
	}
}

func (s *CookieSession) Save(data *shared.SessionData) error {
	// Implement logic to save user session, e.g., set cookie
	session := sessions.Default(s.context)
	session.Set("userID", data.UserID)
	if err := session.Save(); err != nil {
		return err
	}
	return nil // Replace with actual implementation
}

func (s *CookieSession) Get() (*shared.SessionData, error) {
	// Ensure ctx is of type *gin.Context
	// Implement logic to retrieve user session, e.g., get cookie
	session := sessions.Default(s.context)
	userID := session.Get("userID")
	if userID == nil {
		return nil, nil // No session found
	}
	return &shared.SessionData{UserID: userID.(string)}, nil // Replace with actual implementation
}

func (s *CookieSession) Clear() error {
	// Implement logic to clear user session, e.g., delete cookie
	session := sessions.Default(s.context)
	session.Clear()
	return nil // Replace with actual implementation
}
