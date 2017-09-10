package errors

import (
	"strings"

	baseErrors "errors"

	"github.com/pkg/errors"
)

const (
	delim = ";"
)

// New wraps standard library errors
func New(msg string) error {
	return baseErrors.New(msg)
}

// Wrap wraps pkg error
func Wrap(err error, messages ...string) error {
	msg := strings.Join(messages, delim)
	return errors.Wrap(err, msg)
}

type ErrorStatus int

//// IError ...

// IError ...
type IError interface {
	error
	isError() bool
	Status() ErrorStatus
}

type errorImp struct {
	error
	status ErrorStatus
}

func (e *errorImp) isError() bool {
	return true
}

func (e *errorImp) Error() string {
	return e.Error()
}

func (e *errorImp) Status() ErrorStatus {
	return e.status
}

// IsError ...
func IsError(err error) bool {
	e, ok := err.(IError)
	return ok && e.isError()
}

// TODO

// BadReq constructor
func BadReq(err error, message string) error {
	return Wrap(&errorImp{
		status: 
	}, strings.Join([]string{err.Error(), message}, delim)) // TODO: This could probably be more computationally efficient
}

// Internal constructor
func Internal(err error, message string) error {
	return Wrap(&internal{}, strings.Join([]string{err.Error(), message}, delim))
}

// Unreachable constructor
func Unreachable(err error, message string) error {
	return Wrap(&unreachable{}, strings.Join([]string{err.Error(), message}, delim))
}
