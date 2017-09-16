package errors

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Doer ...
type Doer func(err error)

// LogError ...
func LogError(logFields log.Fields) Option {

	return func(cfg *handlerConfig) {

		do := func(err error) {
			log.WithError(err).WithFields(logFields).Error("Error")
		}

		cfg.doers = append(cfg.doers, do)
	}
}

// WriteHeader ...
func WriteHeader(w http.ResponseWriter, defaultHeader int) Option {

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
