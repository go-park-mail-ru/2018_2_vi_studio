package handlers

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"main/middleware"
	"main/proto"
	"net/http"
	"time"
)

type SessionHandler struct {
	sb proto.AuthServices
}

func NewSessionsHandler(sb proto.AuthServices) *SessionHandler {
	return &SessionHandler{
		sb: sb,
	}
}

func (sh *SessionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		sh.Create(w, r)
	case "DELETE":
		sh.Delete(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

// -------------------------------------------------- Structures --------------------------------------------------

type CreateSessionRequest struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

func (csr CreateSessionRequest) IsValid() bool {
	return true
}

type CreateSessionResponse struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

// -------------------------------------------------- Handlers --------------------------------------------------

func (sh *SessionHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Decode
	decoder := json.NewDecoder(r.Body)
	var request CreateSessionRequest
	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Validate
	if !request.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Authorize
	user, err := sh.sb.Users.SignIn(context.Background(), &proto.UserSignIn{
		Nickname: request.Nickname,
		Password: request.Password,
	})
	if err != nil {
		body, err := NewError("Incorrect username or password").MarshalJSON()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusForbidden)
		_, err = w.Write(body)
		if err != nil {
			middleware.Logger(r.Context()).Error(err.Error())
		}
		return
	}

	// Get or create token
	var accessToken string
	session, err := sh.sb.Sessions.ByUserUID(context.Background(), user.Uid)

	if err == nil {
		accessToken = session.AccessToken.Value
	} else {
		tokenUUID, err := uuid.NewRandom()
		if err != nil {
			middleware.Logger(r.Context()).Error(err.Error())
		}

		accessToken = tokenUUID.String()
		_, err = sh.sb.Sessions.Add(context.Background(), &proto.Session{
			AccessToken: &proto.UUID{Value: accessToken},
			UserUID:     user.Uid,
		})
		if err != nil {
			middleware.Logger(r.Context()).Error(err.Error())
		}
	}

	// Set cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	})

	// Write response
	body, err := CreateSessionResponse{
		Nickname: user.Nickname,
		Email:    user.Email,
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

func (sh *SessionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// Remove cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		Expires:  time.Unix(0, 0),
	})
	w.WriteHeader(http.StatusOK)
}
