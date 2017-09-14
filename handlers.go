package errors

import (
	log "github.com/sirupsen/logrus"
)

// TODO: Give option of method. Method is more computationally efficient because cfg is only instantiated once
// Does it break orthogonality to allow both function and method? Maybe, but function is more convenient
// while developing

// HandleErrer ...
func HandleErrer(err error, options ...Option) (error, bool) { // Ignore warning, errors should be the last return when they are not always expected. Here they are the very point of the function
	if IsErrer(err) {
		e := err.(IErrer)

		cfg := &handlerConfig{}
		for _, option := range options {
			option(cfg)
		}

		for _, converter := range cfg.converters {
			err = converter.Do(e)
		}

		if IsErrer(err) {
			e = err.(IErrer)
		} else {
			log.WithError(err).Warn("Converted error is not IErrer, discarding results of conversion")
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
