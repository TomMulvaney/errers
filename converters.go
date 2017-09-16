package errors

import (
	"strconv"
)

// Converter ...
type Converter func(err error) error // The functions below could be terser if this has NError param and return type

// TODO: These functions should return Option

// StatusMessage overwrites the error message with the status (obfuscating the internal working of a system)
func StatusMessage(err error) error {

	if IsNError(err) {
		e := err.(NError)

		statusMessage := strconv.Itoa(e.Status())

		return New(statusMessage, e.Status())
	}

	// TODO: Log warning?

	return err
}

// ToHTTPStatus ...
func ToHTTPStatus(err error) error {

	if IsNError(err) {
		e := err.(NError)

		httpStatus, ok := toHTTP[e.Status()]

		if ok {
			return WrapStatus(e, httpStatus)
		}
	}

	return Internal(err)
}

// ToErrorAPIStatus is for converting error statuses for internal use to error statuses for API clients
// For example, convert UpstreamUnreachable to Internal to obfuscate the system to end users
// Should this be done in the handlers?
func ToErrorAPIStatus(status int) int {
	return StatusUnknown
}
