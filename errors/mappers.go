package errors

import (
	"net/http"
)

// TODO: Deprecate StatusMapper, the same thing can be achieved with ErrorConverter.
// Using ErrorConverter is slightly less computationally efficient, but terser and more orthogonal
// Keep status mapping functions for error conversion

// StatusMapper ...
type StatusMapper interface {
	Do(status int) int
}

type statusMapper struct {
	do statusMapperFunc
}

// NewStatusMapper ...
func NewStatusMapper(do statusMapperFunc) StatusMapper {
	return &statusMapper{
		do: do,
	}
}

func (sm *statusMapper) Do(status int) int {
	return sm.do(status)
}

type statusMapperFunc func(status int) int

var (
	fromHTTP map[int]int // Populated in init

	toHTTP = map[int]int{
		StatusBadReq:              http.StatusBadRequest,
		StatusForbidden:           http.StatusForbidden,
		StatusNotFound:            http.StatusNotFound,
		StatusWrongAcceptType:     http.StatusNotAcceptable,
		StatusReqTimeout:          http.StatusRequestTimeout,
		StatusPreconditionFailed:  http.StatusPreconditionFailed,
		StatusTooManyReqs:         http.StatusTooManyRequests,
		StatusInternal:            http.StatusInternalServerError,
		StatusUnimplemented:       http.StatusNotImplemented,
		StatusUpstreamUnreachable: http.StatusBadGateway,
		StatusUnavailable:         http.StatusServiceUnavailable,
	}
)

// ToHTTPStatus ...
func ToHTTPStatus(status int) int {
	httpStatus, ok := toHTTP[status]

	if ok {
		return httpStatus
	}

	return http.StatusInternalServerError
}

// FromHTTPStatus ...
func FromHTTPStatus(httpStatus int) int {
	status, ok := fromHTTP[httpStatus]

	if ok {
		return status
	}

	return StatusUnknown
}

// ToErrerAPIStatus is for converting errer statuses for internal use to errer statuses for API clients
// For example, convert UpstreamUnreachable to Internal to obfuscate the system to end users
// Should this be done in the handlers?
func ToErrerAPIStatus(status int) int {
	return StatusUnknown
}
