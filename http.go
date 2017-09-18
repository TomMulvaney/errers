package errors

import (
	"net/http"
)

var (
	FromHTTPMap map[int]int // Populated in init

	ToHTTPMap = map[int]int{
		StatusOK:                 http.StatusOK,
		StatusCancelled:          http.StatusInternalServerError, // TODO
		StatusBadReq:             http.StatusBadRequest,
		StatusInvalidArgument:    http.StatusBadRequest,
		StatusAlreadyExists:      http.StatusBadRequest,
		StatusFailedPrecondition: http.StatusPreconditionFailed,
		StatusUnauthenticated:    http.StatusUnauthorized,
		StatusForbidden:          http.StatusForbidden,
		StatusNotFound:           http.StatusNotFound,
		StatusWrongAcceptType:    http.StatusNotAcceptable,
		StatusReqTimeout:         http.StatusRequestTimeout,
		StatusTooManyReqs:        http.StatusTooManyRequests,

		StatusInternal:            http.StatusInternalServerError,
		StatusUnimplemented:       http.StatusNotImplemented,
		StatusUpstreamUnavailable: http.StatusBadGateway,
		StatusUnavailable:         http.StatusServiceUnavailable,
		StatusAborted:             http.StatusInternalServerError,
		StatusDataLoss:            http.StatusInternalServerError,
		StatusResourceExhausted:   http.StatusInternalServerError,
	}
)
