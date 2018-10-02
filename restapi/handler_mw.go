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

func logMW(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customRW := ResponseWriter{w, 200}
		handlerFunc(&customRW, r)
		log.Printf("%d [%s] %s\n", customRW.status, r.Method, r.URL.Path)
	}
}

func corsMW(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Origin validation on regular expressions
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		handlerFunc(w, r)
	}
}

func optionMW(handlerFunc http.HandlerFunc) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		} else {
			handlerFunc(w, r)
		}
	}
}

func AuthMW(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("access_token")
		if err == nil {
			handlerFunc(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}

func CommonMW(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return logMW(corsMW(optionMW(handlerFunc)))
}
