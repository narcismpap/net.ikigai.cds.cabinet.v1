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

func (o *TransactionOperation) CounterRegister(counter *pb.Counter) error{
	counterPerms := &perms.Count{}

	cntIRI, err := iri.ResolveCounterIRI(counter, &o.IdMap, counterPerms)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, RPCErrorIRISpecific, err)
	}

	incVal, err := Int64ToBytes(int64(0))
	CheckFatalError(err)

	for x := range CounterKeys{
		o.tr.Set(cntIRI.GetKey(o.server.dbCount, CounterKeys[x]), incVal)
	}

	if DebugServerRequests {
		o.server.logEvent(fmt.Sprintf("T.CounterRegister(%v) = %v", o.action, cntIRI.GetPath()))
	}

	return o.stream.Send(&pb.TransactionActionResponse{
		Status: pb.MutationStatus_SUCCESS,
		ActionId: o.action.ActionId,
	})
}
