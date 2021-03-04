package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("AcáVaUnaFraseSecreta")

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido, lograste desbloquear la información dentro de ésta página, gracias a JWT que generaste!!")
}

// isAuthorized func => Creo una funcion para comprobar si está autorizado
// Si no está autorizado, dispara mensaje. Si no, le devuelve la info que se encuentra
// en 'homePage'.
// isAuthorized toma un request y valida si el token header está seteado, si no => no autorizado
// Si está, intenta parsear el tokenusando HMAC y nuestro signingKey
func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			// parseo el token header
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Hubo un 32202")
				}
				return mySigningKey, nil
			})
			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}

		} else {
			fmt.Fprintf(w, "NO autorizado kpo!")
		}
	})
}

func handleRequest() {
	http.Handle("/", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	fmt.Println("simple server")
	handleRequest()
}
