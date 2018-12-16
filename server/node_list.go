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
)

func (s *CDSCabinetServer) NodeList(nodeRq *pb.NodeListRequest, stream pb.CDSCabinet_NodeListServer) error{
	_, err := s.fdb.ReadTransact(func (rtr fdb.ReadTransaction) (interface{}, error) {
		var readRange = s.dbNode.Sub(nodeRq.NodeType)

		if DebugServerRequests {
			s.logEvent(fmt.Sprintf("NodeList(%v)", nodeRq))
		}

		ri := rtr.GetRange(readRange, fdb.RangeOptions{
			Limit: 		int(nodeRq.Opt.PageSize),
			Reverse: 	nodeRq.Opt.Reverse,
		}).Iterator()

		for ri.Advance() {
			kv := ri.MustGet()
			nodeKeys, err := s.dbSequence.Unpack(kv.Key) // [node_type, node_id] = properties

			if err != nil {
				return nil, status.Errorf(codes.DataLoss, RPCErrorDataCorrupted, "node.id")
			}

			obj := &pb.Node{}

			if nodeRq.IncludeId {
				obj.Id = nodeKeys[1].(string)
			}

			if nodeRq.IncludeType {
				obj.Type = nodeRq.NodeType
			}

			if nodeRq.IncludeProp{
				obj.Properties = kv.Value
			}

			if err := stream.Send(obj); err != nil {
				return nil, err
			}
		}

		return nil, nil
	})

	return err
}
