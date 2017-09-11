package errors

// ErrorConverter ...
type ErrorConverter interface {
	Do(err error) error
}

type errorConverter struct {
	do errorConverterFunc
}

func (ec *errorConverter) Do(err error) error {
	return ec.do(err)
}

func NewErrorConverter(do errorConverterFunc) *errorConverter { // This errors when return interface
	return &errorConverter{
		do: do,
	}
}

type errorConverterFunc func(err error) error

// ToProxyConverter strips the error message for obfuscating the internal working of a system
func ToProxyError(err error) error {
	// TODO: Strip message
	return err
}
