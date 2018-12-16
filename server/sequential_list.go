// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package server

import (
	pb "cds.ikigai.net/cabinet.v1/rpc"
	"fmt"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"strings"
)

func (s *CDSCabinetServer) SequentialList(seq *pb.SequentialListRequest, stream pb.CDSCabinet_SequentialListServer) error{
	if len(seq.GetType()) == 0 {
		return status.Error(codes.InvalidArgument, fmt.Sprintf(RPCErrorFieldRequired, ".type"))
	}else if seq.Opt.GetPageSize() == 0{
		return status.Error(codes.InvalidArgument, fmt.Sprintf(RPCErrorFieldRequired, "opt.page_size"))
	}

	_, err := s.fdb.ReadTransact(func (rtr fdb.ReadTransaction) (interface{}, error) {
		var readRange = s.dbSequence.Sub(seq.GetType()).Sub("i")

		if DebugServerRequests {
			s.logEvent(fmt.Sprintf("SequentialList(%v)", seq))
		}

		ri := rtr.GetRange(readRange, fdb.RangeOptions{
			Limit: int(seq.Opt.PageSize),
		}).Iterator()

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

			obj := &pb.Sequential{
				Uuid: string(kv.Value),
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
