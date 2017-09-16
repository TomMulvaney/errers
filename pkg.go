package errors

import (
	"strings"

	pkgErrors "github.com/pkg/errors"
)

// Wrap wraps pkgErrors.Wrap
func Wrap(err error, message string) error {
	return &nError{
		error:  pkgErrors.Wrap(err, message),
		status: StatusUnknown, // TODO: Get status from err param, if has no status then set StatusUnknown
	}
}

// WrapStatus ...
func WrapStatus(err error, status int, messages ...string) error {
	// Reverse messages so that the end of the slice (most recently appended) comes first
	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}

	message := strings.Join(messages, delim) // Use same delim and pkgErrors

	err = pkgErrors.Wrap(err, message)

	return &nError{
		error:  err,
		status: status,
	}
}
