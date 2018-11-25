package handlers

import (
	"context"
	"encoding/json"
	"main/middleware"

	//"main/middleware"
	"main/proto"
	"net/http"
)

type UsersHandler struct {
	sb proto.ServiceBundle
}

func NewUsersHandler(sb proto.ServiceBundle) *UsersHandler {
	return &UsersHandler{
		sb: sb,
	}
}

func (uh *UsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		uh.Create(w, r)
	case "GET":
		uh.Get(w, r)
	case "PUT":
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

// -------------------------------------------------- Structures --------------------------------------------------

type CreateUserRequest struct {
	Nickname string `json:"nickname"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func (cur CreateUserRequest) IsValid() bool {
	return true
}

type CreateUserResponse struct {

}

type GetUserResponse struct {
	Nickname string `json:"nickname"`
	Email string `json:"email"`
	Points int `json:"points"`
}

// -------------------------------------------------- Handlers --------------------------------------------------

func (uh *UsersHandler) Create(w http.ResponseWriter, r *http.Request) {
	//Decode
	decoder := json.NewDecoder(r.Body)
	var request CreateUserRequest
	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Validate
	if !request.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Add user to storage
	_, err = uh.sb.Users.SignUp(context.Background(), &proto.UserSignUp{
		Nickname: request.Nickname,
		Email: request.Email,
		Password: request.Password,
	})
	if err == nil {
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(CreateUserResponse{
			// TODO: write some
		})
		if err != nil {
			middleware.Logger(r.Context()).Error(err.Error())
		}
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
}

func (uh *UsersHandler) Get(w http.ResponseWriter, r *http.Request) {
	session, ok := middleware.Session(r.Context())

	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(NewError("user is not authorized"))
		if err != nil {
			middleware.Logger(r.Context()).Error(err.Error())
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	user, _ := uh.sb.Users.ByUID(context.Background(), session.UserUID)
	err := json.NewEncoder(w).Encode(GetUserResponse{
		Nickname: user.Nickname,
		Email: user.Email,
		Points: int(user.Points),
	})
	if err != nil {
		middleware.Logger(r.Context()).Error(err.Error())
	}
}
