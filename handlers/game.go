package handlers

import (
	"2018_2_vi_studio_server/game"
	"2018_2_vi_studio_server/middleware"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

type GameHandler struct {
	upgrader websocket.Upgrader
	game     *game.Game
}

func NewGameHandler(game *game.Game) *GameHandler {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	return &GameHandler{
		upgrader: upgrader,
		game:     game,
	}
}

func (gh *GameHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ws, err := gh.upgrader.Upgrade(w, r, nil)
	if err != nil {
		middleware.Logger(r.Context()).Fatal(fmt.Sprint(err))
		return
	}
	game.NewClient(ws, gh.game) // !!!!!
}
