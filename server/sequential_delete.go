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

func (s *CDSCabinetServer) SequentialDelete(ctx context.Context, seq *pb.Sequential) (*pb.MutationResponse, error){
	vldError := validateSequentialRequest(seq, []string{"t", "u"}, []string{"s"})

	if vldError != nil{
		return nil, vldError
	}

	_, err := s.fdb.Transact(func (tr fdb.Transaction) (ret interface{}, err error) {
		seqIRI := iri.Sequence{Type: seq.Type, UUID: seq.Uuid}

		if vldErr := seqIRI.ValidateIRI(); vldErr != nil{
			return nil, status.Errorf(codes.InvalidArgument, RPCErrorIRISpecific, vldErr)
		}

		dbSeqID := tr.Get(seqIRI.GetReverseKey(s.dbSequence)).MustGet()

		if dbSeqID == nil{
			return nil, status.Error(codes.NotFound, RPCErrorNotFound)
		}

		seqID, err := BytesToInt(dbSeqID)

		if err != nil{
			return nil, status.Error(codes.DataLoss, fmt.Sprintf(RPCErrorDataCorrupted, "seqId"))
		}

		seqIRI.SeqID = uint32(seqID)

		tr.Clear(seqIRI.GetKey(s.dbSequence))
		tr.Clear(seqIRI.GetReverseKey(s.dbSequence))

		if DebugServerRequests {
			s.logEvent(fmt.Sprintf("SequentialDelete(%v) = %v", seq, seqIRI.GetPath()))
		}

		return nil, nil
	})

	if err != nil{
		return nil, err
	}

	return &pb.MutationResponse{Status: pb.MutationStatus_SUCCESS}, nil
}
