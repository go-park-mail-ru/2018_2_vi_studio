package handlers

import (
	"2018_2_vi_studio_server/middleware"
	"2018_2_vi_studio_server/resources"
	"2018_2_vi_studio_server/resources/users"
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
		json.NewEncoder(w).Encode(user)
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
}

func (uh *UsersHandler) Get(w http.ResponseWriter, r *http.Request) {
	session, ok := middleware.Session(r.Context())

	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
	}

	w.WriteHeader(http.StatusOK)
	user, _ := uh.sb.Users.ById(*session.UserId)
	json.NewEncoder(w).Encode(users.User{
		Nickname: user.Nickname,
		Email:    user.Email,
	})
}
