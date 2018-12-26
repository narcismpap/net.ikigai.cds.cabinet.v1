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

func (s *CDSCabinetServer) NodeList(nodeRq *pb.NodeListRequest, stream pb.CDSCabinet_NodeListServer) error {
	_, err := s.fdb.ReadTransact(func(rtr fdb.ReadTransaction) (interface{}, error) {
		if DebugServerRequests {
			s.logEvent(fmt.Sprintf("NodeList(%v)", nodeRq))
		}

		listIRI := &iri.Node{Type: uint16(nodeRq.NodeType)}
		ri := listIRI.GetListRange(s.dbNode, rtr, nodeRq.Opt).Iterator()

		for ri.Advance() {
			kv := ri.MustGet()
			nodeKeys, err := s.dbNode.Unpack(kv.Key) // [node_type, node_id] = properties

			if err != nil {
				return nil, status.Errorf(codes.DataLoss, RPCErrorDataCorrupted, "node.id")
			}

			obj := &pb.Node{}

			if nodeRq.IncludeId {
				obj.Id = nodeKeys[1].(string)
			}

			if nodeRq.IncludeType {
				nType, err := iri.SmallKeyToSequence(nodeKeys[0].([]byte))

				if err != nil {
					return nil, status.Errorf(codes.DataLoss, RPCErrorDataCorrupted, "node.type")
				}

				obj.Type = uint32(nType)
			}

			if nodeRq.IncludeProp {
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
