package messages

type Message struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}
