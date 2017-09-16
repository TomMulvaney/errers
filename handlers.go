package errors

// HandleError ...
// TODO: func HandleError(err error, doers ...Doer) error {
func HandleError(err error, options ...Option) error { // Ignore warning, errors should be the last return when they are not always expected. Here they are the very point of the function

	cfg := &handlerConfig{} // TODO: Deprecate Option
	for _, option := range options {
		option(cfg)
	}

	for _, convert := range cfg.converters {
		err = convert(err)
	}

	for _, do := range cfg.doers { // Only doers but make type func(error) error
		do(err)
	}

	return err
}
