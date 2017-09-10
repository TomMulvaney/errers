package main

import (
	"fmt"

	"github.com/TomMulvaney/errors/cheney/errors"
	// TODO: Use pkg/errors.Wrap
)

func httpMiddleware() error {
	return errors.Unreachable() // TODO
}

func callCrater() error {
	fmt.Println("Calling Crater")
	fmt.Println("Crater is Unreachable")

	// Relay crater error with Wrap

	// TODO: Return Unreachable
	return nil
}

func galaxyAPIHandler() error {
	// TODO: Call Crater, get error
	// Add self to lineage
	// Relay crater error
	return nil
}

func handleError(err error) error {
	// Log error

	// Convert: unreachable to Internal
	// Relay: badReq and Internal

	// Write HTTP status
	return nil
}

func main() {
	fmt.Println("Hello World! I am a dummy Galaxy server")

	fmt.Println("I have just received a request")

	err := galaxyAPIHandler()

	err = handleError(err)

	fmt.Println("Writing error: ", err)
}
