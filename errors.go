package errors

import (
	"strings"

	baseErrors "errors"

	"github.com/pkg/errors"
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
	return baseErrors.New(msg)
}

// Wrap wraps pkg error
func Wrap(err error, messages ...string) error {
	msg := strings.Join(messages, delim)
	return errors.Wrap(err, msg)
}

// IError ...
type IError interface {
	error
	Status() ErrorStatus
}

type errorImp struct {
	error
	status ErrorStatus
}

func (e *errorImp) Error() string {
	return e.Error()
}

func (e *errorImp) Status() ErrorStatus {
	return e.status
}

func new(status ErrorStatus, err error, messages ...string) error {
	if len(messages) > 0 {
		for _, message := range messages {
			err = Wrap(err, message)
		}
	}

	return &errorImp{
		error:  err,
		status: status,
	}
}

// BadReq constructor
func BadReq(err error, messages ...string) error {
	return new(StatusBadReq, err, messages...)
}

// Internal constructor
func Internal(err error, messages ...string) error {
	return new(StatusInternal, err, messages...)
}

// Unreachable constructor
func Unreachable(err error, messages ...string) error {
	return new(StatusUnreachable, err, messages...)
}
