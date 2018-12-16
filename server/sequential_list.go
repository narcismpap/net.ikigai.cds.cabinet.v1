// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package server

import (
	"cds.ikigai.net/cabinet.v1/iri"
	pb "cds.ikigai.net/cabinet.v1/rpc"
	"fmt"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"strings"
)

func (s *CDSCabinetServer) SequentialList(seqRq *pb.SequentialListRequest, stream pb.CDSCabinet_SequentialListServer) error{
	if len(seqRq.GetType()) == 0 {
		return status.Error(codes.InvalidArgument, fmt.Sprintf(RPCErrorFieldRequired, ".type"))
	}else if seqRq.Opt.GetPageSize() == 0{
		return status.Error(codes.InvalidArgument, fmt.Sprintf(RPCErrorFieldRequired, "opt.page_size"))
	}

	_, err := s.fdb.ReadTransact(func (rtr fdb.ReadTransaction) (interface{}, error) {
		if DebugServerRequests {
			s.logEvent(fmt.Sprintf("SequentialList(%v)", seqRq))
		}

		listIRI := &iri.Sequence{Type: seqRq.Type}
		ri := listIRI.GetListRange(s.dbSequence, rtr, seqRq.Opt).Iterator()

		for ri.Advance() {
			kv := ri.MustGet()
			sqO, err := s.dbSequence.Unpack(kv.Key) // {Type, i/u, SeqID} = kv.Value

			if err != nil {
				return nil, status.Error(codes.Unavailable, RPCErrorListIterator)
			}

			seqID, err := strconv.ParseUint(strings.TrimLeft(sqO[2].(string), "0"), 10, 32)

			if err != nil {
				return nil, err
			}

			obj := &pb.Sequential{}

			if seqRq.IncludeType{
				obj.Type = seqRq.Type
			}

			if seqRq.IncludeSeqid{
				obj.Seqid = uint32(seqID)
			}

			if seqRq.IncludeUuid{
				obj.Uuid = string(kv.Value)
			}

			if err := stream.Send(obj); err != nil {
				return nil, err
			}
		}

		return nil, nil
	})

	return err
}
