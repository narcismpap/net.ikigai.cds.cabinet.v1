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
	"fmt"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *CDSCabinetServer) IndexDrop(ctx context.Context, index *pb.IndexDropRequest) (*pb.MutationResponse, error) {
	indexIRI := &iri.NodeIndex{
		IndexId: uint16(index.Index),
	}

	_, err := s.fdb.Transact(func(tr fdb.Transaction) (ret interface{}, err error) {
		idxPerms := &perms.Index{AllowNodeWildcard: true, AllowValueWildcard: true}

		if vldErr := indexIRI.ValidateIRI(idxPerms); vldErr != nil {
			return nil, status.Errorf(codes.InvalidArgument, RPCErrorIRISpecific, vldErr)
		}

		tr.ClearRange(indexIRI.GetClearRange(s.dbIndex))
		tr.ClearRange(indexIRI.GetCounterClearRange(s.dbIndexCnt))

		if DebugServerRequests {
			s.logEvent(fmt.Sprintf("IndexDrop(%v)", index))
		}

		return nil, nil
	})

	if err != nil {
		return nil, err
	}

	return &pb.MutationResponse{Status: pb.MutationStatus_SUCCESS}, nil
}
