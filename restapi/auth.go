package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"

	"github.com/go-park-mail-ru/2018_2_vi_studio/restapi/dao"
	"github.com/go-park-mail-ru/2018_2_vi_studio/restapi/models"
)

type AuthHandler struct {
	users  *dao.UserDAO
	tokens *dao.TokenDAO
}

type SignInUser struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type SignInToken struct {
	AccessToken uuid.UUID
}

type UserSlice []models.User

func NewAuthHandler() AuthHandler {
	return AuthHandler{
		users:  dao.NewUserDAO(),
		tokens: dao.NewTokenDAO(),
	}
}

func (ah *AuthHandler) signIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var sIUser SignInUser
	err := decoder.Decode(&sIUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for id, user := range ah.users.GetAll()  {
		if user.Nickname == sIUser.Nickname && user.Password == sIUser.Password {
			w.WriteHeader(http.StatusOK)

			if token, ok := ah.tokens.Get(id); ok {
				json.NewEncoder(w).Encode(SignInToken{AccessToken:token.UUID})
			} else {
				accessToken := uuid.New()
				ah.tokens.Save(models.Token{
					UserId: id,
					UUID:   accessToken,
				})
				json.NewEncoder(w).Encode(SignInToken{AccessToken:accessToken})
			}
			return
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
}

func (ah *AuthHandler) signUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var user models.User
	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = ah.users.Save(user)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
}
