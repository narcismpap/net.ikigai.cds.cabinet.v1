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


func (o *TransactionOperation) MetaClear(meta *pb.Meta) error{
	metaIRI, err := iri.ResolveMetaIRI(meta, &o.IdMap)

	if err != nil {
		return status.Error(codes.InvalidArgument, RPCErrorInvalidIRI)
	}

	o.tr.ClearRange(metaIRI.GetClearRange(o.server.dbMeta))

	if DebugServerRequests {
		o.server.logEvent(fmt.Sprintf("T.MetaClear(%v) = %v", o.action, metaIRI.GetPath()))
	}

	return o.stream.Send(&pb.TransactionActionResponse{
		Status: pb.MutationStatus_SUCCESS,
		ActionId: o.action.ActionId,
	})
}
