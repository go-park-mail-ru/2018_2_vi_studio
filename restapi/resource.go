package main

import (
	"encoding/json"
	"net/http"
)

type ResourceHendler struct {
	*AuthHandler
	leaderSlice  []Leader
}

type Leader struct {
	Nickname string `json:"nickname"`
	Points int `json:"points"`
}


func NewResourceHendler(ah *AuthHandler) ResourceHendler {
	return ResourceHendler{
		AuthHandler: ah,
		leaderSlice:  []Leader{
			{
				Nickname: "Viewsharp",
				Points: 7352,
			},
		},
	}
}

func (rh *ResourceHendler) leaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(rh.leaderSlice)

	w.WriteHeader(http.StatusOK)
}

func (rh *ResourceHendler) profile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	accessToken := r.Header.Get("Authorization")

	for _, token := range rh.tokens.GetAll() {
		if token.UUID.String() == accessToken {
			json.NewEncoder(w).Encode(rh.users.Get(token.UserId))
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
}