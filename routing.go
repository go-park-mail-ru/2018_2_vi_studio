package main

import (
	"2018_2_vi_studio_server/game"
	"2018_2_vi_studio_server/handlers"
	mw "2018_2_vi_studio_server/middleware"
	"2018_2_vi_studio_server/resources"
	"database/sql"
	"net/http"
)

const DSN = "host=127.0.0.1 port=5432 user=docker password=docker dbname=docker"

func SetRoutes() {
	db, err := sql.Open("postgres", DSN)
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	sb := resources.NewStorageBundle(db)
	commonMW := func(handlerFunc http.HandlerFunc) http.HandlerFunc {
		return mw.LoggerMW(mw.PanicMW(mw.CorsMW(mw.OptionMW(mw.SessionMW(handlerFunc, sb.Sessions)))))
	}

	leadersHandler := handlers.NewLeadersHandler(sb)
	http.HandleFunc("/leader", commonMW(leadersHandler.ServeHTTP))

	sessionsHandler := handlers.NewSessionsHandler(sb)
	http.HandleFunc("/session", commonMW(sessionsHandler.ServeHTTP))

	usersHandler := handlers.NewUsersHandler(sb)
	http.HandleFunc("/user", commonMW(usersHandler.ServeHTTP))

	gm := game.NewGame()

	gameHandler := handlers.NewGameHandler(gm)
	http.HandleFunc("/game-ws", gameHandler.ServeHTTP)

}
