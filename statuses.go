package errors

const (
	StatusUnknown int = iota

	StatusOK

	StatusCancelled

	// Client Errors
	StatusBadReq
	StatusInvalidArgument
	StatusAlreadyExists
	StatusFailedPrecondition
	StatusUnauthenticated
	StatusForbidden
	StatusNotFound
	StatusWrongAcceptType
	StatusReqTimeout
	StatusTooManyReqs

	// Server Errors
	StatusInternal
	StatusUnimplemented
	StatusUpstreamUnavailable
	StatusUnavailable
	StatusAborted
	StatusDataLoss
	StatusResourceExhausted

	delim = ": " // Same delim as pkgErrors
)

// Unknown constructor
func Unknown(err error, messages ...string) error { // Variadic params make string messages optional
	return WrapStatus(err, StatusUnknown, messages...)
}

// BadReq constructor
func BadReq(err error, messages ...string) error {
	return WrapStatus(err, StatusBadReq, messages...)
}

// Unauthenticated constructor
func Unauthenticated(err error, messages ...string) error {
	return WrapStatus(err, StatusUnauthenticated, messages...)
}

// Forbidden constructor
func Forbidden(err error, messages ...string) error {
	return WrapStatus(err, StatusForbidden, messages...)
}

// NotFound constructor
func NotFound(err error, messages ...string) error {
	return WrapStatus(err, StatusNotFound, messages...)
}

// WrongAcceptType constructor
func WrongAcceptType(err error, messages ...string) error {
	return WrapStatus(err, StatusWrongAcceptType, messages...)
}

// ReqTimeout constructor
func ReqTimeout(err error, messages ...string) error {
	return WrapStatus(err, StatusReqTimeout, messages...)
}

// FailedPrecondition constructor
func FailedPrecondition(err error, messages ...string) error {
	return WrapStatus(err, StatusFailedPrecondition, messages...)
}

// TooManyReqs constructor
func TooManyReqs(err error, messages ...string) error {
	return WrapStatus(err, StatusTooManyReqs, messages...)
}

// Internal constructor
func Internal(err error, messages ...string) error {
	return WrapStatus(err, StatusInternal, messages...)
}

// UpstreamUnavailable constructor
func UpstreamUnavailable(err error, messages ...string) error {
	return WrapStatus(err, StatusUpstreamUnavailable, messages...)
}

// Unimplemented constructor
func Unimplemented(err error, messages ...string) error {
	return WrapStatus(err, StatusUnimplemented, messages...)
}

// Unavailable constructor
func Unavailable(err error, messages ...string) error {
	return WrapStatus(err, StatusUnavailable, messages...)
}
