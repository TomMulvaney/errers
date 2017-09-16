package errors

// HandleError ...
func HandleError(err error, options ...Option) (error, bool) { // Ignore warning, errors should be the last return when they are not always expected. Here they are the very point of the function

	if IsNError(err) { // Is there a real advantage to doing this way instead of: if e, ok := err.(NError); ok {

		cfg := &handlerConfig{}
		for _, option := range options {
			option(cfg)
		}

		for _, convert := range cfg.converters {
			err = convert(err)
		}

		for _, do := range cfg.doers {
			do(err)
		}

		return err, true
	}

	return err, false
}
