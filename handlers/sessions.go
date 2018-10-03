package handlers

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2018_2_vi_studio/resources/sessions"
	"github.com/go-park-mail-ru/2018_2_vi_studio/resources/users"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type SessionHandler struct {
	UserStorage    *users.UserIMS
	SessionStorage *sessions.SessionIMS
}

type signInUser struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

func (sh *SessionHandler)ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	switch r.Method {
	case "POST":
		sh.Create(w ,r)
	case "DELETE":
		sh.Delete(w, r)
	}
	w.WriteHeader(http.StatusNotFound)
}

func (sh *SessionHandler)Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var sIUser signInUser
	err := decoder.Decode(&sIUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if user, ok := sh.UserStorage.ByNickname(sIUser.Nickname); ok && user.Password == sIUser.Password {
		var token uuid.UUID

		if session, ok := sh.SessionStorage.ByUserId(user.Id); ok {
			token = session.AccessToken
		} else {
			token = uuid.New()
			sh.SessionStorage.Add(sessions.Session{AccessToken: token, UserId:user.Id})
		}

		http.SetCookie(w, &http.Cookie{
			Name: "access_token",
			Value: token.String(),
			HttpOnly: true,
			Secure:false,
			Path: "/",
		})

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(users.User{
			Nickname:user.Nickname,
			Email:user.Email,
		})
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
}

func (sh *SessionHandler)Delete(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "access_token",
		HttpOnly: true,
		Secure:false,
		Path: "/",
		Expires: time.Unix(0, 0),
	})
	w.WriteHeader(http.StatusOK)
}
