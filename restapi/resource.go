package main

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2018_2_vi_studio/restapi/models"
	"github.com/google/uuid"
	"net/http"
)

type ResourceHendler struct {
	*AuthHandler
	leaderSlice []Leader
}

type Leader struct {
	Nickname string `json:"nickname"`
	Points   int    `json:"points"`
}

func NewResourceHendler(ah *AuthHandler) ResourceHendler {
	return ResourceHendler{
		AuthHandler: ah,
		leaderSlice: []Leader{
			{
				Nickname: "Viewsharp",
				Points:   7352,
			},
		},
	}
}

func (rh *ResourceHendler) leaders(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(rh.leaderSlice)

	w.WriteHeader(http.StatusOK)
}

func (rh *ResourceHendler) profile(w http.ResponseWriter, r *http.Request) {
	GET := func(w http.ResponseWriter, r *http.Request) {
		tokenCookie, _ := r.Cookie("access_token")
		token, err := uuid.Parse(tokenCookie.Value)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if id, ok := rh.tokens.GetUserId(token); ok {
			w.WriteHeader(http.StatusOK)

			user := rh.users.Get(id)
			json.NewEncoder(w).Encode(models.User{
				Nickname:user.Nickname,
				Email:user.Email,
			})
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}

	switch r.Method {
	case "GET":
		GET(w,r)
	case "PUT":
		// TODO: make profile put
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
