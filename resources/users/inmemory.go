package users

import (
	"errors"
)

type UserIMS struct {
	data  map[int]User
	maxId int
}

func NewUserIMS() *UserIMS {
	return &UserIMS{data: make(map[int]User), maxId: 1}
}

func (md *UserIMS) All() []User {
	result := make([]User, 0, len(md.data))
	for _, obj := range md.data {
		result = append(result, obj)
	}
	return result
}

func (md *UserIMS) ById(id int) (User, bool) {
	result, ok := md.data[id]
	return result, ok
}

func (md *UserIMS) ByNickname(nickname string) (User, bool) {
	for _, user := range md.data {
		if user.Nickname == nickname {
			return user, true
		}
	}
	return User{}, false
}

func (md *UserIMS) Add(obj User) error {
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

func (md *UserIMS) Remove(id int) {
	delete(md.data, id)
}
