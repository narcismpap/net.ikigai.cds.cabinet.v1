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
	"context"
	"fmt"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *CDSCabinetServer) ReadCheck(ctx context.Context, rcRq *pb.ReadCheckRequest) (*pb.ReadCheckResponse, error) {
	if DebugServerRequests {
		s.logEvent(fmt.Sprintf("ReadCheck(%v)", rcRq))
	}

	pbr, err := s.fdb.Transact(func(tr fdb.Transaction) (ret interface{}, err error) {
		return readCheckLogic(tr, s, rcRq)
	})

	return &pb.ReadCheckResponse{Result: pbr.(bool)}, err
}

func readCheckLogic(tr fdb.Transaction, s *CDSCabinetServer, rcRq *pb.ReadCheckRequest) (bool, error) {
	source, err := getIRIValue(tr, s, rcRq.Source)
	var target string

	if err != nil {
		return false, err
	}

	switch targetObj := rcRq.Target.Target.(type) {
	case *pb.CheckTarget_Iri:
		targetBytes, err := getIRIValue(tr, s, targetObj.Iri)
		if err != nil {
			return false, err
		}

		target = string(targetBytes)

	case *pb.CheckTarget_Val:
		target = targetObj.Val

	default:
		return false, status.Errorf(codes.InvalidArgument, RPCErrorFieldSpecific, "unknown target type", "readCheck.target")
	}

	switch rcRq.Operator {
	case pb.CheckOperators_EXISTS:
		return source != nil, nil
	case pb.CheckOperators_EQUAL:
		return string(source) == target, nil
	case pb.CheckOperators_NOT_EQUAL:
		return string(source) != target, nil
	case pb.CheckOperators_TOUCH:
		return true, nil
	default:
		return false, status.Errorf(codes.InvalidArgument, RPCErrorFieldSpecific, "unknown or unimplemented operator", "readCheck.operator")
	}

}

func getIRIValue(tr fdb.Transaction, s *CDSCabinetServer, reqIRI string) ([]byte, error) {
	varIRI, err := iri.Parse(reqIRI)
	var source []byte

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, RPCErrorIRISpecific, "source cannot be parsed")
	}

	switch sourceIRI := varIRI.(type) {

	case *iri.Edge:
		edgePerms := &perms.Edge{}

		if vldErr := sourceIRI.ValidateIRI(edgePerms); vldErr != nil {
			return nil, status.Errorf(codes.InvalidArgument, RPCErrorIRISpecific, vldErr)
		}

		source = tr.Get(sourceIRI.GetKey(s.dbEdge)).MustGet()

	case *iri.NodeIndex:
		idxPerms := &perms.Index{}

		if vldErr := sourceIRI.ValidateIRI(idxPerms); vldErr != nil {
			return nil, status.Errorf(codes.InvalidArgument, RPCErrorIRISpecific, vldErr)
		}

		source = tr.Get(sourceIRI.GetKey(s.dbIndex)).MustGet()

	case *iri.NodeMeta:
		metaPerms := &perms.Meta{}

		if vldErr := sourceIRI.ValidateIRI(metaPerms); vldErr != nil {
			return nil, status.Errorf(codes.InvalidArgument, RPCErrorIRISpecific, vldErr)
		}

		source = tr.Get(sourceIRI.GetKey(s.dbMeta)).MustGet()

	case *iri.EdgeMeta:
		metaPerms := &perms.Meta{}

		if vldErr := sourceIRI.ValidateIRI(metaPerms); vldErr != nil {
			return nil, status.Errorf(codes.InvalidArgument, RPCErrorIRISpecific, vldErr)
		}

		source = tr.Get(sourceIRI.GetKey(s.dbMeta)).MustGet()

	case *iri.Node:
		nodePerms := &perms.Node{}

		if vldErr := sourceIRI.ValidateIRI(nodePerms); vldErr != nil {
			return nil, status.Errorf(codes.InvalidArgument, RPCErrorIRISpecific, vldErr)
		}

		source = tr.Get(sourceIRI.GetKey(s.dbNode)).MustGet()

	default:
		return nil, status.Error(codes.Unimplemented, RPCErrorInvalidAction)
	}

	return source, nil
}
