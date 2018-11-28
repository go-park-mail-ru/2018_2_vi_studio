package main

import (
	"context"
	"fmt"
	"game/core"
	"game/proto"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
)

type GameHandler struct {
	upgrader     *websocket.Upgrader
	logger       *zap.Logger
	authServices *AuthServices
	game         *game.Game
}

func NewGameHandler(authServices *AuthServices) *GameHandler {
	// websocket upgrader
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	// logger
	logger, _ := zap.NewProduction()
	//defer logger.Sync()

	return &GameHandler{
		upgrader:     &upgrader,
		logger:       logger,
		authServices: authServices,
		game:         game.NewGame(logger),
	}
}

func (gh *GameHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	accessToken, err := r.Cookie("access_token")
	if err != nil {
		gh.logger.Error(err.Error())
		return
	}

	session, err := gh.authServices.Sessions.ByToken(context.Background(), &proto.UUID{Value: accessToken.Value})
	if err != nil {
		gh.logger.Error(err.Error())
		return
	}

	user, err := gh.authServices.Users.ByUID(context.Background(), session.UserUID)
	if err != nil {
		gh.logger.Error(err.Error())
		return
	}

	gh.logger.Info("CONNECT")

	ws, err := gh.upgrader.Upgrade(w, r, nil)
	if err != nil {
		gh.logger.Fatal(fmt.Sprint(err))
		return
	}

	client := game.NewClient(user, ws, gh.logger)
	gh.game.AddClient(client)
}
