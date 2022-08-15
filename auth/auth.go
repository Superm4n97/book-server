package auth

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/Superm4n97/Book-Server/model"
	"github.com/golang-jwt/jwt"
	"strings"
)

func BasicAuth(enStr string) error {
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

func BearerAuth(tokenStr string) error {
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

func GetBasicToken() string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", model.UNAME, model.UPASS)))
}

func GetBearerToken() string {
	return CreateJwtToken(model.UNAME, model.UPASS)
}
