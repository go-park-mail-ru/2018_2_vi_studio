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
)

func main() {
	grcpConn, err := grpc.Dial(
		"127.0.0.1:9000",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer grcpConn.Close()

	sessionService := proto.NewSessionServiceClient(grcpConn)
	userService := proto.NewUserServiceClient(grcpConn)
	serviceBundle := proto.ServiceBundle{
		Sessions: sessionService,
		Users: userService,
	}

	commonMW := func(handlerFunc http.HandlerFunc) http.HandlerFunc {
		return mw.LoggerMW(mw.PanicMW(mw.CorsMW(mw.OptionMW(mw.SessionMW(handlerFunc, sessionService)))))
	}

	// routing (begin)
	leadersHandler := handlers.NewLeadersHandler(serviceBundle)
	http.HandleFunc("/leader", commonMW(leadersHandler.ServeHTTP))

	sessionsHandler := handlers.NewSessionsHandler(serviceBundle)
	http.HandleFunc("/session", commonMW(sessionsHandler.ServeHTTP))

	usersHandler := handlers.NewUsersHandler(serviceBundle)
	http.HandleFunc("/user", commonMW(usersHandler.ServeHTTP))

	userAvatarHandler := handlers.NewUserAvatarHandler(serviceBundle)
	http.HandleFunc("/user/avatar", commonMW(userAvatarHandler.ServeHTTP))

	http.Handle("/media/", http.FileServer(http.Dir(".")))
	// routing (end)

	fmt.Println("starting server at :8000")
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
