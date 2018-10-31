package sessions

import (
	"github.com/google/uuid"
)

type Session struct {
	AccessToken *uuid.UUID
	UserId      *int
}

func (s Session) IsValid() bool {
	return *s.UserId >= 0
}
