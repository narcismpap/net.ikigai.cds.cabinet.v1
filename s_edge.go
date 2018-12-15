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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *CDSCabinetServer) EdgeGet(ctx context.Context, edge *pb.EdgeGetRequest) (*pb.Edge, error){
	edgeProp, err := s.fDb.ReadTransact(func (rtr fdb.ReadTransaction) (ret interface{}, err error) {
		edgeProp := rtr.Get((&IRIEdge{
			Subject: edge.Edge.Subject,
			Predicate: uint16(edge.Edge.Predicate),
			Target: edge.Edge.Target,
		}).getKey(s)).MustGet()

		if edgeProp == nil{
			return nil, status.Error(codes.NotFound, RPCErrorNotFound)
		}

		return edgeProp, nil
	})

	if err != nil{
		return nil, err
	}

	return &pb.Edge{Properties: edgeProp.([]byte)}, nil
}

func (s *CDSCabinetServer) EdgeList(edgeRq *pb.EdgeListRequest, stream pb.CDSCabinet_EdgeListServer) error{
	return nil
}




