package errors

// AbortError ...
type AbortError interface {
	error
	isAbortError() bool
}

// AbortErrorImp ...
type AbortErrorImp struct {
	error
}

// NewAbortError ...
func NewAbortError() error {
	return &AbortErrorImp{}
}

func (e *AbortErrorImp) isAbortError() bool {
	return true
}

// IsAbortError ...
func IsAbortError(err error) bool {
	e, ok := err.(AbortError)
	return ok && e.isAbortError()
}
