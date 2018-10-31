package middleware

import (
	"2018_2_vi_studio_server/resources/sessions"
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

		if tokenCookie, err := r.Cookie("access_token"); err == nil {
			if token, err := uuid.Parse(tokenCookie.Value); err == nil {
				if session, err := sessions.ByToken(token); err == nil {
					ctx = WithSession(ctx, *session)
				}
			}
		}

		handlerFunc(w, r.WithContext(ctx))
	}
}