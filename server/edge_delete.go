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

func (o *TransactionOperation) EdgeDelete(edge *pb.Edge) error {
	subjectID, err := NodeResolveId(edge.Subject, &o.IdMap)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, RPCErrorFieldSpecific, "tmp:X is invalid", "edge.subject")
	}

	targetID, err := NodeResolveId(edge.Target, &o.IdMap)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, RPCErrorFieldSpecific, "tmp:X is invalid", "edge.target")
	}

	edgeIRI := &iri.Edge{
		Subject:   subjectID,
		Predicate: uint16(edge.Predicate),
		Target:    targetID,
	}

	edgePerms := &perms.Edge{
		AllowTargetWildcard:    false,
		AllowPredicateWildcard: false,
	}

	if vldErr := edgeIRI.ValidateIRI(edgePerms); vldErr != nil {
		return status.Errorf(codes.InvalidArgument, RPCErrorIRISpecific, vldErr)
	}

	o.tr.Clear(edgeIRI.GetKey(o.server.dbEdge))

	if DebugServerRequests {
		o.server.logEvent(fmt.Sprintf("T.EdgeDelete(%v) = %v", o.action, edgeIRI.GetPath()))
	}

	return o.stream.Send(&pb.TransactionActionResponse{
		Status:   pb.MutationStatus_SUCCESS,
		ActionId: o.action.ActionId,
	})
}
