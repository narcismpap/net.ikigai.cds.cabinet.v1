// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package server

import (
	"cds.ikigai.net/cabinet.v1/perms"
	pb "cds.ikigai.net/cabinet.v1/rpc"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/rand"
	"time"
)

func (o *TransactionOperation) CounterIncrement(counter *pb.Counter) error {
	counterPerms := &perms.Count{}

	cntIRI, err := ResolveCounterIRI(counter, &o.IdMap, counterPerms)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, RPCErrorIRISpecific, err)
	}

	incVal, err := Int64ToBytes(int64(counter.Value))
	if err != nil {
		return status.Error(codes.InvalidArgument, RPCErrorArgumentInvalid)
	}

	// increment a random position in the counter
	rand.Seed(time.Now().UnixNano())
	cKey := cntIRI.GetKey(o.server.dbCount, CounterKeys[rand.Intn(16)])

	o.tr.Add(cKey, incVal)

	if DebugServerRequests {
		o.server.logEvent(fmt.Sprintf("T.CounterIncrement(%v) = %v", o.action, cntIRI.GetPath()))
	}

	return o.stream.Send(&pb.TransactionActionResponse{
		Status:   pb.MutationStatus_SUCCESS,
		ActionId: o.action.ActionId,
	})
}
