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

func (s *CDSCabinetServer) MetaGet(ctx context.Context, meta *pb.Meta) (*pb.MetaGetResponse, error){
	metaValue, err := s.fDb.ReadTransact(func (rtr fdb.ReadTransaction) (ret interface{}, err error) {
		iri, err := resolveMetaIRI(meta, nil)

		if err != nil {
			return nil, &CabinetError{code: CDSErrFieldInvalid}
		}

		metaValue := rtr.Get(iri.getKey(s)).MustGet()

		if metaValue == nil{
			return nil, &CabinetError{code: CDSErrorNotFound}
		}

		return metaValue, nil
	})

	if err != nil{
		return nil, err
	}

	return &pb.MetaGetResponse{Val: metaValue.([]byte)}, nil
}

func (s *CDSCabinetServer) MetaList(metaRq *pb.MetaListRequest, stream pb.CDSCabinet_MetaListServer) error{
	return nil
}
