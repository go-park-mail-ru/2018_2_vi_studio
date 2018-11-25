package messages

type GameOver struct {
}

func NewGameOverMSG() *Message {
	return &Message{
		Event: "GameOver",
	}
}