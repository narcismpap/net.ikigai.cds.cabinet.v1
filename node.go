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

func (s *CDSCabinetServer) NodeGet(ctx context.Context, noteRq *pb.NodeGetRequest) (*pb.Node, error){
	return nil, nil
}

func (s *CDSCabinetServer) NodeList(nodeRq *pb.NodeListRequest, stream pb.CDSCabinet_NodeListServer) error{
	return nil
}
