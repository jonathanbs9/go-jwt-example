package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("AcáVaUnaFraseSecreta")

// GenerateJWT func => función que genera un Json Web Token!
func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Jonathan Brull Schroeder"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("Algo anduvo mal =/  ... %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func main() {
	fmt.Println("Simple")
	tokenString, err := GenerateJWT()
	if err != nil {
		fmt.Println("Can't ! ")
	}
	fmt.Println(tokenString)
}
