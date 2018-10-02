package dao

import (
	"github.com/google/uuid"
)

type TokenDAO struct {
	data map[uuid.UUID]int
}

func NewTokenDAO() *TokenDAO {
	return &TokenDAO{data: make(map[uuid.UUID]int)}
}

func (md *TokenDAO) GetUserId(token uuid.UUID) (int, bool) {
	result, ok := md.data[token]
	return result, ok
}

func (md *TokenDAO) GetToken(userId int) (uuid.UUID, bool) {
	for token, id := range md.data {
		if id == userId {
			return token, true
		}
	}

	return uuid.UUID{}, false
}

func (md *TokenDAO) Set(token uuid.UUID, userId int) {
	md.data[token] = userId
}

func (md *TokenDAO) Remove(token uuid.UUID) {
	delete(md.data, token)
}

