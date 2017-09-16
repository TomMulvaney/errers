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
	opt := errors.WriteHeader(w, http.StatusInternalServerError)
	options = append(options, opt)

	logFields := log.Fields{
		"Hello": "World",
	}
	opt = errors.LogError(logFields)
	options = append(options, opt)

	errors.HandleError(err, options...)
}
