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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *TransactionOperation) NodeUpdate(node *pb.Node) error{
	nodeId, err := iri.NodeResolveId(node.Id, &o.IdMap)
	if err != nil{
		return status.Errorf(codes.InvalidArgument, RPCErrorIRISpecific, "tmp:X is invalid", "node.id")
	}

	nodeIRI := &iri.Node{
		Type: uint16(node.Type),
		Id: nodeId,
	}

	if vldErr := nodeIRI.ValidateIRI(); vldErr != nil{
		return status.Errorf(codes.InvalidArgument, RPCErrorIRISpecific, vldErr)
	}

	o.tr.Set(nodeIRI.GetKey(o.server.dbNode), PreparePayload(node.Properties))

	if DebugServerRequests {
		o.server.logEvent(fmt.Sprintf("T.NodeUpdate(%v) = %v", o.action, nodeIRI.GetPath()))
	}

	return o.stream.Send(&pb.TransactionActionResponse{
		Status: pb.MutationStatus_SUCCESS,
		ActionId: o.action.ActionId,
	})
}
