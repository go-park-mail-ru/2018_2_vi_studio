package main

import (
	"fmt"
	"github.com/go-park-mail-ru/2018_2_vi_studio/handlers"
	"github.com/go-park-mail-ru/2018_2_vi_studio/resources/leaders"
	"github.com/go-park-mail-ru/2018_2_vi_studio/resources/sessions"
	"github.com/go-park-mail-ru/2018_2_vi_studio/resources/users"
	"net/http"
	"os"

	//_ "github.com/lib/pq"
)

//var DSN = "host=127.0.0.1 port=5432 user=docker password=docker dbname=docker"

func main() {
	//db, err := sql.Open("postgres", DSN)
	//err = db.Ping() // вот тут будет первое подключение к базе
	//if err != nil {
	//	panic(err)
	//}

	port := "8080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	userStorage := users.NewUserIMS()
	sessionStorage := sessions.NewSessionIMS()

	userHandler := handlers.UsersHandler{
		UserStorage: userStorage,
		SessionStorage: sessionStorage,
	}

	http.HandleFunc("/user", CommonMW(userHandler.ServeHTTP))


	sessionHandler := handlers.SessionHandler{
		UserStorage: userStorage,
		SessionStorage: sessionStorage,
	}

	http.HandleFunc("/session", CommonMW(sessionHandler.ServeHTTP))

	leadersHandler := handlers.LeadersHandler{
		LeaderSlice: []leaders.Leader{
			{
				Nickname: "viewsharp",
				Points: 2018,
			},
		},
	}

	http.HandleFunc("/leader", CommonMW(leadersHandler.ServeHTTP))


	fmt.Printf("starting server at :%s\n", port)
	http.ListenAndServe(":" + port, nil)
}