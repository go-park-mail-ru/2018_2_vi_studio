package sessions

import "github.com/google/uuid"

type SessionStorage interface {
	ByToken(token uuid.UUID) (Session, bool)
	ByUserId(id int) (Session, bool)
	Save(obj Session)
	RemoveByToken(token uuid.UUID)
}
