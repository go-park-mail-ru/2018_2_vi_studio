package models

import "github.com/google/uuid"

type Token struct {
	UserId int
	UUID uuid.UUID
}

func (t *Token) GetId() int {
	return t.UserId
}
