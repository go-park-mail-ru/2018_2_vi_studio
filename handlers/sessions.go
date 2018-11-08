package handlers

import (
	"2018_2_vi_studio_server/resources"
	"2018_2_vi_studio_server/resources/sessions"
	"2018_2_vi_studio_server/resources/users"
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
		sh.sb.Sessions.Add(sessions.Session{AccessToken: &token, UserId: user.Id})
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
	json.NewEncoder(w).Encode(users.User{
		Nickname: user.Nickname,
		Email:    user.Email,
	})
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
