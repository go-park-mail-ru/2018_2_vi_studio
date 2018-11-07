package messages

type QueuePosition struct {
}

func NewQueuePositionMSG(position int) *Message {
	return &Message{
		Event: "QueuePosition",
		Data:  position,
	}
}
