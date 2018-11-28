package game

type Message struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

const (
	EventReadyToPlay   = "READY_TO_PLAY"
	EventQueuePosition = "QUEUE_POSITION"
	EventGameStart     = "GAME_START"
	//EventGameOver      = "GAME_OVER"
	EventNextTry  = "NEXT_TRY"
	EventDoneTry  = "DONE_TRY"
	EventWrongTry = "WRONG_TRY"
)

// ------------------------------------------------- Queue Position -------------------------------------------------

type msgQueuePosition struct {
	Position int `json:"position"`
}

func NewQueuePositionMessage(position int) Message {
	return Message{
		Event: EventQueuePosition,
		Data: msgQueuePosition{
			Position: position,
		},
	}
}

// ------------------------------------------------- Game Start -------------------------------------------------

type msgPlayer struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type msgGameStart struct {
	UserId  int         `json:"userId"`
	Players []msgPlayer `json:"players"`
}

func NewGameStartMessage(userId int, players []msgPlayer) Message {
	return Message{
		Event: EventGameStart,
		Data: msgGameStart{
			UserId:  userId,
			Players: players,
		},
	}
}

// ------------------------------------------------- Game Over -------------------------------------------------

//type msgGameOver struct {
//}
//
//func NewGameOverMessage() Message {
//	return Message{
//		Event: EventGameOver,
//		Data:  msgGameOver{},
//	}
//}

// ------------------------------------------------- Next try -------------------------------------------------

type msgCurretTry struct {
	PlayerId int `json:"playerId"`
	TileType int `json:"tileType"`
}

type msgGameOver struct {
}

type msgNextTry struct {
	LastTry    msgDoneTry   `json:"lastTry"`
	CurrentTry msgCurretTry `json:"currentTry"`
	GameOver   msgGameOver  `json:"gameOver"`
}

func NewNextTryMessage(lastTry msgDoneTry, currentTry msgCurretTry, gameOver msgGameOver) Message {
	return Message{
		Event: EventNextTry,
		Data: msgNextTry{
			LastTry:    lastTry,
			CurrentTry: currentTry,
			GameOver:   gameOver,
		},
	}
}

// ------------------------------------------------- Wrong try -------------------------------------------------

type msgWrongTry struct {

}

func NewWrongTryMessage() Message {
	return Message{
		Event: EventWrongTry,
		Data:  msgWrongTry{},
	}
}

// ------------------------------------------------- Done try -------------------------------------------------

type msgDoneTry struct {
	Row int `json:"row"`
	Col int `json:"col"`
	Rotation int `json:"rotationCount"`
}

func NewDoneTryMessage(doneTry msgDoneTry) Message {
	return Message{
		Event: EventDoneTry,
		Data:  msgDoneTry{},
	}
}
