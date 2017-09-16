package errors

import (
	pkgErrors "github.com/pkg/errors"
)

// NError ...
type NError interface {
	error
	Status() int       // Status represents NError, GRPC, HTTP status
	StatusU32() uint32 // github.com/grpc/grpc-go use uint32
	isNError() bool    // Is this an NError?
}

// nError ...
type nError struct {
	error
	status int
}

// New ...
func New(msg string, status int) error {
	return &nError{
		error:  pkgErrors.New(msg),
		status: status,
	}
}

// Status ...
func (e *nError) Status() int {
	return e.status
}

// StatusU32 returns status as uint32 (for grpc)
func (e *nError) StatusU32() uint32 {
	return uint32(e.status)
}

// TODO: Find out whether conflict if export this
func (e *nError) isNError() bool {
	return true
}

// IsNError ...
func IsNError(err error) bool {
	e, ok := err.(NError)
	return ok && e.isNError()
}
