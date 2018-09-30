package dao

import (
	"github.com/go-park-mail-ru/2018_2_vi_studio/restapi/models"
)

type TokenDAO struct {
	data map[int]models.Token
}

func NewTokenDAO() *TokenDAO {
	return &TokenDAO{data: make(map[int]models.Token)}
}

func (md *TokenDAO) GetAll() []models.Token {
	result := make([]models.Token, len(md.data))
	for i, obj := range md.data {
		result[i] = obj
	}
	return result
}

func (md *TokenDAO) Get(id int) (models.Token, bool) {
	result, ok := md.data[id]
	return result, ok
}

func (md *TokenDAO) Save(obj models.Token) {
	md.data[obj.GetId()] = obj
}

func (md *TokenDAO) Remove(id int) {
	delete(md.data, id)
}
