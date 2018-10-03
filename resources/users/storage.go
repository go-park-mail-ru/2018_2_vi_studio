package users

type UserStorage interface {
	All() []User
	ById(id int) (User, bool)
	ByNickname(nickname string) (User, bool)
	Add(obj User) error
	Remove(id int)
}
