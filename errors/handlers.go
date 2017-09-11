package errors

import (
	log "github.com/sirupsen/logrus"
)

// HandleErrer ...
func HandleErrer(err error, options ...Option) (error, bool) { // Ignore warning, errors should be the last return when they are not always expected. Here they are the very point of the function
	if IsErrer(err) {
		e := err.(IErrer)

		cfg := &handlerConfig{}
		for _, option := range options {
			option(cfg)
		}

		status := e.Status()
		for _, mapper := range cfg.mappers {
			status = mapper.Do(status)
		}
		err = WrapStatus(e, status)

		if cfg.logFields != nil {
			log.WithError(err).WithFields(*cfg.logFields).Error("E") // TODO: Maybe add Last() to IErrer
		}

		if cfg.overwriteMessageWithStatus {
			// TODO: Overwrite message with status converted to string
		}

		return err, true
	}

	return err, false
}
