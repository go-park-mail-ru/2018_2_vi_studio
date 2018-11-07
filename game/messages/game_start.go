package messages

type GameStart struct {
}

func NewGameStartMSG() *Message {
	return &Message{
		Event: "GameStart",
	}
}
