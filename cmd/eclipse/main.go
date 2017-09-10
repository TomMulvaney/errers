package main

import (
	"fmt"

	"github.com/TomMulvaney/errors"
	log "github.com/sirupsen/logrus"
)

func craterServerGetGroup() error {
	return errors.BadReq(errors.New("Crater has message about what went wrong"), "Crater Server")
}

func craterClientGetGroup() error {
	fmt.Println("Calling Crater")

	err := craterServerGetGroup()

	if err != nil {
		return errors.Wrap(err, "Failed reading group from Crater")
	}

	return nil
}

func galaxyAPIHandler() error {
	// TODO: Call Crater, get error
	// Add self to lineage
	// Relay crater error
	return nil
}

func convertError(err error) error {
	e, ok := err.(errors.IError)
	if ok {
		switch e.Status() {
		case errors.StatusUnreachable: // Convert Unreachable to Internal
			err = errors.Internal(e)
		}
	}

	return err
}

func handleError(err error) error {
	// Log errors here, eases burden on handlers
	log.WithError(err).Error("Failed Request")

	err = convertError(err)

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
