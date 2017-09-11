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

		return err, true
	}

	return err, false
}

type handlerConfig struct {
	mappers   []StatusMapper
	logFields *log.Fields // Pointer lets us do a null check to decide whether to log
}

// Option ...
type Option func(*handlerConfig)

// WithStatusMapper ...
func WithStatusMapper(mapper StatusMapper) Option {
	return func(cfg *handlerConfig) {
		cfg.mappers = append(cfg.mappers, mapper)
	}
}

// WithLogFields ...
func WithLogFields(logFields log.Fields) Option {
	return func(cfg *handlerConfig) {
		cfg.logFields = &logFields
	}
}
