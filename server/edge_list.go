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

func (s *CDSCabinetServer) EdgeList(edgeRq *pb.EdgeListRequest, stream pb.CDSCabinet_EdgeListServer) error {
	_, err := s.fdb.ReadTransact(func(rtr fdb.ReadTransaction) (interface{}, error) {
		if DebugServerRequests {
			s.logEvent(fmt.Sprintf("EdgeList(%v)", edgeRq))
		}

		listIRI := &iri.Edge{Subject: edgeRq.Subject, Predicate: uint16(edgeRq.Predicate)}
		listOpt := &iri.ListOptions{PageSize: int(edgeRq.Opt.PageSize), Reverse: edgeRq.Opt.Reverse}

		ri := listIRI.GetListRange(s.dbEdge, rtr, listOpt).Iterator()

		for ri.Advance() {
			kv := ri.MustGet()
			edgeKeys, err := s.dbEdge.Unpack(kv.Key) // [subject, predicate, target] = properties

			if err != nil {
				return nil, status.Errorf(codes.DataLoss, RPCErrorDataCorrupted, "node.id")
			}

			obj := &pb.Edge{}

			if edgeRq.IncludeSubject {
				obj.Subject = edgeKeys[0].(string)
			}

			if edgeRq.IncludePredicate {
				predicate, err := iri.SmallKeyToSequence(edgeKeys[1].([]byte))

				if err != nil {
					return nil, status.Errorf(codes.DataLoss, RPCErrorDataCorrupted, "edge.predicate")
				}

				obj.Predicate = uint32(predicate)
			}

			if edgeRq.IncludeTarget {
				obj.Target = edgeKeys[2].(string)
			}

			if edgeRq.IncludeProp {
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
