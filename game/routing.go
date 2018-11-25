package main

import (
	"database/sql"
	"game/core"
	"game/handlers"
	mw "game/middleware"
	"game/resources"
	"net/http"
)

func SetRoutes() {


	commonMW := func(handlerFunc http.HandlerFunc) http.HandlerFunc {
		return mw.LoggerMW(mw.PanicMW(mw.CorsMW(mw.OptionMW(mw.SessionMW(handlerFunc, sb.Sessions)))))
	}

	leadersHandler := handlers.NewLeadersHandler(sb)
	http.HandleFunc("/leader", commonMW(leadersHandler.ServeHTTP))

	sessionsHandler := handlers.NewSessionsHandler(sb)
	http.HandleFunc("/session", commonMW(sessionsHandler.ServeHTTP))

	usersHandler := handlers.NewUsersHandler(sb)
	http.HandleFunc("/user", commonMW(usersHandler.ServeHTTP))

	http.Handle("/media/", http.FileServer(http.Dir(".")))
	userAvatarHandler := handlers.NewUserAvatarHandler(sb)
	http.HandleFunc("/user/avatar", commonMW(userAvatarHandler.ServeHTTP))

	gm := game.NewGame()

	gameHandler := handlers.NewGameHandler(gm)
	http.HandleFunc("/game-ws", gameHandler.ServeHTTP)

}
