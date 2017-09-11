package grpc

// Separate package because not everyone wants to import grpc

import (
	"github.com/TomMulvaney/errers/errors"
)

// FromGRPCStatus ...
func FromGRPCStatus(status int) int { // Calling function needs to convert to uint32
	return errors.StatusUnknown
}

// ToErrerStatus ...
func ToErrerStatus(status int) int {
	return errors.StatusUnknown
}
