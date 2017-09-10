package errors

import (
	"strings"

	basIErrors "errors"

	"github.com/pkg/errors"
)

const (
	delim = ";"
)

// New wraps standard library errors
func New(msg string) error {
	return basIErrors.New(msg)
}

// Wrap wraps pkg error
func Wrap(err error, messages ...string) error {
	msg := strings.Join(messages, delim)
	return errors.Wrap(err, msg)
}

// TODO: Rename to Errar, and errar

//// IError ...

// IError ...
type IError interface {
	error
	isIError() bool
}

// IError ... Abstract (no constructor)?
type errar struct {
	error
}

func (e *errar) isIError() bool {
	return true
}

func (e *errar) Error() string {
	return e.Error()
}

// IsIError ...
func IsIError(err error) bool {
	e, ok := err.(IError)
	return ok && e.isIError()
}

//// BadReq ...

// IBadReq ...
type IBadReq interface {
	IError
	isBadReq() bool
}

type badReq struct {
	IError
}

// BadReq constructor
func BadReq(err error, message string) error {
	return Wrap(&badReq{}, strings.Join([]string{err.Error(), message}, delim)) // TODO: This could probably be more computationally efficient
}

// isBadReq method
func (*badReq) isBadReq() bool {
	return true
}

// IsBadReq function
func IsBadReq(err error) bool {
	e, ok := err.(IBadReq)
	return ok && e.isBadReq()
}

//// Internal

// IInternal ...
type IInternal interface {
	IError
	isInternal() bool
}

type internal struct {
	IError
}

// Internal constructor
func Internal(err error, message string) error {
	return Wrap(&badReq{}, strings.Join([]string{err.Error(), message}, delim))
}

// IUnreachable ...
type IUnreachable interface {
	IError
	isUnreachable() bool
}

type unreachable struct {
	IError
}

// Unreachable constructor
func Unreachable(err error, message string) error {
	return Wrap(&badReq{}, strings.Join([]string{err.Error(), message}, delim))
}

func (*unreachable) isUnreachable() bool {
	return true
}

// IsUnreachable ...
func IsUnreachable(err error) bool {
	e, ok := err.(IUnreachable)
	return ok && e.isUnreachable()
}
