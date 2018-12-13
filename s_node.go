// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package main

import (
	pb "cds.ikigai.net/cabinet.v1/rpc"
	"context"
	"fmt"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
)

func (s *CDSCabinetServer) NodeGet(ctx context.Context, nodeRq *pb.NodeGetRequest) (*pb.Node, error){
	nodeProp, err := s.fDb.ReadTransact(func (rtr fdb.ReadTransaction) (ret interface{}, err error) {
		nodeIRI := &IRINode{Type: uint16(nodeRq.NodeType), Id: nodeRq.Id}
		nodeProp := rtr.Get(nodeIRI.getKey(s)).MustGet()

		if DebugServerRequests {
			s.logEvent(fmt.Sprintf("NodeGet(%v) = %v", nodeRq, nodeIRI.getPath()))
		}

		if nodeProp == nil{
			return nil, &CabinetError{code: CDSErrorNotFound}
		}

		return nodeProp, nil
	})

	if err != nil{
		return nil, err
	}

	return &pb.Node{Properties: nodeProp.([]byte), Id: nodeRq.Id}, nil
}

func (s *CDSCabinetServer) NodeList(nodeRq *pb.NodeListRequest, stream pb.CDSCabinet_NodeListServer) error{
	return nil
}
