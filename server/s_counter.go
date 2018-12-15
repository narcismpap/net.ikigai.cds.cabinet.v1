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

func (s *CDSCabinetServer) CounterGet(ctx context.Context, counter *pb.Counter) (*pb.CounterValueResponse, error) {
	total, err := s.FdbConn.ReadTransact(func (rtr fdb.ReadTransaction) (ret interface{}, err error) {
		iri, err := iri.ResolveCounterIRI(counter, nil)
		var total int64

		if err != nil {
			return nil, status.Error(codes.InvalidArgument, RPCErrorInvalidIRI)
		}

		// simple SUM(IRI/0-f); real-time counting w/ good scalability
		ri := rtr.GetRange(iri.GetKeyRange(s), fdb.RangeOptions{
			Limit: len(CounterKeys),
		}).Iterator()

		keysSeen := uint8(0)

		for ri.Advance() {
			kv := ri.MustGet()

			if err != nil {
				return nil, err
			}

			cVal, err := BytesToInt(kv.Value)

			if err != nil {
				return nil, status.Error(codes.DataLoss, fmt.Sprintf(RPCErrorDataCorrupted, "counter.slice"))
			}

			total += cVal
			keysSeen += 1
		}

		if keysSeen == 0{
			return nil, status.Error(codes.NotFound, RPCErrorNotFound)
		}

		return total, nil
	})

	if err != nil{
		return nil, err
	}

	return &pb.CounterValueResponse{Value: total.(int64)}, nil
}


var CounterKeys = map[int]string{
	0: "0",
	1: "1",
	2: "2",
	3: "3",
	4: "4",
	5: "5",
	6: "6",
	7: "7",
	8: "8",
	9: "9",
	10: "a",
	11: "b",
	12: "c",
	13: "d",
	14: "e",
	15: "f",
}
