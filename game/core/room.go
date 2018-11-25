package game

type Room struct {
	players []*Client
}

func NewGameRoom(players []*Client) *Room {
	result := &Room{players: players}
	go result.Loop()
	return result
}

func (room *Room)Loop()  {
	for _, player := range room.players {
		player.GameStart()
		player.GameOver()
	}
}
