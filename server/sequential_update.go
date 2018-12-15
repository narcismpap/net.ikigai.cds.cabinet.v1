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

func (s *CDSCabinetServer) SequentialUpdate(ctx context.Context, seq *pb.Sequential) (*pb.MutationResponse, error){
	vldError := validateSequentialRequest(seq, []string{"t", "s", "n"}, []string{})

	if vldError != nil{
		return nil, vldError
	}

	_, err := s.fdb.Transact(func (tr fdb.Transaction) (ret interface{}, err error) {
		seqIRI := iri.Sequence{Type: seq.Type, SeqID: seq.Seqid}
		tr.Set(seqIRI.GetKey(s.dbSequence), []byte(seq.GetNode()))

		if DebugServerRequests {
			s.logEvent(fmt.Sprintf("SequentialUpdate(%v) = %v", seq, seqIRI.GetPath()))
		}

		return nil, nil
	})

	if err != nil{
		return &pb.MutationResponse{Status: pb.MutationStatus_PROCESSING_FAILURE}, err
	}

	return &pb.MutationResponse{Status: pb.MutationStatus_SUCCESS}, nil
}