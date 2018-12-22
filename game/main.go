package main

import (
	"fmt"
	"game/proto"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
)

type AuthServices struct {
	Sessions proto.SessionServiceClient
	Users    proto.UserServiceClient
}

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

	authServices := &AuthServices{
		Sessions: proto.NewSessionServiceClient(grcpConn),
		Users: proto.NewUserServiceClient(grcpConn),
	}

	http.Handle("/", NewGameHandler(authServices))

	port := "8001"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	fmt.Printf("starting server at :%s\n", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
