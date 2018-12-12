// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package main

import (
	pb "cds.ikigai.net/cabinet.v1/rpc"
	"context"
	"fmt"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
	"strconv"
	"strings"
)

func utilSeqMakeKey(n uint32) string{
	return fmt.Sprintf("%05d", n)
	//  strconv.FormatUint(uint64(lastInt32),10)
}

func (s *CDSCabinetServer) SequentialCreate(ctx context.Context, seq *pb.Sequential) (newSeq *pb.Sequential, err error){
	vldError := validateSequentialRequest(seq, []string{"t", "n"}, []string{"s"})

	if vldError != nil{
		return nil, vldError
	}

	newId, err := s.fDb.Transact(func (tr fdb.Transaction) (ret interface{}, err error) {
		var lastKey = s.dbSeq.Pack(tuple.Tuple{"l", seq.GetType()})
		var lastNum = tr.Get(lastKey).MustGet()
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

		tr.Set(s.dbSeq.Pack(tuple.Tuple{seq.GetType(), utilSeqMakeKey(lastInt32)}), []byte(seq.GetNode()))
		tr.Set(lastKey, []byte(strconv.FormatUint(uint64(lastInt32 + 1), 10)))

		return lastInt32, nil
	})

	if err != nil{
		return
	}

	return &pb.Sequential{Seqid: newId.(uint32)}, nil
}

func (s *CDSCabinetServer) SequentialUpdate(ctx context.Context, seq *pb.Sequential) (*pb.MutationResponse, error){
	vldError := validateSequentialRequest(seq, []string{"t", "s", "n"}, []string{})

	if vldError != nil{
		return nil, vldError
	}

	_, err := s.fDb.Transact(func (tr fdb.Transaction) (ret interface{}, err error) {
		var key = s.dbSeq.Pack(tuple.Tuple{seq.GetType(), utilSeqMakeKey(seq.GetSeqid())})
		tr.Set(key, []byte(seq.GetNode()))

		return nil, nil
	})

	if err != nil{
		return &pb.MutationResponse{Status: pb.MutationStatus_PROCESSING_FAILURE}, err
	}

	return &pb.MutationResponse{Status: pb.MutationStatus_SUCCESS}, nil
}

func (s *CDSCabinetServer) SequentialDelete(ctx context.Context, seq *pb.Sequential) (*pb.MutationResponse, error){
	vldError := validateSequentialRequest(seq, []string{"t", "s"}, []string{"n"})

	if vldError != nil{
		return nil, vldError
	}

	_, err := s.fDb.Transact(func (tr fdb.Transaction) (ret interface{}, err error) {
		var key = s.dbSeq.Pack(tuple.Tuple{seq.GetType(), utilSeqMakeKey(seq.GetSeqid())})
		tr.Clear(key)

		return nil, nil
	})

	if err != nil{
		return &pb.MutationResponse{Status: pb.MutationStatus_PROCESSING_FAILURE}, err
	}

	return &pb.MutationResponse{Status: pb.MutationStatus_SUCCESS}, nil
}

func (s *CDSCabinetServer) SequentialGet(ctx context.Context, seq *pb.Sequential) (*pb.Sequential, error){
	vldError := validateSequentialRequest(seq, []string{"t", "s"}, []string{"n"})

	if vldError != nil{
		return nil, vldError
	}

	nodeId, err := s.fDb.ReadTransact(func (rtr fdb.ReadTransaction) (ret interface{}, err error) {
		var key = s.dbSeq.Pack(tuple.Tuple{seq.GetType(), utilSeqMakeKey(seq.GetSeqid())})
		sVal := rtr.Get(key).MustGet()

		if sVal == nil{
			return nil, &CabinetError{code: CDSErrorNotFound}
		}

		return string(sVal), nil
	})

	if err != nil{
		return nil, err
	}

	return &pb.Sequential{Node: nodeId.(string)}, nil
}

func (s *CDSCabinetServer) SequentialList(seq *pb.SequentialListRequest, stream pb.CDSCabinet_SequentialListServer) error{
	if len(seq.GetType()) == 0 {
		return &CabinetError{code: CDSErrorFieldRequired, field: "type"}
	}else if seq.Opt.GetPageSize() == 0{
		return &CabinetError{code: CDSListNoPagination, field: "opt.page_size"}
	}

	_, err := s.fDb.ReadTransact(func (rtr fdb.ReadTransaction) (interface{}, error) {
		var readRange = s.dbSeq.Sub(seq.GetType())

		ri := rtr.GetRange(readRange, fdb.RangeOptions{
			Limit: int(seq.Opt.PageSize),
		}).Iterator()

		for ri.Advance() {
			kv := ri.MustGet()
			sqO, err := s.dbSeq.Unpack(kv.Key) // {Type, SeqID} = kv.Value

			if err != nil {
				return nil, err
			}

			seqID, err := strconv.ParseUint(strings.TrimLeft(sqO[1].(string), "0"), 10, 32)

			if err != nil {
				return nil, err
			}

			obj := &pb.Sequential{
				Type: seq.Type,
				Node: string(kv.Value),
				Seqid: uint32(seqID),
			}

			if err := stream.Send(obj); err != nil {
				return nil, err
			}
		}

		return nil, nil
	})

	return err
}

func validateSequentialRequest(seq *pb.Sequential, required []string, unexpected []string) error{
	for i := range required {
		if required[i] == "t" && len(seq.GetType()) == 0 {
			return &CabinetError{code: CDSErrorFieldRequired, field: "type"}
		}else if required[i] == "n" && len(seq.GetNode()) == 0 {
			return &CabinetError{code: CDSErrorFieldRequired, field: "node"}
		}else if required[i] == "s" && seq.GetSeqid() == 0 {
			return &CabinetError{code: CDSErrorFieldRequired, field: "seqId"}
		}
	}

	for i := range unexpected {
		if unexpected[i] == "t" && len(seq.GetType()) > 0 {
			return &CabinetError{code: CDSErrorFieldUnexpected, field: "type"}
		}else if unexpected[i] == "n" && len(seq.GetNode()) > 0 {
			return &CabinetError{code: CDSErrorFieldUnexpected, field: "node"}
		}else if unexpected[i] == "s" && seq.GetSeqid() != 0 {
			return &CabinetError{code: CDSErrorFieldUnexpected, field: "seqId"}
		}
	}

	return nil
}
