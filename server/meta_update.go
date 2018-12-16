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

func (o *TransactionOperation) MetaUpdate(meta *pb.Meta) error{
	metaIRI, err := iri.ResolveMetaIRI(meta, &o.IdMap)

	if err != nil {
		return status.Errorf(codes.InvalidArgument, RPCErrorIRISpecific, err)
	}

	o.tr.Set(metaIRI.GetKey(o.server.dbMeta), PreparePayload(meta.Val))

	if DebugServerRequests {
		o.server.logEvent(fmt.Sprintf("T.MetaUpdate(%v) = %v\n\n", o.action, metaIRI.GetPath()))
	}

	return o.stream.Send(&pb.TransactionActionResponse{
		Status: pb.MutationStatus_SUCCESS,
		ActionId: o.action.ActionId,
	})
}
