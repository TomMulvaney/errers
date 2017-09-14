package hypnos

import (
	"github.com/nskeleton/errors"
)

func hypnosToHTTPStatus(status int) int {
	return errors.StatusUnknown
}

// HandleHypnosError ...
func HandleHypnosError(err error) (error, bool) {
	if IsHypnosError(err) {
		e := err.(HypnosError)
		err = errors.WrapStatus(e, hypnosToHTTPStatus(e.Status()))

		return err, true
	}

	return err, false
}
