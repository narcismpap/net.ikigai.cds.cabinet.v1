// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package server

import (
	pb "cds.ikigai.net/cabinet.v1/rpc"
)


func (s *CDSCabinetServer) EdgeList(edgeRq *pb.EdgeListRequest, stream pb.CDSCabinet_EdgeListServer) error{
	return nil
}


