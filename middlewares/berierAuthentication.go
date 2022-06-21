package middlewares

import (
	"fmt"
	"github.com/Superm4n97/Book-Server/model"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

func BearerAuthentication(req http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		inputString := r.Header.Get("Authorization")
		fmt.Println(inputString)

		infoString := strings.Split(inputString, " ")

		if len(infoString) != 2 || infoString[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("invalid bearer token"))
			return
		}

		tokenStr := infoString[1]

		token, err := jwt.Parse(tokenStr, func(tkn *jwt.Token) (interface{}, error) {
			if _, ok := tkn.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", tkn.Header["alg"])
			}

			return []byte(model.ServerSecretKey), nil
		})

		if clms, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			//fmt.Println("Token Matched !!!!")
			//w.Write([]byte("Token Matched, valid user"))
			fmt.Println(clms["userName"])
		} else {
			fmt.Println(err)
		}
	})
}
