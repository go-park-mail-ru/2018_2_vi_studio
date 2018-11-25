package handlers

import (
	"game/middleware"
	"game/resources"
	"game/resources/sessions"
	"game/resources/users"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type SessionHandler struct {
	sb *resources.StorageBundle
}

func NewSessionsHandler(sb *resources.StorageBundle) *SessionHandler  {
	return &SessionHandler{
		sb: sb,
	}
}

func (sh *SessionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "POST":
		sh.Create(w, r)
	case "DELETE":
		sh.Delete(w, r)
	}
	w.WriteHeader(http.StatusNotFound)
}

func (sh *SessionHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Decode
	decoder := json.NewDecoder(r.Body)
	var userAuth users.UserAuth
	err := decoder.Decode(&userAuth)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Validate
	if !userAuth.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Authorize
	user, err := sh.sb.Users.Auth(userAuth)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get or create token
	var token uuid.UUID
	session, err := sh.sb.Sessions.ByUserId(*user.Id)

	if err == nil {
		token = *session.AccessToken
	} else if err == sessions.ErrNotFound {
		token, _ = uuid.NewRandom()
		err := sh.sb.Sessions.Add(sessions.Session{AccessToken: &token, UserId: user.Id})
		if err != nil {
			middleware.Logger(r.Context()).Error(err.Error())
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Set cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    token.String(),
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
	})

	// Write response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(users.User{
		Nickname: user.Nickname,
		Email:    user.Email,
	})
	if err != nil {
		middleware.Logger(r.Context()).Error(err.Error())
	}
}

func (sh *SessionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// Remove cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
		Expires:  time.Unix(0, 0),
	})
	w.WriteHeader(http.StatusOK)
}
