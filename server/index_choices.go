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
)

func (s *CDSCabinetServer) IndexChoices(indexRq *pb.IndexChoiceRequest, stream pb.CDSCabinet_IndexChoicesServer) error {
	_, err := s.fdb.ReadTransact(func(rtr fdb.ReadTransaction) (interface{}, error) {
		if DebugServerRequests {
			s.logEvent(fmt.Sprintf("IndexChoices(%v)", indexRq))
		}

		listIRI := &iri.NodeIndex{IndexId: uint16(indexRq.Index), Value: "*"}
		listOpt := &iri.ListOptions{PageSize: int(indexRq.Opt.PageSize), Reverse: indexRq.Opt.Reverse}

		ri := listIRI.GetCounterListRange(s.dbIndexCnt, rtr, listOpt).Iterator()

		for ri.Advance() {
			kv := ri.MustGet()
			idxChoiceKeys, err := s.dbIndexCnt.Unpack(kv.Key) // [type, value] = atomic<cnt>

			if err != nil {
				return nil, status.Errorf(codes.DataLoss, RPCErrorDataCorrupted, "index.key")
			}

			cVal, err := BytesToInt(kv.Value)

			if err != nil {
				return nil, status.Error(codes.DataLoss, fmt.Sprintf(RPCErrorDataCorrupted, "index.choice.key"))
			}

			// sanity checks, when counter drops <0 means there is an application bug (T.IndexDelete > T.IndexCreate)
			if cVal < 0 {
				s.logError(NewKeyReport(KeyReportBelowZero, indexRq, "val <0", []byte(kv.Key), kv.Value))
				cVal = 0
			}

			obj := &pb.IndexChoice{
				Value: idxChoiceKeys[1].(string),
				Count: uint32(cVal),
			}

			if err := stream.Send(obj); err != nil {
				return nil, err
			}
		}

		return nil, nil
	})

	return err
}
