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

func (s *CDSCabinetServer) IndexGet(ctx context.Context, indexGet *pb.IndexGetRequest) (*pb.Index, error){
	indexProp, err := s.FdbConn.ReadTransact(func (rtr fdb.ReadTransaction) (ret interface{}, err error) {
		indexProp := rtr.Get((&iri.NodeIndex{
			Node: indexGet.Index.Node,
			IndexId: uint16(indexGet.Index.Type),
			Value: indexGet.Index.Value,
		}).GetKey(s.DbIndex)).MustGet()

		if indexProp == nil{
			return nil, status.Error(codes.NotFound, RPCErrorNotFound)
		}

		return indexProp, nil
	})

	if err != nil{
		return nil, err
	}

	return &pb.Index{Properties: indexProp.([]byte)}, nil
}


func (s *CDSCabinetServer) IndexList(indexRq *pb.IndexListRequest, stream pb.CDSCabinet_IndexListServer) error{
	return nil
}
