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

func (o *TransactionOperation) EdgeDelete(edge *pb.Edge) error{
	edgeIRI := &iri.Edge{
		Subject: iri.NodeResolveId(edge.Subject, &o.IdMap),
		Predicate: uint16(edge.Predicate),
		Target: iri.NodeResolveId(edge.Target, &o.IdMap),
	}

	o.tr.Clear(edgeIRI.GetKey(o.server.dbEdge))

	if DebugServerRequests {
		o.server.logEvent(fmt.Sprintf("T.EdgeDelete(%v) = %v", o.action, edgeIRI.GetPath()))
	}

	return o.stream.Send(&pb.TransactionActionResponse{
		Status: pb.MutationStatus_SUCCESS,
		ActionId: o.action.ActionId,
	})
}
