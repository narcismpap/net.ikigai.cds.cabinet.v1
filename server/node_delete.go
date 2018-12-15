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
)

func (o *TransactionOperation) NodeDelete(node *pb.Node) error{
	nodeIRI := &iri.Node{
		Type: uint16(node.Type),
		Id: iri.NodeResolveId(node.Id, &o.IdMap),
	}

	o.tr.Clear(nodeIRI.GetKey(o.server.dbNode))

	if DebugServerRequests {
		o.server.logEvent(fmt.Sprintf("T.NodeDelete(%v) = %v", o.action, nodeIRI.GetPath()))
	}

	return o.stream.Send(&pb.TransactionActionResponse{
		Status: pb.MutationStatus_SUCCESS,
		ActionId: o.action.ActionId,
	})
}
