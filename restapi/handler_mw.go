package main

import (
	"log"
	"net/http"
)

type ResponseWriter struct {
	http.ResponseWriter
	status int
}

func (w *ResponseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func handlerMW(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customRW := ResponseWriter{w, 200}
		handlerFunc(&customRW, r)
		log.Printf("%d [%s] %s\n", customRW.status, r.Method, r.URL.Path)
	}
}