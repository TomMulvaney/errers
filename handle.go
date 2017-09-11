package errors

import (
	log "github.com/sirupsen/logrus"
)

// ConvertCommon ...
func ConvertCommon(err error) error {
	e, ok := err.(IError)
	if ok {
		switch e.Status() {
		case StatusUnreachable: // Convert Unreachable to Internal
			err = Internal(e)
		}
	}

	return err
}

// HandleCommon ...
func HandleCommon(err error) error {
	// Log errors here, eases burden on handlers
	log.WithError(err).Error("Error")

	return ConvertCommon(err)
}
