package game

import (
	"go.uber.org/zap"
	"time"
)

const (
	RoomSize = 2

	RoundTimeInterval = 15 * time.Second

	ClientTimeOut = 15 * time.Minute

	MaxMessageSize = 1024
)

type Game struct {
	queue  *ClientQueue
	logger *zap.Logger
}

func NewGame(logger *zap.Logger) *Game {
	return &Game{
		queue:  NewClientQueue(),
		logger: logger,
	}
}

func (game *Game) waitClientGameStart(client *Client) {
	select {
	case message := <-client.Output():
		if message.Event != EventReadyToPlay {
			game.logger.Error("Wrong client event")
			return
		}

		if game.queue.len() > 0 {
			clients := make([]*Client, 0, RoomSize)
			clients = append(clients, client)

			for i := 1; i < RoomSize; i++ {
				clients = append(clients, game.queue.pop())
			}

			NewGameRoom(clients, game.logger)
		} else {
			game.queue.push(client)
		}

		for i, client := range game.queue.Clients {
			client.Input(NewQueuePositionMessage(i))
		}
	case <-time.After(ClientTimeOut):
		return
	}
}

func (game *Game) AddClient(client *Client) {
	go game.waitClientGameStart(client)
}
