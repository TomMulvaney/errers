package errors

import (
	"strings"

	pkgErrors "github.com/pkg/errors"
)

// Wrap wraps pkgErrors.Wrap
func Wrap(err error, message string) error {
	return &NErrorImp{
		error:  pkgErrors.Wrap(err, message),
		status: Status(err),
	}
}

// WrapStatus ...
func WrapStatus(err error, status int, messages ...string) error {

	if len(messages) > 0 {
		// Reverse messages so that the end of the slice (most recently appended) comes first
		for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
			messages[i], messages[j] = messages[j], messages[i]
		}

		message := strings.Join(messages, delim) // Use same delim and pkgErrors

		err = pkgErrors.Wrap(err, message)
	}

	return &NErrorImp{
		error:  err,
		status: status,
	}
}
