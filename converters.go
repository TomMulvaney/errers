package errors

import (
	"strconv"
)

// TODO: Deprecate - Doer should be type func(error) error

// Converter is a function executed (before doers) in HandleError
type Converter func(err error) error // The functions below could be terser if this has NError param and return type

// StatusMessage overwrites the error message with the status (obfuscating the internal working of a system)
func StatusMessage() Option {

	return func(cfg *handlerConfig) {
		converter := func(err error) error {
			if IsNError(err) {
				e := err.(NError)

				statusMessage := strconv.Itoa(e.Status())

				return New(statusMessage, e.Status())
			}

			// TODO: Log warning?

			return err
		}

		cfg.converters = append(cfg.converters, converter)
	}

}

// ToHTTPStatus ...
func ToHTTPStatus() Option {

	return func(cfg *handlerConfig) {
		converter := func(err error) error {
			if IsNError(err) {
				e := err.(NError)

				httpStatus, ok := ToHTTPMap[e.Status()]

				if ok {
					return WrapStatus(e, httpStatus)
				}
			}

			return Internal(err)
		}

		cfg.converters = append(cfg.converters, converter)
	}

}

// ToErrorAPIStatus is for converting error statuses for internal use to error statuses for API clients
// For example, convert UpstreamUnavailable to Internal to obfuscate the system to end users
// Should this be done in the handlers?
func ToErrorAPIStatus(status int) int {
	return StatusUnknown
}
