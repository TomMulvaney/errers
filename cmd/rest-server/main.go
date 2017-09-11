package main

import (
	"fmt"

	"github.com/TomMulvaney/errors"
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

func fooAPIHandler() error {

	err := barClient()

	if err != nil {
		return errors.Wrap(err, "Failed doing something in Bar")
	}

	return nil
}

func fooAPIMiddleware() {

	err := fooAPIHandler()

	err = errors.HandleCommon(err)

	fmt.Println("Returning err: ", err)

	// TODO: Write HTTP status
}

func main() {
	fmt.Println("Hello World! I am a dummy REST API")

	fmt.Println("I have just received a request")

	// err := barServer()
	// fmt.Println(err.Error())

	errors.Test4()

	// fooAPIMiddleware()
}
