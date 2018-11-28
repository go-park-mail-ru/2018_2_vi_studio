package game

import (
	"game/proto"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"time"
)

type Client struct {
	user    *proto.User
	ws      *websocket.Conn
	runLoop bool
	logger  *zap.Logger
	input   chan Message
	output  chan Message
}

func NewClient(user *proto.User, ws *websocket.Conn, logger *zap.Logger) *Client {
	result := &Client{
		user:   user,
		ws:     ws,
		logger: logger,
		input:  make(chan Message),
		output: make(chan Message),
	}
	go result.Init()
	return result
}

func (client *Client) Input(message Message) {
	client.input <- message
}

func (client *Client) Output() <-chan Message {
	return client.output
}

func (client *Client) readPump(wsReadChan chan<- Message) {
	defer func() {
		_ = client.ws.Close()
		close(wsReadChan)
	}()

	// init ws connect
	client.ws.SetReadLimit(MaxMessageSize)
	err := client.ws.SetReadDeadline(time.Now().Add(ClientTimeOut))
	client.ws.SetPongHandler(func(string) error { return client.ws.SetReadDeadline(time.Now().Add(ClientTimeOut)) })
	if err != nil {
		client.logger.Error(err.Error())
	}

	// read from web socket to channel
	for {
		var message Message

		_, bytes, err := client.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				client.logger.Error(err.Error())
			}
			break
		}

		err = message.UnmarshalJSON(bytes)
		if err != nil {
			client.logger.Error(err.Error())
			break
		}

		wsReadChan <- message
	}
}

func (client *Client) Init() {
	wsReadChan := make(chan Message)

	go client.readPump(wsReadChan)

	// web socket read loop
	for {
		select {
		case message := <-wsReadChan:
			client.output <- message
		case message := <-client.input:
			bytes, err := message.MarshalJSON()
			if err != nil {
				client.logger.Error(err.Error())
			}

			err = client.ws.WriteMessage(0, bytes)
			if err != nil {
				client.logger.Error(err.Error())
			}
		case <-time.After(ClientTimeOut):
			err := client.ws.Close()
			if err != nil {
				client.logger.Error(err.Error())
			}

			break
		}
	}
}
