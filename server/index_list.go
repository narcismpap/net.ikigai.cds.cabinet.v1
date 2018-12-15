// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package server

import (
	pb "cds.ikigai.net/cabinet.v1/rpc"
)


func (s *CDSCabinetServer) IndexList(indexRq *pb.IndexListRequest, stream pb.CDSCabinet_IndexListServer) error{
	return nil
}
