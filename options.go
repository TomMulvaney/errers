package errors

type handlerConfig struct {
	converters []Converter
	doers      []Doer
}

// Option ...
type Option func(*handlerConfig)

// Convert ...
func Convert(converter Converter) Option {
	return func(cfg *handlerConfig) {
		cfg.converters = append(cfg.converters, converter)
	}
}

// Do ...
func Do(doer Doer) Option {
	return func(cfg *handlerConfig) {
		cfg.doers = append(cfg.doers, doer)
	}
}
