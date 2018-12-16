// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package server

import (
	"cds.ikigai.net/cabinet.v1/iri"
	"cds.ikigai.net/cabinet.v1/perms"
	pb "cds.ikigai.net/cabinet.v1/rpc"
	"context"
	"fmt"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *CDSCabinetServer) CounterGet(ctx context.Context, counter *pb.Counter) (*pb.CounterValueResponse, error) {
	total, err := s.fdb.ReadTransact(func (rtr fdb.ReadTransaction) (ret interface{}, err error) {
		counterPerms := &perms.Count{}

		cntIRI, err := iri.ResolveCounterIRI(counter, nil, counterPerms)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, RPCErrorIRISpecific, err)
		}

		// simple SUM(IRI/0-f); real-time counting w/ good scalability
		var total int64

		ri := rtr.GetRange(cntIRI.GetKeyRange(s.dbCount), fdb.RangeOptions{
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
