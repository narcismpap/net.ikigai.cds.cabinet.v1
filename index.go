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

func (s *CDSCabinetServer) IndexGet(ctx context.Context, indexGet *pb.IndexGetRequest) (*pb.Index, error){
	indexProp, err := s.fDb.ReadTransact(func (rtr fdb.ReadTransaction) (ret interface{}, err error) {
		indexProp := rtr.Get((&IRINodeIndex{
			Node: indexGet.Index.Subject,
			IndexId: uint16(indexGet.Index.Type),
			Value: indexGet.Index.Value,
		}).getKey(s)).MustGet()

		if indexProp == nil{
			return nil, &CabinetError{code: CDSErrorNotFound}
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
