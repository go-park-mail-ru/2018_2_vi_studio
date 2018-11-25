package sessions

import (
	"auth/resources/proto"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type SessionService struct {
	Sessions *SessionStorage
}

func NewSessionService(sessions *SessionStorage) *SessionService {
	return &SessionService{
		Sessions: sessions,
	}
}

func (au *SessionService) ByToken(ctx context.Context, token *proto.UUID) (*proto.Session, error) {
	tokenUUID, _ := uuid.Parse(token.Value)
	session, err := au.Sessions.ByToken(tokenUUID)
	if err != nil {
		return &proto.Session{}, err
	}

	return &proto.Session{
		AccessToken: &proto.UUID{Value: session.AccessToken.String()},
		UserUID:     &proto.UUID{Value: session.UserUID.String()},
	}, nil
}

func (au *SessionService) ByUserUID(ctx context.Context, userUID *proto.UUID) (*proto.Session, error) {
	userUUID, _ := uuid.Parse(userUID.Value)
	session, err := au.Sessions.ByUserUId(userUUID)
	if err != nil {
		return &proto.Session{}, err
	}

	return &proto.Session{
		AccessToken: &proto.UUID{Value: session.AccessToken.String()},
		UserUID:     &proto.UUID{Value: session.UserUID.String()},
	}, nil
}

func (au *SessionService) Add(ctx context.Context, session *proto.Session) (*proto.Nothing, error) {
	accessToken, _ := uuid.Parse(session.AccessToken.Value)
	userUID, _ := uuid.Parse(session.UserUID.Value)

	err := au.Sessions.Add(&Session{
		AccessToken: &accessToken,
		UserUID:     &userUID,
	})

	return &proto.Nothing{}, err
}

func (au *SessionService) RemoveByToken(ctx context.Context, token *proto.UUID) (*proto.Nothing, error) {
	tokenUUID, _ := uuid.Parse(token.Value)

	err := au.Sessions.RemoveByToken(&tokenUUID)

	return &proto.Nothing{}, err
}
