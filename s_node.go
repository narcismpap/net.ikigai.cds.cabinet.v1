// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package main

import (
	pb "cds.ikigai.net/cabinet.v1/rpc"
	"context"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
)

func (s *CDSCabinetServer) NodeGet(ctx context.Context, noteRq *pb.NodeGetRequest) (*pb.Node, error){
	nodeProp, err := s.fDb.ReadTransact(func (rtr fdb.ReadTransaction) (ret interface{}, err error) {
		nodeProp := rtr.Get((&IRINode{
			Type: uint16(noteRq.NodeType),
			Id: noteRq.Id,
		}).getKey(s)).MustGet()

		if nodeProp == nil{
			return nil, &CabinetError{code: CDSErrorNotFound}
		}

		return nodeProp, nil
	})

	if err != nil{
		return nil, err
	}

	return &pb.Node{Properties: nodeProp.([]byte)}, nil
}

func (s *CDSCabinetServer) NodeList(nodeRq *pb.NodeListRequest, stream pb.CDSCabinet_NodeListServer) error{
	return nil
}
