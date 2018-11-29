package game

import (
	"go.uber.org/zap"
	"time"
)

type Room struct {
	players    []*Client
	input      chan Message
	runLoop    bool
	moveNumber int
	tiles      []Tile
	lastTry    msgDoneTry
	logger     *zap.Logger
}

func NewGameRoom(players []*Client, logger *zap.Logger) *Room {
	result := &Room{
		players:    players,
		input:      make(chan Message),
		runLoop:    true,
		tiles:      NewTileSlice(),
		moveNumber: 0,
		lastTry:    msgDoneTry{},
		logger:     logger,
	}
	go result.Init()
	return result
}

func (room *Room) Init() {
	players := make([]msgPlayer, 0, RoomSize)

	for i, player := range room.players {
		players = append(players, msgPlayer{
			Id:       i,
			Nickname: player.user.Nickname,
			Avatar:   player.user.Avatar,
		})
	}

	// game start
	for i, player := range room.players {
		player.Input(NewGameStartMessage(i, players))
	}

	// game loop
	for room.runLoop {
		for i, player := range room.players {
			room.move(i, player)

			if room.moveNumber == len(room.tiles) {
				room.runLoop = false
			}
		}
	}

	// game over
	room.playersOutput(NewNextTryMessage(room.lastTry, msgCurretTry{}, msgGameOver{}))
}

func (room *Room) move(playerId int, player *Client) {
	currentTry := msgCurretTry{
		PlayerId: playerId,
		TileType: room.tiles[room.moveNumber][0],
	}

	room.playersOutput(NewNextTryMessage(room.lastTry, currentTry, msgGameOver{}))
	room.lastTry = msgDoneTry{}

	for {
		select {
		case message := <-player.Output():
			if message.Event == EventDoneTry {
				data := message.Data.(map[string]interface{})
				room.lastTry = msgDoneTry{
					Row:      int(data["row"].(float64)),
					Col:      int(data["col"].(float64)),
					Rotation: int(data["rotationCount"].(float64)),
				}

				room.moveNumber++
				return
			} else {
				room.logger.Error("Wrong client event")
				return
			}

		case <-time.After(RoundTimeInterval):
			return
		}
	}
}

func (room *Room) playersOutput(message Message) {
	for _, player := range room.players {
		player.Input(message)
	}
}
