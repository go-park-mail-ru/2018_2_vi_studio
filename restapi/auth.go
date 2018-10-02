package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"time"

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

func NewAuthHandler(usersDAO *dao.UserDAO) AuthHandler {
	return AuthHandler{
		users:  usersDAO,
		tokens: dao.NewTokenDAO(),
	}
}

func (ah *AuthHandler) signIn(w http.ResponseWriter, r *http.Request) {
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

	for _, user := range ah.users.GetAll() {
		if user.Nickname == sIUser.Nickname && user.Password == sIUser.Password {
			var accessToken uuid.UUID

			if token, ok := ah.tokens.GetToken(user.Id); ok {
				accessToken = token
			} else {
				accessToken = uuid.New()
				ah.tokens.Set(accessToken, user.Id)
			}

			http.SetCookie(w, &http.Cookie{
				Name: "access_token",
				Value: accessToken.String(),
				HttpOnly: true,
				Secure:false,
				Path: "/",
			})

			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(models.User{
				Nickname:user.Nickname,
				Email:user.Email,
			})
			return
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
}

func (ah *AuthHandler) signUp(w http.ResponseWriter, r *http.Request) {
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

func (ah *AuthHandler) signOut(w http.ResponseWriter, r *http.Request)  {
	http.SetCookie(w, &http.Cookie{
		Name: "access_token",
		Value: "",
		HttpOnly: true,
		Secure:false,
		Path: "/",
		Expires: time.Unix(0, 0),
	})
	w.WriteHeader(http.StatusOK)
}
