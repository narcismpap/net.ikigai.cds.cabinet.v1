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

func (o *TransactionOperation) IndexUpdate(index *pb.Index) error {
	nodeId, err := iri.NodeResolveId(index.Node, &o.IdMap)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, RPCErrorIRISpecific, "tmp:X is invalid", "index.node")
	}

	indexIRI := &iri.NodeIndex{
		Node:    nodeId,
		IndexId: uint16(index.Type),
		Value:   index.Value,
	}

	indexPerms := &perms.Index{}

	if vldErr := indexIRI.ValidateIRI(indexPerms); vldErr != nil {
		return status.Errorf(codes.InvalidArgument, RPCErrorIRISpecific, vldErr)
	}

	o.tr.Set(indexIRI.GetKey(o.server.dbIndex), PreparePayload(index.Properties))

	if DebugServerRequests {
		o.server.logEvent(fmt.Sprintf("T.IndexUpdate(%v) = %v", o.action, indexIRI.GetPath()))
	}

	return o.stream.Send(&pb.TransactionActionResponse{
		Status:   pb.MutationStatus_SUCCESS,
		ActionId: o.action.ActionId,
	})
}
