package middleware

import (
	"fmt"
	"net/http"
	"regexp"
)

var regexOrigin, _ = regexp.Compile(`^https?\:\/\/((2018\-2\-vi\-studio\-\w+\.now\.sh)|((127\.0\.0\.1|localhost|95\.163\.180\.8)(\:\d+)?))$`)

func CorsMW(handlerFunc http.HandlerFunc) http.HandlerFunc {
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

func OptionMW(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		} else {
			handlerFunc(w, r)
		}
	}
}

func PanicMW(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				Logger(r.Context()).Panic(fmt.Sprint("Recovered panic: ", rec))
			}
		}()

		handlerFunc(w, r)
	}
}