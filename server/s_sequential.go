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
	"strings"
)

func (s *CDSCabinetServer) SequentialCreate(ctx context.Context, seq *pb.Sequential) (newSeq *pb.Sequential, err error){
	vldError := validateSequentialRequest(seq, []string{"t", "n"}, []string{"s"})

	if vldError != nil{
		return nil, vldError
	}

	newId, err := s.FdbConn.Transact(func (tr fdb.Transaction) (ret interface{}, err error) {
		baseSeqIRI := &iri.IRISequential{Type: seq.Type}
		lastKey := baseSeqIRI.GetIncrementKey(s)
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

		seqIRI := &iri.IRISequential{Type: seq.Type, SeqID: lastInt32}

		tr.Set(seqIRI.GetKey(s), []byte(seq.GetNode()))
		tr.Set(lastKey, []byte(strconv.FormatUint(uint64(lastInt32 + 1), 10)))

		if DebugServerRequests {
			s.logEvent(fmt.Sprintf("SequentialCreate(%v) = %v", seq, seqIRI.GetPath()))
		}

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

	_, err := s.FdbConn.Transact(func (tr fdb.Transaction) (ret interface{}, err error) {
		seqIRI := iri.IRISequential{Type: seq.Type, SeqID: seq.Seqid}
		tr.Set(seqIRI.GetKey(s), []byte(seq.GetNode()))

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

func (s *CDSCabinetServer) SequentialDelete(ctx context.Context, seq *pb.Sequential) (*pb.MutationResponse, error){
	vldError := validateSequentialRequest(seq, []string{"t", "s"}, []string{"n"})

	if vldError != nil{
		return nil, vldError
	}

	_, err := s.FdbConn.Transact(func (tr fdb.Transaction) (ret interface{}, err error) {
		seqIRI := iri.IRISequential{Type: seq.Type, SeqID: seq.Seqid}

		tr.Clear(seqIRI.GetKey(s))
		tr.Clear(seqIRI.GetIncrementKey(s))

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

func (s *CDSCabinetServer) SequentialGet(ctx context.Context, seq *pb.Sequential) (*pb.Sequential, error){
	vldError := validateSequentialRequest(seq, []string{"t", "s"}, []string{"n"})

	if vldError != nil{
		return nil, vldError
	}

	nodeId, err := s.FdbConn.ReadTransact(func (rtr fdb.ReadTransaction) (ret interface{}, err error) {
		seqIRI := iri.IRISequential{Type: seq.Type, SeqID: seq.Seqid}
		sVal := rtr.Get(seqIRI.GetKey(s)).MustGet()

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

func (s *CDSCabinetServer) SequentialList(seq *pb.SequentialListRequest, stream pb.CDSCabinet_SequentialListServer) error{
	if len(seq.GetType()) == 0 {
		return status.Error(codes.InvalidArgument, fmt.Sprintf(RPCErrorFieldRequired, ".type"))
	}else if seq.Opt.GetPageSize() == 0{
		return status.Error(codes.InvalidArgument, fmt.Sprintf(RPCErrorFieldRequired, "opt.page_size"))
	}

	_, err := s.FdbConn.ReadTransact(func (rtr fdb.ReadTransaction) (interface{}, error) {
		var readRange = s.DbSeq.Sub(seq.GetType())

		if DebugServerRequests {
			s.logEvent(fmt.Sprintf("SequentialList(%v)", seq))
		}

		ri := rtr.GetRange(readRange, fdb.RangeOptions{
			Limit: int(seq.Opt.PageSize),
		}).Iterator()

		for ri.Advance() {
			kv := ri.MustGet()
			sqO, err := s.DbSeq.Unpack(kv.Key) // {Type, SeqID} = kv.Value

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
			return status.Error(codes.InvalidArgument, fmt.Sprintf(RPCErrorFieldRequired, "seq.type"))
		}else if required[i] == "n" && len(seq.GetNode()) == 0 {
			return status.Error(codes.InvalidArgument, fmt.Sprintf(RPCErrorFieldRequired, "seq.node"))
		}else if required[i] == "s" && seq.GetSeqid() == 0 {
			return status.Error(codes.InvalidArgument, fmt.Sprintf(RPCErrorFieldRequired, "seq.seqId"))
		}
	}

	for i := range unexpected {
		if unexpected[i] == "t" && len(seq.GetType()) > 0 {
			return status.Error(codes.InvalidArgument, fmt.Sprintf(RPCErrorFieldUnexpected, "seq.type"))
		}else if unexpected[i] == "n" && len(seq.GetNode()) > 0 {
			return status.Error(codes.InvalidArgument, fmt.Sprintf(RPCErrorFieldUnexpected, "seq.node"))
		}else if unexpected[i] == "s" && seq.GetSeqid() != 0 {
			return status.Error(codes.InvalidArgument, fmt.Sprintf(RPCErrorFieldUnexpected, "seq.seqId"))
		}
	}

	return nil
}
