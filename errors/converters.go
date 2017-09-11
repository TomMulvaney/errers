package errors

// ErrorConverter ...
type ErrorConverter interface {
	Do(err error) error
}

type errorConverter struct {
	Do errorConverterFunc
}

func NewErrorConverter(do errorConverterFunc) *errorConverter { // This errors when return interface
	return &errorConverter{
		Do: do,
	}
}

type errorConverterFunc func(err error) error

// ToProxyConverter strips the error message for obfuscating the internal working of a system
func ToProxyError(err error) error {
	// TODO: Strip message
	return err
}
