package game

import (
	"2018_2_vi_studio_server/game/messages"
	"github.com/gorilla/websocket"
)

type Client struct {
	ws *websocket.Conn
	loop bool
	game *Game

}

func NewClient(ws *websocket.Conn, game *Game) *Client {
	result := &Client{
		ws: ws,
		game: game,
		loop: true,
	}
	go result.Loop()
	return result
}

// TODO: rewrite

func (client *Client) Loop() {
	for client.loop {
		var message messages.Message
		err := client.ws.ReadJSON(&message)
		if err != nil {
			return // TODO: Handle error
		}

		switch message.Event {
		case "ReadyToPlay":
			client.readyToPlay()
		case "DoneTry":
			client.doneTry()
		default:
		}
	}
}


func (client *Client) QueuePosition(position int) {
	client.ws.WriteJSON(messages.NewQueuePositionMSG(position))
}

func (client *Client) GameStart() {
	client.ws.WriteJSON(messages.NewGameStartMSG())
}

func (client *Client) StartTry() {
	client.ws.WriteJSON(messages.NewQueuePositionMSG(0))
}

func (client *Client) TimeOut() {
	client.ws.WriteJSON(messages.NewQueuePositionMSG(0))
}

func (client *Client) GameOver() {
	client.ws.WriteJSON(messages.NewGameOverMSG())
}

func (client *Client) readyToPlay() {
	client.game.AddClient(client)
}

func (client *Client) doneTry() {
	client.game.AddClient(client)
}

/*
------------------------------------------ RESPONSES ------------------------------------------

// game-event-QueuePosition
event: QueuePosition
Clients: {
    position int
}

// game-event-GameStart
event: GameStart
Clients: {
    players: [
        {   id int // local (game) id
            username string
            avatar string // image url
        }
    ],
    clientId int
}

// game-event-StartTry
event: StartTry
Clients: {
    playerId int
    tile: {
        type int
    }
}

// game-event-TimeOut
event: TimeOut
Clients: {
}

// game-event-GameOver
event: GameOver
Clients: {
    players {
        id int
        points int
    }
}

------------------------------------------ REQUESTS ------------------------------------------

// game-event-ReadyToPlay
event: ReadyToPlay
Clients: {
}

// game-event-DoneTry
event: DoneTry
Clients: {
    row int
    col int
    rotation int
}
 */
