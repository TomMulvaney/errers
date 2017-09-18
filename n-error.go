package errors

import (
	pkgErrors "github.com/pkg/errors"
)

// NError ...
type NError interface {
	error
	GetStatus() int // Status represents NError, GRPC, HTTP status
	isNError() bool // Is this an NError?
}

// NErrorImp ...
type NErrorImp struct {
	error
	status int
}

// New ...
func New(msg string, status int) error {
	return &NErrorImp{
		error:  pkgErrors.New(msg),
		status: status,
	}
}

// GetStatus ...
func (e *NErrorImp) GetStatus() int {
	return e.status
}

// Status returns status from NError, and StatusUnknown from other errors
func Status(err error) int {
	if IsNError(err) {
		e := err.(NError)
		return e.GetStatus()
	}

	return StatusUnknown
}

// StatusU32 returns status as uint32 (for grpc)
func StatusU32(err error) uint32 {
	return uint32(Status(err))
}

// TODO: Find out whether conflict if export this
func (e *NErrorImp) isNError() bool {
	return true
}

// IsNError ...
func IsNError(err error) bool {
	e, ok := err.(NError)
	return ok && e.isNError()
}
