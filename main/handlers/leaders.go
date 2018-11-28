package handlers

import (
	"context"
	"main/middleware"
	"main/proto"
	"math"
	"net/http"
	"strconv"
)

type LeadersHandler struct {
	sb proto.AuthServices
}

func NewLeadersHandler(sb proto.AuthServices) *LeadersHandler {
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

	body, err := GetLeadersResponse{
		Leaders: response.Users,
		PageCount: int(math.Ceil(float64(response.Count) / limit)),
	}.MarshalJSON()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(body)
	if err != nil {
		middleware.Logger(r.Context()).Error(err.Error())
	}
}