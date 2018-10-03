package handlers

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2018_2_vi_studio/resources/leaders"
	"net/http"
)

type LeadersHandler struct {
	LeaderSlice[] leaders.Leader
}

func (lh *LeadersHandler)ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(lh.LeaderSlice)

	w.WriteHeader(http.StatusOK)
}

