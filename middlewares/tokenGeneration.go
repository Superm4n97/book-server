package middlewares

import (
	"fmt"
	"github.com/Superm4n97/Book-Server/model"
	"github.com/golang-jwt/jwt"
)

func CreateJwtToken(username string, password string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		//"foo": "bar",
		//"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		username: password,
	})

	fmt.Println("Secret key: ", model.ServerSecretKey)
	//tokenString, err := token.SignedString(model.ServerSecretKey)
	tokenString, _ := token.SignedString([]byte(model.ServerSecretKey))
	//fmt.Println(tokenString, err)
	return tokenString
}
