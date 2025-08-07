package shared

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	shared "github.com/n3xtchen/gin-3at/internal/service/shared"
)

const (
	DefaultKey = "session-3at"
)

type CookieSession struct {
	session gin.HandlerFunc
	context *gin.Context
}

// NewCookieSessionContext creates a new CookieSessionStore instance.
// func NewCookieSession(name string, store cookie.Store) shared.Session {
func NewCookieSession(name string, store cookie.Store) *CookieSession {
	session := sessions.Sessions(name, store)
	return &CookieSession{
		session: session,
		context: nil,
	}
}

func (s *CookieSession) MiddleWares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		s.session,
		func(c *gin.Context) {
			s.context = c
			data, _ := s.Get()
			if data != nil {
				c.Set(DefaultKey, data)
			}
			c.Next()
		},
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
	return &shared.SessionData{UserID: userID.(int)}, nil // Replace with actual implementation
}

func (s *CookieSession) Clear() error {
	// Implement logic to clear user session, e.g., delete cookie
	session := sessions.Default(s.context)
	session.Clear()
	return nil // Replace with actual implementation
}
