package middleware

import (
	"context"
	"github.com/google/uuid"
	"main/proto"
	"net/http"
)

const sessionKey = "SESSION"

func Session(ctx context.Context) (proto.Session, bool) {
	session := ctx.Value(sessionKey)
	if session == nil {
		return proto.Session{}, false
	} else {
		return session.(proto.Session), true
	}
}

func WithSession(ctx context.Context, session proto.Session) context.Context {
	return context.WithValue(ctx, sessionKey, session)
}

func SessionMW(handlerFunc http.HandlerFunc, service proto.SessionServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if accessToken, err := r.Cookie("access_token"); err == nil {
			if session, err := service.ByToken(context.Background(), &proto.UUID{Value: accessToken.Value}); err == nil {
				ctx = WithSession(ctx, *session)

			}
		} else {
			var anonymToken uuid.UUID
			if anonymTokenString, err := r.Cookie("access_token"); err == nil {
				anonymToken, _ = uuid.Parse(anonymTokenString.Value)
			} else {
				anonymToken, _ = uuid.NewRandom()
				http.SetCookie(w, &http.Cookie{
					Name:     "anonym_token",
					Value:    anonymToken.String(),
					HttpOnly: true,
					Secure:   true,
					Path:     "/",
				})
			}
		}

		handlerFunc(w, r.WithContext(ctx))
	}
}
