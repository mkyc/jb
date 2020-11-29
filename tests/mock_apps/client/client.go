package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func poke(c *http.Client, url string) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
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

func main() {
	client := &http.Client{}
	for {
		err := poke(client, "http://server:8080/")
		if err != nil {
			log.Println(err)
		}
		time.Sleep(1 * time.Second)
	}
}
