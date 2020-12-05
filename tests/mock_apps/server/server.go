package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
)

var key = []byte("someverystrongsigningkey")

func main() {
	http.Handle("/", ifAuthorized(root))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func root(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprint(w, "root endpoint")
	log.Println("hit: root")
}

func ifAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("incorrect casting")
				}
				return key, nil
			})
			if err != nil {
				_, _ = fmt.Fprintf(w, err.Error())
			}
			if token != nil && token.Valid {
				claims := token.Claims.(jwt.MapClaims)
				fmt.Println("client: " + claims["client"].(string))
				endpoint(w, r)
			}
		} else {
			_, _ = fmt.Fprintf(w, "Not Authorized")
		}
	})
}
