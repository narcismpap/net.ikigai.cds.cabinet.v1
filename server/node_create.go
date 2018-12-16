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
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

func (o *TransactionOperation) NodeCreate(node *pb.Node) error{
	newIDBytes, err := ksuid.New().MarshalText()
	CheckFatalError(err)

	newID := string(newIDBytes)
	nodeIRI := &iri.Node{Type: uint16(node.Type), Id: newID}

	if vldErr := nodeIRI.ValidateIRI(); vldErr != nil{
		return status.Errorf(codes.InvalidArgument, RPCErrorIRISpecific, vldErr)
	}

	o.tr.Set(nodeIRI.GetKey(o.server.dbNode), PreparePayload(node.Properties))
	o.IdMap[strings.TrimLeft(node.Id, "tmp:")] = newID

	if DebugServerRequests {
		o.server.logEvent(fmt.Sprintf("T.NodeCreate(%v) = %v", o.action, nodeIRI.GetPath()))
	}

	return o.stream.Send(&pb.TransactionActionResponse{
		Status: pb.MutationStatus_SUCCESS,
		ActionId: o.action.ActionId,
		Response: &pb.TransactionActionResponse_NodeCreate{NodeCreate: &pb.NodeCreateResponse{Id: nodeIRI.Id}},
	})
}
