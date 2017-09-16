package main

import (
	"fmt"
	"net/http"
	"path"

	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Hello World! I am the Hypnos REST API")
	fmt.Println("I am named after the greek god of sleep (REST, get it?)")

	fmt.Println("Received GET request for  dream")

	w := NewResponseWriter()

	r, err := http.NewRequest("GET", path.Join(urlBase, "dreams/happy-place"), nil)

	if err != nil {
		log.WithError(err).Fatal("Failed creating dummy req")
	}

	APIMiddleware(w, r)
}
