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


func (s *CDSCabinetServer) IndexList(indexRq *pb.IndexListRequest, stream pb.CDSCabinet_IndexListServer) error{
	_, err := s.fdb.ReadTransact(func (rtr fdb.ReadTransaction) (interface{}, error) {
		if DebugServerRequests {
			s.logEvent(fmt.Sprintf("IndexList(%v)", indexRq))
		}

		listIRI := &iri.NodeIndex{IndexId: uint16(indexRq.Index), Value: indexRq.Value}
		ri := listIRI.GetListRange(s.dbIndex, rtr, indexRq.Opt).Iterator()

		for ri.Advance() {
			kv := ri.MustGet()
			indexKeys, err := s.dbIndex.Unpack(kv.Key) // [type, value, node] = properties

			if err != nil {
				return nil, status.Errorf(codes.DataLoss, RPCErrorDataCorrupted, "node_key")
			}

			obj := &pb.Index{}

			if indexRq.IncludeIndex {
				obj.Type = indexRq.Index
			}

			if indexRq.IncludeValue {
				obj.Value = indexKeys[1].(string)
			}

			if indexRq.IncludeProp{
				obj.Properties = kv.Value
			}

			if indexRq.IncludeNode{
				obj.Node = indexKeys[2].(string)
			}

			if err := stream.Send(obj); err != nil {
				return nil, err
			}
		}

		return nil, nil
	})

	return err
}
