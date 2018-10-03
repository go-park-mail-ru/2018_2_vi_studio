package handlers

import "net/http"

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
