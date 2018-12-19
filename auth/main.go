package main

import (
	"auth/resources/proto"
	"auth/resources/sessions"
	"auth/resources/users"
	"database/sql"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"

	_ "github.com/lib/pq"
)

const DSN = "host=127.0.0.1 port=5432 user=docker password=docker dbname=docker sslmode=disable"

func main() {
	db, err := sql.Open("postgres", DSN)
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	sessionService := sessions.NewSessionService(sessions.NewSessionStorage(db))
	userService := users.NewUserService(users.NewUserStorage(db))

	server := grpc.NewServer()

	proto.RegisterSessionServiceServer(server, sessionService)
	proto.RegisterUserServiceServer(server, userService)

	fmt.Println("starting server at :8080")

	err = server.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
