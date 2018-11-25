package handlers

import (
	"game/middleware"
	"game/resources"
	"game/resources/users"
	"encoding/json"
	"net/http"
)

type UsersHandler struct {
	sb *resources.StorageBundle
}

func NewUsersHandler(sb *resources.StorageBundle) *UsersHandler {
	return &UsersHandler{
		sb: sb,
	}
}

func (uh *UsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		uh.Create(w, r)
	case "GET":
		uh.Get(w, r)
	case "PUT":
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (uh *UsersHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Decode
	decoder := json.NewDecoder(r.Body)
	var user users.User
	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Validate
	if !user.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Add user to storage
	err = uh.sb.Users.Add(user)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(user)
		if err != nil {
			middleware.Logger(r.Context()).Error(err.Error())
		}
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
}

func (uh *UsersHandler) Get(w http.ResponseWriter, r *http.Request) {
	session, ok := middleware.Session(r.Context())

	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(NewError("user is not authorized"))
		if err != nil {
			middleware.Logger(r.Context()).Error(err.Error())
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	user, _ := uh.sb.Users.ById(*session.UserId)
	err := json.NewEncoder(w).Encode(users.User{
		Nickname: user.Nickname,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Points:   user.Points,
	})
	if err != nil {
		middleware.Logger(r.Context()).Error(err.Error())
	}
}
