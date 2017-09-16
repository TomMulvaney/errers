package errors

const (
	StatusUnknown int = iota

	// Client Errors
	StatusBadReq
	StatusUnauthenticated
	StatusForbidden
	StatusNotFound
	StatusWrongAcceptType
	StatusReqTimeout
	StatusPreconditionFailed
	StatusTooManyReqs

	// Server Errors
	StatusInternal
	StatusUnimplemented
	StatusUpstreamUnreachable
	StatusUnavailable

	delim = ": " // Same delim as pkgErrors
)

// Unknown constructor
func Unknown(err error, messages ...string) error { // Variadic params make string messages optional
	return WrapStatus(err, StatusUnknown, messages...)
}

// BadReq constructor
func BadReq(err error, messages ...string) error { // Variadic params make string messages optional
	return WrapStatus(err, StatusBadReq, messages...)
}

// Internal constructor
func Internal(err error, messages ...string) error {
	return WrapStatus(err, StatusInternal, messages...)
}

// UpstreamUnreachable constructor
func UpstreamUnreachable(err error, messages ...string) error {
	return WrapStatus(err, StatusUpstreamUnreachable, messages...)
}

// Unimplemented constructor
func Unimplemented(err error, messages ...string) error {
	return WrapStatus(err, StatusUnimplemented, messages...)
}
