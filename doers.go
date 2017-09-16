package errors

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Doer is a function executed (after converters) in HandleError
type Doer func(err error)

// LogError ...
func LogError(logFields log.Fields) Option { // This function is misnamed

	return func(cfg *handlerConfig) {

		do := func(err error) {
			log.WithError(err).WithFields(logFields).Error("")
		}

		cfg.doers = append(cfg.doers, do)
	}
}

// WriteHeader ...
func WriteHeader(w http.ResponseWriter, defaultHeader int) Option { // This function is misnamed

	return func(cfg *handlerConfig) {

		do := func(err error) {
			status := defaultHeader

			if IsNError(err) {
				e := err.(NError)
				status = e.Status()
			}

			w.WriteHeader(status)
		}

		cfg.doers = append(cfg.doers, do)
	}
}
