// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package server

import (
	"cds.ikigai.net/cabinet.v1/iri"
	"cds.ikigai.net/cabinet.v1/perms"
	pb "cds.ikigai.net/cabinet.v1/rpc"
	"context"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *CDSCabinetServer) IndexGet(ctx context.Context, indexGet *pb.IndexGetRequest) (*pb.Index, error){
	indexProp, err := s.fdb.ReadTransact(func (rtr fdb.ReadTransaction) (ret interface{}, err error) {
		indexIRI := &iri.NodeIndex{
			Node: indexGet.Index.Node,
			IndexId: uint16(indexGet.Index.Type),
			Value: indexGet.Index.Value,
		}

		indexPerms := &perms.Index{}

		if vldErr := indexIRI.ValidateIRI(indexPerms); vldErr != nil{
			return nil, status.Errorf(codes.InvalidArgument, RPCErrorIRISpecific, vldErr)
		}

		indexProp := rtr.Get(indexIRI.GetKey(s.dbIndex)).MustGet()

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

