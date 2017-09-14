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
