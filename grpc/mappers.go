package grpc

// Separate package because not everyone wants to import grpc

import (
	"github.com/nskeleton/errors"
	"google.golang.org/grpc/codes"
)

// FromGRPCStatus ...
func FromGRPCStatus(status int) int { // Calling function needs to convert to uint32
	return errors.StatusUnknown
}

// ToErrerStatus ...
func ToGRPCStatus(status int) int {
	return int(codes.Unknown)
}
