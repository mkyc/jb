package main

import (
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var key = []byte("someverystrongsigningkey")

func main() {
	client := &http.Client{}
	for {
		err := poke(client, "http://server:8080/", false)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(1 * time.Second)
	}
}

func poke(c *http.Client, url string, authorized bool) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	if authorized {
		token, err := generateJWT()
		if err != nil {
			return err
		}
		req.Header.Set("Token", token)
	}

	res, err := c.Do(req)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	log.Println(string(body))
	return nil
}

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["client"] = "Someone"
	claims["exp"] = time.Now().Add(1 * time.Minute).Unix()

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
