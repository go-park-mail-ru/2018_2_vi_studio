package handlers

import (
	"game/middleware"
	"game/resources"
	"game/resources/users"
	"encoding/json"
	"math"
	"net/http"
	"strconv"
)

type LeadersHandler struct {
	sb *resources.StorageBundle
}

func NewLeadersHandler(sb *resources.StorageBundle) *LeadersHandler {
	return &LeadersHandler{
		sb: sb,
	}
}

type LeadersResponse struct {
	leaders *[]users.User
	pages   int
}

func (lh *LeadersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	const limit = 10

	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	page := 1
	if pageString, ok := r.URL.Query()["page"]; ok {
		var err error
		page, err = strconv.Atoi(pageString[0])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	leaders, err := lh.sb.Users.Leaders(limit, limit*(page-1))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	userCount, err := lh.sb.Users.Count()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(LeadersResponse{
		leaders: leaders,
		pages:   int(math.Ceil(float64(userCount) / limit)),
	})
	if err != nil {
		middleware.Logger(r.Context()).Error(err.Error())
	}
}
