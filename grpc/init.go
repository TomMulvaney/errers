package grpc

import (
	"google.golang.org/grpc/codes"
)

// Separate package because not everyone wants to import grpc

func init() {
	FromGRPCMap = make(map[codes.Code]int, len(ToGRPCMap))

	for k, v := range ToGRPCMap {
		FromGRPCMap[v] = k
	}
}
