package handlers

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2018_2_vi_studio/resources/sessions"
	"github.com/go-park-mail-ru/2018_2_vi_studio/resources/users"
	"github.com/google/uuid"
	"net/http"
)

type UsersHandler struct {
	UserStorage    *users.UserIMS
	SessionStorage *sessions.SessionIMS
}

func (uh *UsersHandler)ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	switch r.Method {
	case "POST":
		uh.Create(w ,r)
	case "GET":
		AuthMW(uh.Get)(w, r)
	}
	w.WriteHeader(http.StatusNotFound)
}

func (uh *UsersHandler)Create(w http.ResponseWriter, r *http.Request)  {
	decoder := json.NewDecoder(r.Body)
	var user users.User
	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = uh.UserStorage.Add(user)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (uh *UsersHandler)Get(w http.ResponseWriter, r *http.Request)  {
	tokenCookie, _ := r.Cookie("access_token")
	token, err := uuid.Parse(tokenCookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if session, ok := uh.SessionStorage.ByToken(token); ok {
		w.WriteHeader(http.StatusOK)

		user, _ := uh.UserStorage.ById(session.UserId)
		json.NewEncoder(w).Encode(users.User{
			Nickname:user.Nickname,
			Email:user.Email,
		})
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

