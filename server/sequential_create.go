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
	"strconv"
)

func (s *CDSCabinetServer) SequentialCreate(ctx context.Context, seq *pb.Sequential) (newSeq *pb.Sequential, err error){
	vldError := validateSequentialRequest(seq, []string{"t", "u"}, []string{"s"})

	if vldError != nil{
		return nil, vldError
	}

	newId, err := s.fdb.Transact(func (tr fdb.Transaction) (ret interface{}, err error) {
		baseSeqIRI := &iri.Sequence{Type: seq.Type}

		// check UUID is unique
		if tr.Get((&iri.Sequence{Type: seq.Type, UUID: seq.Uuid}).GetReverseKey(s.dbSequence)).MustGet() != nil{
			return nil, status.Error(codes.AlreadyExists, RPCErrorDuplicateRecord)
		}

		lastKey := baseSeqIRI.GetIncrementKey(s.dbSequence)
		lastNum := tr.Get(lastKey).MustGet()

		var lastInt32 uint32

		if lastNum == nil{
			lastInt32 = uint32(1)
		}else {
			lastInt, err := strconv.ParseUint(string(lastNum), 10, 32)

			if err != nil {
				return nil, err
			}

			lastInt32 = uint32(lastInt)
		}

		seqIRI := &iri.Sequence{Type: seq.Type, SeqID: lastInt32, UUID: seq.Uuid}
		rVal, err := Int64ToBytes(int64(lastInt32))
		CheckFatalError(err)

		// store as /s/i/{seqId} = {UUID}; /s/u/{UUID} = {seqId}
		tr.Set(seqIRI.GetKey(s.dbSequence), []byte(seq.Uuid))
		tr.Set(seqIRI.GetReverseKey(s.dbSequence), rVal)

		tr.Set(lastKey, []byte(strconv.FormatUint(uint64(lastInt32 + 1), 10)))

		if DebugServerRequests {
			s.logEvent(fmt.Sprintf("SequentialCreate(%v) = %v", seq, seqIRI.GetPath()))
		}

		return lastInt32, nil
	})

	if err != nil{
		return
	}

	return &pb.Sequential{Seqid: newId.(uint32), Uuid: seq.Uuid}, nil
}
