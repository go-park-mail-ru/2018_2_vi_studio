package handlers

import (
	"context"
	"encoding/json"
	"main/middleware"
	"main/proto"
	"math"
	"net/http"
	"strconv"
)

type LeadersHandler struct {
	sb proto.ServiceBundle
}

func NewLeadersHandler(sb proto.ServiceBundle) *LeadersHandler {
	return &LeadersHandler{
		sb: sb,
	}
}

func (lh *LeadersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		lh.Get(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

// -------------------------------------------------- Structures --------------------------------------------------

type GetLeadersResponse struct {
	Leaders []*proto.User `json:"leaders"`
	PageCount int `json:"pageCount"`
}

// -------------------------------------------------- Handlers --------------------------------------------------

func (lh *LeadersHandler) Get(w http.ResponseWriter, r *http.Request)  {
	const limit = 10

	page := 1
	if pageString, ok := r.URL.Query()["page"]; ok {
		var err error
		page, err = strconv.Atoi(pageString[0])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	response, err := lh.sb.Users.Leaders(
		context.Background(),
		&proto.LeadersRequest{Limit: limit, Offset: int32(limit * (page - 1))},
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(GetLeadersResponse{
		Leaders: response.Users,
		PageCount: int(math.Ceil(float64(response.Count) / limit)),
	})

	if err != nil {
		middleware.Logger(r.Context()).Error(err.Error())
	}
}