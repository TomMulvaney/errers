package errors

import (
	log "github.com/sirupsen/logrus"
)

type handlerConfig struct {
	converters                 []ErrorConverter
	mappers                    []StatusMapper
	logFields                  *log.Fields // Pointer lets us do a null check to decide whether to log
	overwriteMessageWithStatus bool
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

// WithStatusMessage overwrites the error's message with the error's status
func WithStatusMessage() Option {
	return func(cfg *handlerConfig) {
		cfg.overwriteMessageWithStatus = true
	}
}
