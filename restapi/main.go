package main

import (
	"fmt"
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

	ah := NewAuthHandler()
	http.HandleFunc("/auth/sign-in", handlerMW(ah.signIn))
	http.HandleFunc("/auth/sign-up", handlerMW(ah.signUp))

	rh := NewResourceHendler()
	http.HandleFunc("/resource/leaders", handlerMW(rh.leaders))


	fmt.Printf("starting server at :%s\n", port)
	http.ListenAndServe(":" + port, nil)
}