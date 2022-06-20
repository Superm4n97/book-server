package middlewares

import (
	"encoding/base64"
	"github.com/Superm4n97/Book-Server/info"
	"net/http"
	"strings"
)

func BasicAuthentication(req http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//uname, pass, ok := r.BasicAuth()
		authHeader := r.Header.Get("Authorization")
		// authHeader: "Basic ak83djkfdfj84dj"

		if authHeader == "" {
			http.Error(w, http.StatusText(401), 401)
			w.Write([]byte("Unauthorized!!!!"))

			return
		}

		authorizationInfo := strings.Split(authHeader, " ")

		if len(authorizationInfo) != 2 {
			http.Error(w, http.StatusText(401), 401)
			w.Write([]byte("Invalid Authorization!!!!"))

			return
		}

		decodedInfo, err := base64.StdEncoding.DecodeString(authorizationInfo[1])

		if err != nil {

			http.Error(w, http.StatusText(401), 401)
			w.Write([]byte("Invalid Authorization!!!!"))

			return
		}

		usernamePassword := strings.Split(string(decodedInfo), ":")

		//fmt.Println("username :", usernamePassword[0])
		//fmt.Println("password :", usernamePassword[1])

		if info.UserInfo[usernamePassword[0]] != usernamePassword[1] {

			http.Error(w, http.StatusText(401), 401)
			w.Write([]byte("Invalid Authorization!!!!"))

			return
		}

		req.ServeHTTP(w, r)

		//w.Write([]byte("middleman is working!!!!!!!!!!"))
	})
}
