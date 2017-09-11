package errors

import (
	"fmt"

	pkgErrors "github.com/pkg/errors"
)

type ErrorStatus int

const (
	delim = ";"

	StatusBadReq ErrorStatus = iota
	StatusInternal
	StatusUnreachable
)

// New wraps standard library errors
func New(msg string) error {
	return pkgErrors.New(msg)
}

// Wrap wraps pkg error
func Wrap(err error, message string) error {
	return pkgErrors.Wrap(err, message)
}

// IError ...
type IError interface {
	error
	Status() ErrorStatus
	isIError() bool
}

type errorImp struct {
	error
	status ErrorStatus
}

// func (e *errorImp) Error() string {
// 	return e.Error()
// }

func (e *errorImp) Status() ErrorStatus {
	return e.status
}

func (e *errorImp) isIError() bool {
	return true
}

// IsIError ...
func IsIError(err error) bool {
	e, ok := err.(IError)
	return ok && e.isIError()
}

// NewStatus ...
func NewStatus(status ErrorStatus, err error, messages ...string) error {
	for _, message := range messages {
		err = pkgErrors.Wrap(err, message+delim)
	}

	return &errorImp{
		error:  err,
		status: status,
	}
}

// NewStatusTest ...
func NewStatusTest(status ErrorStatus, err error, message string) error {
	fmt.Println("NewStatusTest is wrapping")
	err = pkgErrors.Wrap(err, message)
	fmt.Println("NewStatusTest has wrapped")

	return &errorImp{
		error:  err,
		status: status,
	}
}

// BadReq constructor
func BadReq(err error, messages ...string) error {
	return NewStatus(StatusBadReq, err, messages...)
}

// Internal constructor
func Internal(err error, messages ...string) error {
	return NewStatus(StatusInternal, err, messages...)
}

// Unreachable constructor
func Unreachable(err error, messages ...string) error {
	return NewStatus(StatusUnreachable, err, messages...)
}
