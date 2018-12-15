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

func (o *TransactionOperation) IndexUpdate(index *pb.Index) error{
	indexIRI := &iri.NodeIndex{
		Node: iri.NodeResolveId(index.Node, &o.IdMap),
		IndexId: uint16(index.Type),
		Value: index.Value,
	}

	o.tr.Set(indexIRI.GetKey(o.server.dbIndex), PreparePayload(index.Properties))

	if DebugServerRequests {
		o.server.logEvent(fmt.Sprintf("T.IndexUpdate(%v) = %v", o.action, indexIRI.GetPath()))
	}

	return o.stream.Send(&pb.TransactionActionResponse{
		Status: pb.MutationStatus_SUCCESS,
		ActionId: o.action.ActionId,
	})
}
