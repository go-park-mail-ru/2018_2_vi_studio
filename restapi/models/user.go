package models

type User struct {
	Id int `json:"id,omitempty"`
	Nickname string `json:"nickname"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func (u *User) GetId() int {
	return u.Id
}
