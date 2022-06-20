package middlewares

import (
	"encoding/base64"
	"fmt"
	"github.com/Superm4n97/Book-Server/model"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

func basicAuth(req http.Handler, w http.ResponseWriter, r *http.Request, enStr string) {
	decodedInfo, err := base64.StdEncoding.DecodeString(enStr)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	usernamePassword := strings.Split(string(decodedInfo), ":")

	//fmt.Println("username :", usernamePassword[0])
	//fmt.Println("password :", usernamePassword[1])

	if model.UserInfo[usernamePassword[0]] != usernamePassword[1] {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	req.ServeHTTP(w, r)
}

func bearerAuth(req http.Handler, w http.ResponseWriter, r *http.Request, tokenStr string) {
	fmt.Println(tokenStr)
	token, err := jwt.Parse(tokenStr, func(tkn *jwt.Token) (interface{}, error) {
		if _, ok := tkn.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", tkn.Header["alg"])
		}

		return []byte(model.ServerSecretKey), nil
	})

	clms, ok := token.Claims.(jwt.MapClaims)

	fmt.Println(clms)
	fmt.Println(ok)
	fmt.Println(token.Valid)

	if ok && token.Valid {
		//fmt.Println("Token Matched !!!!")
		fmt.Println(clms["userName"])
		//w.Write([]byte("Token Matched, valid user"))

		req.ServeHTTP(w, r)
	} else {
		fmt.Println("last error ", err)
	}
}

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//uname, pass, ok := r.BasicAuth()
		authHeader := r.Header.Get("Authorization")
		// authHeader: "Basic ak83djkfdfj84dj"

		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		authorizationInfo := strings.Split(authHeader, " ")

		if len(authorizationInfo) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if authorizationInfo[0] == "Basic" {
			basicAuth(next, w, r, authorizationInfo[1])
		} else if authorizationInfo[0] == "Bearer" {
			bearerAuth(next, w, r, authorizationInfo[1])
		}

		return
	})
}
