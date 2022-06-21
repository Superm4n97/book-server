package middlewares

import (
	"github.com/Superm4n97/Book-Server/auth"
	"net/http"
	"strings"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//uname, pass, ok := r.BasicAuth()
		authHeader := r.Header.Get("Authorization")
		// authHeader: "Basic ak83djkfdfj84dj"

		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("user authorization missing"))
			return
		}

		authorizationInfo := strings.Split(authHeader, " ")

		if len(authorizationInfo) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("invalid user authorization info"))
			return
		}

		var err error

		if authorizationInfo[0] == "Basic" {
			err = auth.BasicAuth(authorizationInfo[1])
		} else if authorizationInfo[0] == "Bearer" {
			err = auth.BearerAuth(authorizationInfo[1])
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid user authorization info"))
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}
		next.ServeHTTP(w, r)

		return
	})
}
