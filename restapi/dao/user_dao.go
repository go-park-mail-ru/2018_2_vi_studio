package dao

import (
	"errors"
	"github.com/go-park-mail-ru/2018_2_vi_studio/restapi/models"
)

type UserDAO struct {
	data map[int]models.User
	maxId int
}

func NewUserDAO() *UserDAO {
	return &UserDAO{data: make(map[int]models.User), maxId: 1}
}

func (md *UserDAO) GetAll() []models.User {
	result := make([]models.User, 0, len(md.data))
	for _, obj := range md.data {
		result = append(result, obj)
	}
	return result
}

func (md *UserDAO) Get(id int) models.User {
	return md.data[id]
}

func (md *UserDAO) Save(obj models.User) error {
	for _, user := range md.data {
		if user.Nickname == obj.Nickname {
			return errors.New("nickname conflict")
		}
		if user.Email == obj.Email {
			return errors.New("email conflict")
		}
	}

	if obj.Id == 0 {
		obj.Id = md.maxId
		md.maxId += 1
	}

	md.data[obj.Id] = obj

	return nil
}

func (md *UserDAO) Remove(id int) {
	delete(md.data, id)
}
