// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package main

import (
	pb "cds.ikigai.net/cabinet.v1/rpc"
	"context"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
)

func (s *CDSCabinetServer) CounterGet(ctx context.Context, counter *pb.Counter) (*pb.CounterValueResponse, error) {
	total, err := s.fDb.ReadTransact(func (rtr fdb.ReadTransaction) (ret interface{}, err error) {
		iri, err := resolveCounterIRI(counter, nil)
		var total int64

		if err != nil {
			return nil, &CabinetError{code: CDSErrFieldInvalid}
		}

		// simple SUM(IRI/0-f); real-time counting w/ good scalability
		ri := rtr.GetRange(iri.getKeyRange(s), fdb.RangeOptions{
			Limit: 16,
		}).Iterator()

		keysSeen := uint8(0)

		for ri.Advance() {
			kv := ri.MustGet()

			if err != nil {
				return nil, err
			}

			cVal, err := BytesToInt(kv.Value)

			if err != nil {
				return nil, &CabinetError{code: CDSErrorBadRecord, err: "Expected int64, unable to parse"}
			}

			total += cVal
			keysSeen += 1
		}

		if keysSeen == 0{
			return nil, &CabinetError{code: CDSErrorNotFound}
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
