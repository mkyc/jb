package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleRequests() {
	http.HandleFunc("/", root)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func root(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprint(w, "root endpoint")
	log.Println("hit: root")
}

func main() {
	handleRequests()
}
