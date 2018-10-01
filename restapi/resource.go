package main

import (
	"encoding/json"
	"net/http"
)

type ResourceHendler struct {
	leaderSlice  []Leader
}

type Leader struct {
	Nickname string `json:"nickname"`
	Points int `json:"points"`
}


func NewResourceHendler() ResourceHendler {
	return ResourceHendler{
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