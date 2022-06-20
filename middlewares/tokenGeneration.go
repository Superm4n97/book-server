package middlewares

import (
	"fmt"
	"github.com/Superm4n97/Book-Server/info"
	"github.com/golang-jwt/jwt"
)

func CreateJwtToken(username string, password string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		//"foo": "bar",
		//"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"userName": username,
		"pass":     password,
	})

	fmt.Println("Secret key: ", info.ServerSecretKey)
	//tokenString, err := token.SignedString(info.ServerSecretKey)
	tokenString, _ := token.SignedString([]byte(info.ServerSecretKey))
	//fmt.Println(tokenString, err)
	return tokenString
}
