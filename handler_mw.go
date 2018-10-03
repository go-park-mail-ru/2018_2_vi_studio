package main

import (
	"log"
	"net/http"
	"regexp"
)

var regexOrigin, _ = regexp.Compile(`^https?\:\/\/((2018\-2\-vi\-studio\-\w+\.now\.sh)|((127\.0\.0\.1|localhost)(\:\d+)?))$`)

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
		log.Printf("%d\t[%s]\t%s\tOrigin: %s\n", customRW.status, r.Method, r.URL.Path, r.Header.Get("Origin"))
	}
}

func corsMW(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if regexOrigin.MatchString(r.Header.Get("Origin")) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}
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

func CommonMW(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return logMW(corsMW(optionMW(handlerFunc)))
}
