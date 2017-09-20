package grpc

import (
	"github.com/nskeleton/errors"
	"google.golang.org/grpc/codes"
)

// Separate package because not everyone wants to import grpc

var (
	FromGRPCMap map[codes.Code]int // Populated in init

	ToGRPCMap = map[int]codes.Code{
		errors.StatusOK:                 codes.OK,
		errors.StatusUnknown:            codes.Unknown,
		errors.StatusCancelled:          codes.Canceled,
		errors.StatusBadReq:             codes.InvalidArgument,
		errors.StatusInvalidArgument:    codes.InvalidArgument,
		errors.StatusAlreadyExists:      codes.AlreadyExists,
		errors.StatusFailedPrecondition: codes.FailedPrecondition,
		errors.StatusUnauthenticated:    codes.Unauthenticated,
		errors.StatusForbidden:          codes.PermissionDenied,
		errors.StatusNotFound:           codes.NotFound,
		errors.StatusWrongAcceptType:    codes.Unknown,
		errors.StatusReqTimeout:         codes.DeadlineExceeded,
		errors.StatusTooManyReqs:        codes.Unknown,

		errors.StatusInternal:            codes.Internal,
		errors.StatusUnimplemented:       codes.Unimplemented,
		errors.StatusUpstreamUnavailable: codes.Unavailable,
		errors.StatusUnavailable:         codes.Unavailable,
		errors.StatusAborted:             codes.Aborted,
		errors.StatusDataLoss:            codes.DataLoss,
		errors.StatusResourceExhausted:   codes.ResourceExhausted,
	}
)

// ToGRPC ...
func ToGRPC(err error) error { // TODO: Use GRPC constructor
	// grpcStatus, ok := ToGRPCMap[errors.Status(err)]

	// if ok {
	// 	status.Error(grpcStatus, err.Message())
	// }

	// return status.Error(codes.Unknown, err.Message())

	return ToGRPCStatus(err)
}

// ToGRPCStatus ...
func ToGRPCStatus(err error) error {
	grpcStatus, ok := ToGRPCMap[errors.Status(err)]

	if ok {
		errors.WrapStatus(err, int(grpcStatus))
	}

	return errors.WrapStatus(err, int(codes.Unknown))
}

// FromGRPCStatus ...
func FromGRPCStatus(err error) error {
	grpcStatus := codes.Code(errors.StatusU32(err))

	status, ok := FromGRPCMap[grpcStatus]

	if ok {
		return errors.WrapStatus(err, status)
	}

	return errors.WrapStatus(err, errors.StatusUnknown)
}
