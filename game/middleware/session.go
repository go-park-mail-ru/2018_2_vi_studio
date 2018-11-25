package middleware

import (
	"game/resources/sessions"
	"context"
	"github.com/google/uuid"
	"net/http"
)

const sessionKey = "SESSION"

func Session(ctx context.Context) (sessions.Session, bool) {
	session := ctx.Value(sessionKey)
	if session == nil {
		return sessions.Session{}, false
	} else {
		return session.(sessions.Session), true
	}
}

func WithSession(ctx context.Context, session sessions.Session) context.Context {
	return context.WithValue(ctx, sessionKey, session)
}

func SessionMW(handlerFunc http.HandlerFunc, sessions *sessions.SessionStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if accessTokenString, err := r.Cookie("access_token"); err == nil {
			if accessToken, err := uuid.Parse(accessTokenString.Value); err == nil {
				if session, err := sessions.ByToken(accessToken); err == nil {
					ctx = WithSession(ctx, *session)
				}
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
					Secure:   false,
					Path:     "/",
				})
			}
		}

		handlerFunc(w, r.WithContext(ctx))
	}
}