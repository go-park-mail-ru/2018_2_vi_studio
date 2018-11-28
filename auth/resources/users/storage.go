package users

import (
	"database/sql"
	"github.com/google/uuid"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (us *UserStorage) Leaders(limit int, offset int) (*[]User, error) {
	result := make([]User, 0, limit)

	rows, err := us.db.Query("SELECT nickname, email, points FROM users ORDER BY points DESC LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, ErrUnknown
	}

	for rows.Next() {
		var user User
		err = rows.Scan(&user.Nickname, &user.Email, &user.Points)
		if err != nil {
			return nil, ErrUnknown
		}
		result = append(result, user)
	}

	return &result, nil
}

func (us *UserStorage) BySessionToken(token uuid.UUID) (*User, error) {
	var result User
	err := us.db.QueryRow(
		"SELECT u.nickname, u.email, u.points FROM users u JOIN sessions s ON u.uid = s.user_uid WHERE token = $1",
		token,
	).Scan(&result.Nickname, &result.Email, &result.Points)
	if err == nil {
		return &result, nil
	}

	return nil, ErrUnknown
}

func (us *UserStorage) ByUID(uid uuid.UUID) (*User, error) {
	var result User
	err := us.db.QueryRow(
		"SELECT nickname, email, points, avatar FROM users WHERE uid = $1",
		uid,
	).Scan(&result.Nickname, &result.Email, &result.Points, &result.Avatar)
	if err == nil {
		return &result, nil
	}

	return nil, ErrUnknown
}

func (us *UserStorage) Add(obj User) error {
	_, err := us.db.Exec(
		"INSERT INTO users(nickname, email, password) VALUES ($1, $2, $3)",
		obj.Nickname,
		obj.Email,
		obj.Password,
	)

	if err != nil {
		return ErrUnknown
	}

	return nil
}

func (us *UserStorage) AddAvatar(uid uuid.UUID, avatar string) error {
	_, err := us.db.Exec("UPDATE users SET avatar=$1 WHERE uid=$2", avatar, uid)

	if err != nil {
		return ErrUnknown
	}

	return nil
}

func (us *UserStorage) Auth(user UserAuth) (*User, error) {
	var result User
	err := us.db.QueryRow(
		"SELECT uid, nickname, email, points FROM users WHERE nickname = $1 and password = $2",
		user.Nickname, user.Password,
	).Scan(&result.UID, &result.Nickname, &result.Email, &result.Points)

	if err == nil {
		return &result, nil
	}

	if err.Error() == "sql: no rows in result set" {
		return nil, ErrNotFound
	}

	return nil, ErrUnknown
}

func (us *UserStorage) Count() (int, error) {
	var result int
	err := us.db.QueryRow("SELECT count(*) FROM users").Scan(&result)

	if err != nil {
		return 0, ErrUnknown
	}

	return result, nil
}
