package sessions

import "github.com/google/uuid"

type Session struct {
	AccessToken uuid.UUID
	UserId int
}
