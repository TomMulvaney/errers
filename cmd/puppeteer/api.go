package main

import (
	"github.com/nskeleton/errors"
	errorsGRPC "github.com/nskeleton/errors/grpc"
	log "github.com/sirupsen/logrus"
)

const (
	isExist = true
)

// GRPCServer ...
type GRPCServer struct {
}

// NewServer ...
func NewServer() *GRPCServer {
	return &GRPCServer{}
}

// HandleError ...
func (g *GRPCServer) HandleError(err error, method string) {
	var doers []errors.Doer

	logFields := log.Fields{
		"Server Method": method,
	}

	doers = append(doers, errors.LogError(logFields))

	doers = append(doers, errors.Upstream)

	doers = append(doers, errorsGRPC.ToGRPC)

	return errors.HandleError(err, doers)
}

// CreatePuppet ...
func (g *GRPCServer) CreatePuppet() {
	err := errors.New("Puppet already exists", errors.StatusBadReq)

	g.HandleError(err, "CreatePuppet")
}
