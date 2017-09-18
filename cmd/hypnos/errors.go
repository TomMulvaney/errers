package main

import (
	"github.com/nskeleton/errors"
)

// Error ...
type Error interface { // Stutter is for demo purposes, you should proably call this NError or something
	errors.NError
	isHypnosError() bool
}

// ErrorImp ...
type ErrorImp struct {
	errors.NErrorImp
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

		// e := err.(HypnosError)
		// err = errors.WrapStatus(e, hypnosToHTTPStatus(e.Status()))
	}

	return err
}
