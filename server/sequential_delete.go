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
)

func (s *CDSCabinetServer) SequentialDelete(ctx context.Context, seq *pb.Sequential) (*pb.MutationResponse, error){
	vldError := validateSequentialRequest(seq, []string{"t", "s"}, []string{"n"})

	if vldError != nil{
		return nil, vldError
	}

	_, err := s.fdb.Transact(func (tr fdb.Transaction) (ret interface{}, err error) {
		seqIRI := iri.Sequence{Type: seq.Type, SeqID: seq.Seqid}

		tr.Clear(seqIRI.GetKey(s.dbSequence))
		tr.Clear(seqIRI.GetIncrementKey(s.dbSequence))

		if DebugServerRequests {
			s.logEvent(fmt.Sprintf("SequentialDelete(%v) = %v", seq, seqIRI.GetPath()))
		}

		return nil, nil
	})

	if err != nil{
		return &pb.MutationResponse{Status: pb.MutationStatus_PROCESSING_FAILURE}, err
	}

	return &pb.MutationResponse{Status: pb.MutationStatus_SUCCESS}, nil
}
