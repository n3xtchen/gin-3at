package shared

type SessionData struct {
	UserID int
}

type Session interface {
	Save(data *SessionData) error
	Get() (*SessionData, error)
	Clear() error
}
