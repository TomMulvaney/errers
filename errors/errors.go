package errors

import (
	"strings"

	pkgErrors "github.com/pkg/errors"
)

const (
	StatusUnknown int = iota
	StatusBadReq
	StatusInternal
	StatusUnreachable
	StatusUnimplemented

	delim = ": " // Same delim as pkgErrors
)

// IErrer ...
type IErrer interface {
	error
	Status() int          // Status can represent IErrer or HTTP Status
	StatusUInt32() uint32 // github.com/grpc/grpc-go use uint32
	isErrer() bool
}

// Errer ...
type Errer struct {
	error
	status int
}

// New wraps pkgErrors.New (which wraps baseError.New)
func New(msg string) error {
	return &Errer{
		error:  pkgErrors.New(msg),
		status: StatusUnknown,
	}
}

// Wrap wraps pkgErrors.Wrap
func Wrap(err error, message string) error {
	return &Errer{
		error:  pkgErrors.Wrap(err, message),
		status: StatusUnknown, // TODO: Convert from HTTP and GRPC
	}
}

// WrapStatus ...
func WrapStatus(err error, status int, messages ...string) error { // Variadic params makes messages optional
	// Reverse messages so that the end of the slice (most recently appended) comes first
	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}

	message := strings.Join(messages, delim) // Use same delim and pkgErrors

	err = pkgErrors.Wrap(err, message)

	return &Errer{
		error:  err,
		status: status,
	}
}

// Status ...
func (e *Errer) Status() int {
	return e.status
}

// StatusU32 ...
func (e *Errer) StatusU32() uint32 {
	return uint32(e.status)
}

// isErrer ...
func (e *Errer) isErrer() bool {
	return true
}

// IsErrer ...
func IsErrer(err error) bool {
	e, ok := err.(IErrer)
	return ok && e.isErrer()
}

// Unknown constructor
func Unknown(err error, messages ...string) error { // Variadic params make string messages optional
	return WrapStatus(err, StatusUnknown, messages...)
}

// BadReq constructor
func BadReq(err error, messages ...string) error { // Variadic params make string messages optional
	return WrapStatus(err, StatusBadReq, messages...)
}

// Internal constructor
func Internal(err error, messages ...string) error {
	return WrapStatus(err, StatusInternal, messages...)
}

// Unreachable constructor
func Unreachable(err error, messages ...string) error {
	return WrapStatus(err, StatusUnreachable, messages...)
}

// Unimplemented constructor
func Unimplemented(err error, messages ...string) error {
	return WrapStatus(err, StatusUnimplemented, messages...)
}
