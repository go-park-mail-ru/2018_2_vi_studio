package sessions

import (
	"github.com/google/uuid"
)

type Session struct {
	AccessToken *uuid.UUID
	UserUID     *uuid.UUID
}

func (s Session) IsValid() bool {
	return true
}