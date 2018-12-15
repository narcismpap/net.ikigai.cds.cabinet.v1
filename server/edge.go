// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package server

import (
	"cds.ikigai.net/cabinet.v1/iri"
	pb "cds.ikigai.net/cabinet.v1/rpc"
	"context"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *CDSCabinetServer) EdgeGet(ctx context.Context, edge *pb.EdgeGetRequest) (*pb.Edge, error){
	edgeProp, err := s.FdbConn.ReadTransact(func (rtr fdb.ReadTransaction) (ret interface{}, err error) {
		edgeProp := rtr.Get((&iri.Edge{
			Subject: edge.Edge.Subject,
			Predicate: uint16(edge.Edge.Predicate),
			Target: edge.Edge.Target,
		}).GetKey(s.DbEdge)).MustGet()

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
