package users

import (
	"database/sql"
	//"errors"
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

	rows, err := us.db.Query("SELECT nickname, email, points FROM users LIMIT $1 OFFSET $2", limit, offset)
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
		"SELECT u.nickname, u.email, u.points FROM users u JOIN sessions s ON u.id = s.user_id WHERE token = $1",
		token,
	).Scan(&result.Nickname, &result.Email, &result.Points)
	if err == nil {
		return &result, nil
	}

	return nil, ErrUnknown
}

func (us *UserStorage) ByNickname(nickname string) (*User, error) {
	var result User
	err := us.db.QueryRow(
		"SELECT nickname, email, points FROM users WHERE nickname = $1",
		nickname,
	).Scan(&result.Nickname, &result.Email, &result.Points)
	if err == nil {
		return &result, nil
	}

	return nil, ErrUnknown
}

func (us *UserStorage) ById(id int) (*User, error) {
	var result User
	err := us.db.QueryRow(
		"SELECT nickname, email, points FROM users WHERE id = $1",
		id,
	).Scan(&result.Nickname, &result.Email, &result.Points)
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

func (us *UserStorage) Auth(user UserAuth) (*User, error) {
	var result User
	err := us.db.QueryRow(
		"SELECT id, nickname, email, points FROM users WHERE nickname = $1 and password = $2",
		user.Nickname, user.Password,
	).Scan(&result.Id, &result.Nickname, &result.Email, &result.Points)

	if err == nil {
		return &result, nil
	}

	if err.Error() == "sql: no rows in result set" {
		return nil, ErrNotFound
	}

	return nil, ErrUnknown
}
