package game

import "fmt"

type Game struct {
	queue         ClientQueue
	newClientChan chan *Client
}

func NewGame() *Game {
	result := &Game{
		newClientChan: make(chan *Client),
	}
	go result.Loop()
	return result
}

func (game *Game)Loop() {
	for {
		client := <- game.newClientChan
		if game.queue.len() > 0 {
			NewGameRoom([]*Client{client, game.queue.pop()}) // !!!!!!!!
		} else {
			game.queue.push(client)
		}

		fmt.Println(game.queue.len())

		for i, client := range game.queue.Clients {
			client.QueuePosition(i)
		}
	}
}

func (game *Game)AddClient(client *Client)  {
	game.newClientChan <- client
}
