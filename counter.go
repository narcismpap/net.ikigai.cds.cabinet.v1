// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package main

import (
	pb "cds.ikigai.net/cabinet.v1/rpc"
	"context"
)

func (s *CDSCabinetServer) CounterGet(ctx context.Context, counter *pb.Counter) (*pb.CounterValueResponse, error) {
	//val := pb.CounterValueResponse{Value:1}
	//return &val, nil

	// IRI Reference
	return nil, nil
}
