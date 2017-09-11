package main

import (
	"fmt"

	"github.com/TomMulvaney/errers/errors"
)

func barServer() error {
	err := errors.New("Something, something, low-level error")
	fmt.Println(err)
	return errors.BadReq(err, "Message from Bar Server API Handler")
}

func barClient() error {
	fmt.Println("Calling Bar server")

	err := barServer()

	if err != nil { // This error should already be converted by network library that client depends on
		return err // Probably no need to wrap this error, but we could wrap a message
	}

	return nil
}

func hypnosAPIHandler() error {

	err := barClient()

	if err != nil {
		return errors.Wrap(err, "Failed doing something in Bar")
	}

	return nil
}

func handleHypnosError(err error) (error, bool) {
	return err, false
}

func hypnosAPIMiddlerware() {

	err := hypnosAPIHandler()

	err, ok := handleHypnosError(err) // This bit is optional, maybe there is no HypnosError type

	if !ok {
		err, ok = errors.HandleIErrer(err)

		if !ok {
			// Log Warning
			err = errors.Unknown(err, "Error is unknown type. Consider updating the Hypnos error handler")
		}
	}

	httpStatus := errors.ErrerToHTTP(err)

	fmt.Println("Writing Header: ", httpStatus)

	fmt.Println("Returning err: ", err)
}

func main() {
	fmt.Println("Hello World! I am a dummy REST API")
	fmt.Println("I am called Hypnos, after the greek goddess of sleep")

	fmt.Println("Received request")

	err := errors.New("Nightmare")

	err = errors.Wrap(err, "Nightmare about not being able to wake up from the nightmare")

	err = errors.Internal(err, "We are experiencing some technical difficulties", "Please try to relax with this smooth jazz")

	fmt.Println(err)
}
