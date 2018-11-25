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
	err := ss.db.QueryRow("SELECT user_uid, token FROM sessions WHERE token = $1", token).Scan(&result.UserUID, &result.AccessToken)
	if err == nil {
		return &result, nil
	}

	return nil, ErrUnknown
}

func (ss *SessionStorage) ByUserUId(id uuid.UUID) (*Session, error) {
	var result Session
	err := ss.db.QueryRow("SELECT user_uid, token FROM sessions WHERE user_uid = $1", id).Scan(&result.UserUID, &result.AccessToken)
	if err == nil {
		return &result, nil
	}

	if err.Error() == "sql: no rows in result set" {
		return nil, ErrNotFound
	}

	return nil, ErrUnknown
}

func (ss *SessionStorage) Add(session *Session) error {
	_, err := ss.db.Exec("INSERT INTO sessions (user_uid, token) VALUES ($1, $2) ON CONFLICT(user_uid) DO UPDATE SET token = $2", session.UserUID, session.AccessToken)
	if err == nil {
		return nil
	}

	return ErrUnknown
}

func (ss *SessionStorage) RemoveByToken(token *uuid.UUID) error {
	_, err := ss.db.Exec("DELETE FROM sessions WHERE token = $1", *token)
	if err == nil {
		return nil
	}

	return ErrUnknown
}
