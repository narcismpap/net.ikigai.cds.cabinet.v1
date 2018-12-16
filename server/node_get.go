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

func (s *CDSCabinetServer) NodeGet(ctx context.Context, nodeRq *pb.NodeGetRequest) (*pb.Node, error){
	nodeProp, err := s.fdb.ReadTransact(func (rtr fdb.ReadTransaction) (ret interface{}, err error) {
		nodeIRI := &iri.Node{Type: uint16(nodeRq.NodeType), Id: nodeRq.Id}
		nodePerms := &perms.Node{}

		if vldErr := nodeIRI.ValidateIRI(nodePerms); vldErr != nil{
			return nil, status.Errorf(codes.InvalidArgument, RPCErrorIRISpecific, vldErr)
		}

		nodeProp := rtr.Get(nodeIRI.GetKey(s.dbNode)).MustGet()

		if DebugServerRequests {
			s.logEvent(fmt.Sprintf("NodeGet(%v) = %v", nodeRq, nodeIRI.GetPath()))
		}

		if nodeProp == nil{
			return nil, status.Error(codes.NotFound, RPCErrorNotFound)
		}

		return nodeProp, nil
	})

	if err != nil{
		return nil, err
	}

	return &pb.Node{Properties: nodeProp.([]byte), Id: nodeRq.Id}, nil
}
