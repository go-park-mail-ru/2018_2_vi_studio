package handlers

import (
	"2018_2_vi_studio_server/resources"
	"encoding/json"
	"net/http"
)

type LeadersHandler struct {
	sb *resources.StorageBundle
}

func NewLeadersHandler(sb *resources.StorageBundle) *LeadersHandler{
	return &LeadersHandler{
		sb: sb,
	}
}

func (lh *LeadersHandler)ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	leaders, err := lh.sb.Users.Leaders(10, 0)

	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(leaders)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
}

