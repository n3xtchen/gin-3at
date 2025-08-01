package shared

type SessionData struct {
	UserID string
}

type Session interface {
	Save(data *SessionData) error
	Get() (*SessionData, error)
	Clear() error
}
