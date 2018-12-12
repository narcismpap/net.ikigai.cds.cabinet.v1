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

func (s *CDSCabinetServer) MetaGet(ctx context.Context, meta *pb.Meta) (*pb.MetaGetResponse, error){
	return nil, nil
}

func (s *CDSCabinetServer) MetaList(metaRq *pb.MetaListRequest, stream pb.CDSCabinet_MetaListServer) error{
	return nil
}