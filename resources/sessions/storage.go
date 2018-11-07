package sessions

import (
	"database/sql"
	"github.com/google/uuid"
)

type SessionStorage struct {
	db *sql.DB
}

func NewSessionStorage(db *sql.DB) *SessionStorage {
	return &SessionStorage{db: db}
}

func (ss *SessionStorage) ByToken(token uuid.UUID) (*Session, error) {
	var result Session
	err := ss.db.QueryRow("SELECT user_id, token FROM sessions WHERE token = $1", token).Scan(&result.UserId, &result.AccessToken)
	if err == nil {
		return &result, nil
	}

	return nil, ErrUnknown
}

func (ss *SessionStorage) ByUserId(id int) (*Session, error) {
	var result Session
	err := ss.db.QueryRow("SELECT user_id, token FROM sessions WHERE user_id = $1", id).Scan(&result.UserId, &result.AccessToken)
	if err == nil {
		return &result, nil
	}

	if err.Error() == "sql: no rows in result set" {
		return nil, ErrNotFound
	}

	return nil, ErrUnknown
}

func (ss *SessionStorage) Add(obj Session) error {
	_, err := ss.db.Exec("INSERT INTO sessions (user_id, token) VALUES ($1, $2) ON CONFLICT(user_id) DO UPDATE SET token = $2", obj.UserId, obj.AccessToken)
	if err == nil {
		return nil
	}

	return ErrUnknown
}

func (ss *SessionStorage) RemoveByToken(token uuid.UUID) error {
	_, err := ss.db.Exec("DELETE FROM sessions WHERE token = $1", token)
	if err == nil {
		return nil
	}

	return ErrUnknown
}
