package middlewares

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/Superm4n97/Book-Server/model"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

func basicAuth(enStr string) error {
	decodedInfo, err := base64.StdEncoding.DecodeString(enStr)

	if err != nil {
		return errors.New("unable to decode the encoded string")
	}

	usernamePassword := strings.Split(string(decodedInfo), ":")

	//fmt.Println("username :", usernamePassword[0])
	//fmt.Println("password :", usernamePassword[1])

	if model.UserInfo[usernamePassword[0]] != usernamePassword[1] {
		return errors.New("wrong username or password")
	}

	return nil
	//req.ServeHTTP(w, r)
}

func bearerAuth(tokenStr string) error {
	fmt.Println(tokenStr)
	token, err := jwt.Parse(tokenStr, func(tkn *jwt.Token) (interface{}, error) {
		if _, ok := tkn.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", tkn.Header["alg"])
		}

		return []byte(model.ServerSecretKey), nil
	})

	//clms, ok := token.Claims.(jwt.MapClaims)
	_, ok := token.Claims.(jwt.MapClaims)

	//fmt.Println(clms)
	//fmt.Println(ok)
	//fmt.Println(token.Valid)

	if !ok || !token.Valid || err != nil {
		return errors.New("signature mismatch")
	}
	return nil
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

		var err error

		if authorizationInfo[0] == "Basic" {
			err = basicAuth(authorizationInfo[1])
		} else if authorizationInfo[0] == "Bearer" {
			err = bearerAuth(authorizationInfo[1])
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)

		return
	})
}
