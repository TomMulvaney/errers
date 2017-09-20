package main

import (
	"fmt"

	"github.com/nskeleton/errors"
)

// Error ...
type Error interface { // Stutter is for demo purposes, you should proably call this NError or something
	errors.NError
	isHypnosError() bool
}

// ErrorImp ...
type ErrorImp struct {
	*errors.NErrorImp
}

// NewHypnosError ...
func NewHypnosError(msg string, status int) error {
	return &ErrorImp{
		NErrorImp: errors.New(msg, status).(*errors.NErrorImp),
	}
}

func (e *ErrorImp) isHypnosError() bool {
	return true
}

// IsHypnosError ...
func IsHypnosError(err error) bool {
	e, ok := err.(Error)
	return ok && e.isHypnosError()
}

// HandleHypnosError ...
func HandleHypnosError(err error) error {
	if IsHypnosError(err) {
		// TODO

		fmt.Println("Handling Hypnos Error")

		return errors.NewAbortError()
	}

	return err
}
