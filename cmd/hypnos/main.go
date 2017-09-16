package main

import (
	"fmt"
	"net/http"

	"github.com/nskeleton/errors"
	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Hello World! I am a dummy REST API")
	fmt.Println("I am called Hypnos, after the greek goddess of sleep")

	fmt.Println("Received request")

	err := errors.New("Nightmare", errors.StatusUnknown)

	err = errors.Wrap(err, "Nightmare about not being able to wake up from the nightmare")

	err = errors.Internal(err, "We are experiencing some technical difficulties", "Please try to relax with this smooth jazz")

	var options []errors.Option

	// TODO: Convert to HTTP Status

	w := NewResponseWriter()
	options = append(options, errors.WriteHeader(w, http.StatusInternalServerError))

	logFields := log.Fields{
		"Hello": "World",
	}
	options = append(options, errors.LogError(logFields))

	errors.HandleError(err, options...)
}
