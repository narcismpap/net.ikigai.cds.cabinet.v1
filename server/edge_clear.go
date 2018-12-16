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

func (o *TransactionOperation) EdgeClear(edge *pb.Edge) error{
	edgeIRI := &iri.Edge{
		Subject: 	iri.NodeResolveId(edge.Subject, &o.IdMap),
		Predicate: 	uint16(edge.Predicate),
		Target: 	edge.Target,
	}

	edgePerms := &iri.EdgePermissions{
		AllowTargetWildcard: true,
		AllowPredicateWildcard: false,
	}

	if valdErr := edgeIRI.ValidateIRI(edgePerms); valdErr != nil{
		return status.Errorf(codes.InvalidArgument, RPCErrorIRISpecific, valdErr)
	}

	o.tr.ClearRange(edgeIRI.GetClearRange(o.server.dbEdge))

	if DebugServerRequests {
		o.server.logEvent(fmt.Sprintf("T.EdgeClear(%v) = %v", o.action, edgeIRI.GetPath()))
	}

	return o.stream.Send(&pb.TransactionActionResponse{
		Status: pb.MutationStatus_SUCCESS,
		ActionId: o.action.ActionId,
	})
}
