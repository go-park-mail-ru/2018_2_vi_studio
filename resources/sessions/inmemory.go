package sessions

import (
	"github.com/google/uuid"
)

type SessionIMS struct {
	uuidMap   map[uuid.UUID]*Session
	idMap map[int]*Session
}

func NewSessionIMS() *SessionIMS {
	return &SessionIMS{
		uuidMap: make(map[uuid.UUID]*Session),
		idMap: make(map[int]*Session),
	}
}

func (md *SessionIMS) ByToken(token uuid.UUID) (Session, bool) {
	if result, ok := md.uuidMap[token]; ok {
		return *result, true
	}

	return Session{}, false
}

func (md *SessionIMS) ByUserId(id int) (Session, bool) {
	if result, ok := md.idMap[id]; ok {
		return *result, true
	}

	return Session{}, false
}

func (md *SessionIMS) Add(obj Session) {
	md.uuidMap[obj.AccessToken] = &obj
	md.idMap[obj.UserId] = &obj
}

func (md *SessionIMS) RemoveByToken(token uuid.UUID) {
	delete(md.idMap, md.uuidMap[token].UserId)
	delete(md.uuidMap, token)
}
