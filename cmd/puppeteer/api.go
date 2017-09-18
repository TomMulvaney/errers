package main

import (
	"github.com/nskeleton/errors"
)

const (
	isExist = true
)

type GRPCServer struct {
	errorHandlers []errors.Doer
}

func NewServer() *GRPCServer {
	var errorHandlers []errors.Doer

	return &GRPCServer{
		errorHandlers: errorHandlers,
	}
}

func (g *GRPCServer) CreatePuppet() {
	err := errors.New("Puppet already exists", errors.StatusBadReq)

	errors.HandleError(err)
}
