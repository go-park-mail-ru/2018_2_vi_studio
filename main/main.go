package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"main/handlers"
	mw "main/middleware"
	"main/proto"
	"net/http"
	"os"
)

func main() {
	grcpConn, err := grpc.Dial(
		"127.0.0.1:8080",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := grcpConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sessionService := proto.NewSessionServiceClient(grcpConn)
	userService := proto.NewUserServiceClient(grcpConn)
	authServices := proto.AuthServices{
		Sessions: sessionService,
		Users: userService,
	}

	commonMW := func(handlerFunc http.HandlerFunc) http.HandlerFunc {
		return mw.LoggerMW(mw.PanicMW(mw.CorsMW(mw.OptionMW(mw.SessionMW(handlerFunc, sessionService)))))
	}

	// routing (begin)
	leadersHandler := handlers.NewLeadersHandler(authServices)
	http.HandleFunc("/leader", commonMW(leadersHandler.ServeHTTP))

	sessionsHandler := handlers.NewSessionsHandler(authServices)
	http.HandleFunc("/session", commonMW(sessionsHandler.ServeHTTP))

	usersHandler := handlers.NewUsersHandler(authServices)
	http.HandleFunc("/user", commonMW(usersHandler.ServeHTTP))

	userAvatarHandler := handlers.NewUserAvatarHandler(authServices)
	http.HandleFunc("/user/avatar", commonMW(userAvatarHandler.ServeHTTP))

	http.Handle("/media/", http.FileServer(http.Dir(".")))
	// routing (end)

	port := "8000"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	fmt.Printf("starting server at :%s\n", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
