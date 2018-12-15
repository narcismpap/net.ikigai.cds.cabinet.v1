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
	"fmt"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *CDSCabinetServer) SequentialGet(ctx context.Context, seq *pb.Sequential) (*pb.Sequential, error){
	vldError := validateSequentialRequest(seq, []string{"t", "s"}, []string{"n"})

	if vldError != nil{
		return nil, vldError
	}

	nodeId, err := s.fdb.ReadTransact(func (rtr fdb.ReadTransaction) (ret interface{}, err error) {
		seqIRI := iri.Sequence{Type: seq.Type, SeqID: seq.Seqid}
		sVal := rtr.Get(seqIRI.GetKey(s.dbSequence)).MustGet()

		if DebugServerRequests {
			s.logEvent(fmt.Sprintf("SequentialGet(%v) = %v", seq, seqIRI.GetPath()))
		}

		if sVal == nil{
			return nil, status.Error(codes.NotFound, RPCErrorNotFound)
		}

		return string(sVal), nil
	})

	if err != nil{
		return nil, err
	}

	return &pb.Sequential{Node: nodeId.(string)}, nil
}
