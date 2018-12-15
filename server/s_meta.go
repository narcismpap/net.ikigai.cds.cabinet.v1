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

func (s *CDSCabinetServer) MetaGet(ctx context.Context, meta *pb.Meta) (*pb.MetaGetResponse, error){
	metaValue, err := s.FdbConn.ReadTransact(func (rtr fdb.ReadTransaction) (ret interface{}, err error) {
		iri, err := iri.ResolveMetaIRI(meta, nil)

		if err != nil {
			return nil, status.Error(codes.InvalidArgument, RPCErrorInvalidIRI)
		}

		metaValue := rtr.Get(iri.GetKey(s)).MustGet()

		if metaValue == nil{
			return nil, status.Error(codes.NotFound, RPCErrorNotFound)
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
