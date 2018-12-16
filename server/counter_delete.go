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

func (o *TransactionOperation) CounterDelete(counter *pb.Counter) error{
	cntIRI, err := iri.ResolveCounterIRI(counter, &o.IdMap)

	if err != nil {
		return status.Errorf(codes.InvalidArgument, RPCErrorIRISpecific, err)
	}

	o.tr.ClearRange(cntIRI.GetKeyRange(o.server.dbCount))

	if DebugServerRequests {
		o.server.logEvent(fmt.Sprintf("T.CounterDelete(%v) = %v", o.action, cntIRI.GetPath()))
	}

	return o.stream.Send(&pb.TransactionActionResponse{
		Status: pb.MutationStatus_SUCCESS,
		ActionId: o.action.ActionId,
	})
}
