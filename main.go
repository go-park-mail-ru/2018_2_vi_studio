package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"os"
)

func main() {
	port := "8080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	SetRoutes()

	fmt.Printf("starting server at :%s\n", port)
	http.ListenAndServe(":"+port, nil)
}
