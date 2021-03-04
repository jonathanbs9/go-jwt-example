package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// var mySigningKey = os.getEnv("MY_JWT_TOKEN")
// En consola = set MY_JWT_TOKEN=AcáVaUnaFraseSecreta
var mySigningKey = []byte("AcáVaUnaFraseSecreta")

func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:9000/", nil)
	req.Header.Set("Token", validToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error => %s", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, string(body))
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":9001", nil))
}

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
	handleRequest()
}
