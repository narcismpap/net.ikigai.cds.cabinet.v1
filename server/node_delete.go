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
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *TransactionOperation) NodeDelete(node *pb.Node) error {
	nodeId, err := NodeResolveId(node.Id, &o.IdMap)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, RPCErrorIRISpecific, "tmp:X is invalid", "node.id")
	}

	nodeIRI := &iri.Node{
		Type: uint16(node.Type),
		Id:   nodeId,
	}

	nodePerms := &perms.Node{}

	if vldErr := nodeIRI.ValidateIRI(nodePerms); vldErr != nil {
		return status.Errorf(codes.InvalidArgument, RPCErrorIRISpecific, vldErr)
	}

	o.tr.Clear(nodeIRI.GetKey(o.server.dbNode))

	if DebugServerRequests {
		o.server.logEvent(fmt.Sprintf("T.NodeDelete(%v) = %v", o.action, nodeIRI.GetPath()))
	}

	return o.stream.Send(&pb.TransactionActionResponse{
		Status:   pb.MutationStatus_SUCCESS,
		ActionId: o.action.ActionId,
	})
}
