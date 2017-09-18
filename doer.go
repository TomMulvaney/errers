package errors

import (
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// Doer is a function executed (after converters) in HandleError
type Doer func(err error) error

////// Doers //////

// StatusMessage overwrites the error message with the status (obfuscating the internal working of a system)
func StatusMessage(err error) error {

	status := Status(err)

	statusMessage := strconv.Itoa(status)

	return New(statusMessage, status)
}

// ToHTTPStatus ...
func ToHTTPStatus(err error) error {
	httpStatus, ok := ToHTTPMap[Status(err)]

	if ok {
		return WrapStatus(err, httpStatus)
	}

	return Internal(err)
}

// Upstream converts error statuses received from upstream servers
func Upstream(err error) error {
	switch Status(err) {
	case StatusUnavailable:
		return UpstreamUnavailable(err)
	}

	return err
}

////// Doer Generators //////

// LogError ...
func LogError(logFields log.Fields) Doer { // This function is misnamed

	return func(err error) error {
		log.WithError(err).WithFields(logFields).Error("")
		return err
	}
}

// WriteHTTPHeader writes header from an NError status. Default http.InternalServerError
func WriteHTTPHeader(w http.ResponseWriter) Doer { // This function is misnamed

	return func(err error) error {
		status := http.StatusInternalServerError

		if IsNError(err) {
			e := err.(NError)
			status = e.GetStatus() // Call GetStatus because we want default to be http.StatusInternalServerError not StatusUnknown
		}

		w.WriteHeader(status)

		return err
	}
}
