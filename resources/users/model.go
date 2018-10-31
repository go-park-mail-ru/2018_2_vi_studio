package users

type User struct {
	Id       *int    `json:"id,omitempty"`
	Nickname *string `json:"nickname"`
	Email    *string `json:"email"`
	Password *string `json:"password,omitempty"`
	Points   *int    `json:"points,omitempty"`
}

func (u User)IsValid() bool {
	return *u.Id >= 0 &&
		len(*u.Nickname) > 3 && len(*u.Nickname) < 32 &&
		len(*u.Email) > 3 && len(*u.Email) < 64 &&
		len(*u.Password) > 3 && len(*u.Password) < 32 &&
		*u.Points >= 0
}

type Leader struct {
	Nickname string `json:"nickname"`
	Points   int    `json:"points"`
}

func (l Leader)IsValid() bool {
	return len(l.Nickname) > 3 && len(l.Nickname) < 32 && l.Points >= 0
}

type UserAuth struct {
	Nickname *string `json:"nickname"`
	Password *string `json:"password"`
}

func (u UserAuth)IsValid() bool {
	return len(*u.Nickname) > 3 && len(*u.Nickname) < 32 &&
		len(*u.Password) > 3 && len(*u.Password) < 32
}