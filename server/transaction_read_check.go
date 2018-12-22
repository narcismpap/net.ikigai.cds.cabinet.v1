// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package server

import (
	pb "cds.ikigai.net/cabinet.v1/rpc"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *TransactionOperation) ReadCheck(rc *pb.ReadCheckRequest) error {
	if DebugServerRequests {
		o.server.logEvent(fmt.Sprintf("T.ReadCheck(%v)", rc))
	}

	rcStatus, err := readCheckLogic(o.tr, o.server, rc)

	if err != nil {
		return err
	}

	if rcStatus == false {
		_ = o.stream.Send(&pb.TransactionActionResponse{
			Status:   pb.MutationStatus_READ_CHECK_FAILURE,
			ActionId: o.action.ActionId,
		})

		return status.Error(codes.InvalidArgument, RPCErrorReadCheck)
	}

	return o.stream.Send(&pb.TransactionActionResponse{
		Status:   pb.MutationStatus_SUCCESS,
		ActionId: o.action.ActionId,
	})
}
