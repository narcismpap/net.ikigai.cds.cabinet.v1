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
	vldError := validateSequentialRequest(seq, []string{"t"}, []string{})

	if vldError != nil{
		return nil, vldError
	}

	if seq.Seqid > 0 && len(seq.Uuid) > 0{
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf(RPCErrorFieldSpecific, "mutually exclusive", "seq.uuid,seq.seqId"))
	}

	if seq.Seqid == 0 && len(seq.Uuid) == 0{
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf(RPCErrorFieldSpecific, "one ID required", "seq.uuid|seq.seqId"))
	}

	seqResponse, err := s.fdb.ReadTransact(func (rtr fdb.ReadTransaction) (ret interface{}, err error) {
		seqIRI := iri.Sequence{Type: seq.Type, SeqID: seq.Seqid, UUID: seq.Uuid}

		if vldErr := seqIRI.ValidateIRI(); vldErr != nil{
			return nil, status.Errorf(codes.InvalidArgument, RPCErrorIRISpecific, vldErr)
		}

		if DebugServerRequests {
			s.logEvent(fmt.Sprintf("SequentialGet(%v) = %v", seq, seqIRI.GetPath()))
		}

		if len(seq.Uuid) == 0 {
			sVal := rtr.Get(seqIRI.GetKey(s.dbSequence)).MustGet()

			if sVal == nil{
				return nil, status.Error(codes.NotFound, RPCErrorNotFound)
			}

			seqIRI.UUID = string(sVal)
		}else{
			sVal := rtr.Get(seqIRI.GetReverseKey(s.dbSequence)).MustGet()

			if sVal == nil{
				return nil, status.Error(codes.NotFound, RPCErrorNotFound)
			}

			seqID, err := BytesToInt(sVal)

			if err != nil{
				return nil, status.Error(codes.DataLoss, fmt.Sprintf(RPCErrorDataCorrupted, "seqId"))
			}

			seqIRI.SeqID = uint32(seqID)
		}

		return &pb.Sequential{Seqid: seqIRI.SeqID, Uuid: seqIRI.UUID}, nil
	})

	if err != nil{
		return nil, err
	}

	return seqResponse.(*pb.Sequential), nil
}
