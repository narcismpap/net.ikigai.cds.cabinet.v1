// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package main

import (
	pb "cds.ikigai.net/cabinet.v1/rpc"
	"context"
	"fmt"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"math/rand"
	"strings"
	"time"
)

func (s *CDSCabinetServer) Transaction(bStream pb.CDSCabinet_TransactionServer) error{
	_, err := s.fDb.Transact(func (tr fdb.Transaction) (ret interface{}, err error) {
		idMap := make(map[string]string)
		usedAction := make(map[uint32]bool)

		for {
			tAct, err := bStream.Recv()

			if err == io.EOF {
				return nil, nil
			}else if err != nil {
				return nil, err
			}

			if _, ok := usedAction[tAct.ActionId]; ok {
				return nil, status.Error(codes.Unimplemented, RPCErrorRepeatAction)
			}else{
				usedAction[tAct.ActionId] = true
			}

			switch tOpr := tAct.Action.(type) {

				// Counter
				case *pb.TransactionAction_CounterIncrement:
					cntIRI, err := resolveCounterIRI(tOpr.CounterIncrement, &idMap)

					if err != nil {
						return nil, status.Error(codes.InvalidArgument, RPCErrorInvalidIRI)
					}

					incVal, err := intToBytes(int64(tOpr.CounterIncrement.Value))


					if err != nil {
						return nil, status.Error(codes.InvalidArgument, RPCErrorArgumentInvalid)
					}

					// increment a random position in the counter
					rand.Seed(time.Now().UnixNano())
					cKey := cntIRI.getKey(s, CounterKeys[rand.Intn(16)])

					tr.Add(cKey, incVal)

					if DebugServerRequests {
						s.logEvent(fmt.Sprintf("T.CounterIncrement(%v) = %v", tAct, cntIRI.getPath()))
					}

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
					})

					if err != nil{
						return nil, err
					}

				case *pb.TransactionAction_CounterDelete:
					cntIRI, err := resolveCounterIRI(tOpr.CounterDelete, &idMap)

					if err != nil {
						return nil, status.Error(codes.InvalidArgument, RPCErrorInvalidIRI)
					}

					tr.ClearRange(cntIRI.getKeyRange(s))

					if DebugServerRequests {
						s.logEvent(fmt.Sprintf("T.CounterDelete(%v) = %v", tAct, cntIRI.getPath()))
					}

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
					})

					if err != nil{
						return nil, err
					}

				case *pb.TransactionAction_CounterRegister:
					cntIRI, err := resolveCounterIRI(tOpr.CounterRegister, &idMap)

					if err != nil {
						return nil, status.Error(codes.InvalidArgument, RPCErrorInvalidIRI)
					}

					incVal, err := intToBytes(int64(0))
					CheckFatalError(err)

					for x := range CounterKeys{
						tr.Set(cntIRI.getKey(s, CounterKeys[x]), incVal)
					}

					if DebugServerRequests {
						s.logEvent(fmt.Sprintf("T.CounterRegister(%v) = %v", tAct, cntIRI.getPath()))
					}

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
					})

					if err != nil{
						return nil, err
					}

				// Edge
				case *pb.TransactionAction_EdgeUpdate:
					edgeIRI := &IRIEdge{
						Subject: nodeResolveId(tOpr.EdgeUpdate.Subject, &idMap),
						Predicate: uint16(tOpr.EdgeUpdate.Predicate),
						Target: nodeResolveId(tOpr.EdgeUpdate.Target, &idMap),
					}

					tr.Set(edgeIRI.getKey(s), prepareProperties(tOpr.EdgeUpdate.Properties))

					if DebugServerRequests {
						s.logEvent(fmt.Sprintf("T.EdgeUpdate(%v) = %v", tAct, edgeIRI.getPath()))
					}

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
					})

					if err != nil{
						return nil, err
					}

				case *pb.TransactionAction_EdgeDelete:
					edgeIRI := &IRIEdge{
						Subject: nodeResolveId(tOpr.EdgeDelete.Subject, &idMap),
						Predicate: uint16(tOpr.EdgeDelete.Predicate),
						Target: nodeResolveId(tOpr.EdgeDelete.Target, &idMap),
					}

					tr.Clear(edgeIRI.getKey(s))

					if DebugServerRequests {
						s.logEvent(fmt.Sprintf("T.EdgeDelete(%v) = %v", tAct, edgeIRI.getPath()))
					}

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
					})

					if err != nil{
						return nil, err
					}

				case *pb.TransactionAction_EdgeClear:
					edgeIRI := &IRIEdge{
						Subject: nodeResolveId(tOpr.EdgeClear.Subject, &idMap),
						Predicate: uint16(tOpr.EdgeClear.Predicate),
					}

					tr.ClearRange(edgeIRI.getClearRange(s))

					if DebugServerRequests {
						s.logEvent(fmt.Sprintf("T.EdgeClear(%v) = %v", tAct, edgeIRI.getPath()))
					}

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
					})

					if err != nil{
						return nil, err
					}

				// Indexes
				case *pb.TransactionAction_IndexUpdate:
					indexIRI := &IRINodeIndex{
						Node: nodeResolveId(tOpr.IndexUpdate.Node, &idMap),
						IndexId: uint16(tOpr.IndexUpdate.Type),
						Value: tOpr.IndexUpdate.Value,
					}

					tr.Set(indexIRI.getKey(s), prepareProperties(tOpr.IndexUpdate.Properties))

					if DebugServerRequests {
						s.logEvent(fmt.Sprintf("T.IndexUpdate(%v) = %v", tAct, indexIRI.getPath()))
					}

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
					})

					if err != nil{
						return nil, err
					}

				case *pb.TransactionAction_IndexDelete:
					indexIRI := &IRINodeIndex{
						Node: nodeResolveId(tOpr.IndexDelete.Node, &idMap),
						IndexId: uint16(tOpr.IndexDelete.Type),
						Value: tOpr.IndexDelete.Value,
					}

					tr.Clear(indexIRI.getKey(s))

					if DebugServerRequests {
						s.logEvent(fmt.Sprintf("T.IndexDelete(%v) = %v", tAct, indexIRI.getPath()))
					}

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
					})

					if err != nil{
						return nil, err
					}

				// Meta
				case *pb.TransactionAction_MetaUpdate:
					metaIRI, err := resolveMetaIRI(tOpr.MetaUpdate, &idMap)

					if err != nil {
						return nil, status.Error(codes.InvalidArgument, RPCErrorInvalidIRI)
					}

					tr.Set(metaIRI.getKey(s), prepareProperties(tOpr.MetaUpdate.Val))

					if DebugServerRequests {
						s.logEvent(fmt.Sprintf("T.MetaUpdate(%v) = %v", tAct, metaIRI.getPath()))
					}

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
					})

					if err != nil{
						return nil, err
					}

				case *pb.TransactionAction_MetaDelete:
					metaIRI, err := resolveMetaIRI(tOpr.MetaDelete, &idMap)

					if err != nil {
						return nil, status.Error(codes.InvalidArgument, RPCErrorInvalidIRI)
					}

					tr.Clear(metaIRI.getKey(s))

					if DebugServerRequests {
						s.logEvent(fmt.Sprintf("T.MetaDelete(%v) = %v", tAct, metaIRI.getPath()))
					}

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
					})

					if err != nil{
						return nil, err
					}

				case *pb.TransactionAction_MetaClear:
					metaIRI, err := resolveMetaIRI(tOpr.MetaClear, &idMap)

					if err != nil {
						return nil, status.Error(codes.InvalidArgument, RPCErrorInvalidIRI)
					}

					tr.ClearRange(metaIRI.getClearRange(s))

					if DebugServerRequests {
						s.logEvent(fmt.Sprintf("T.MetaClear(%v) = %v", tAct, metaIRI.getPath()))
					}

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
					})

					if err != nil{
						return nil, err
					}

				// Node
				case *pb.TransactionAction_NodeCreate:
					newIDBytes, err := ksuid.New().MarshalText()
					CheckFatalError(err)

					newID := string(newIDBytes)
					nodeIRI := &IRINode{Type: uint16(tOpr.NodeCreate.Type), Id: newID}
					tr.Set(nodeIRI.getKey(s), prepareProperties(tOpr.NodeCreate.Properties))

					if DebugServerRequests {
						s.logEvent(fmt.Sprintf("T.NodeCreate(%v) = %v", tAct, nodeIRI.getPath()))
					}

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
						Response: &pb.TransactionActionResponse_NodeCreate{NodeCreate: &pb.NodeCreateResponse{Id: nodeIRI.Id}},
					})

					if err != nil{
						return nil, err
					}

					idMap[strings.TrimLeft(tOpr.NodeCreate.Id, "tmp:")] = newID

				case *pb.TransactionAction_NodeUpdate:
					nodeIRI := &IRINode{
						Type: uint16(tOpr.NodeUpdate.Type),
						Id: nodeResolveId(tOpr.NodeUpdate.Id, &idMap),
					}

					tr.Set(nodeIRI.getKey(s), prepareProperties(tOpr.NodeUpdate.Properties))

					if DebugServerRequests {
						s.logEvent(fmt.Sprintf("T.NodeUpdate(%v) = %v", tAct, nodeIRI.getPath()))
					}

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
					})

					if err != nil{
						return nil, err
					}

				case *pb.TransactionAction_NodeDelete:
					nodeIRI := &IRINode{
						Type: uint16(tOpr.NodeDelete.Type),
						Id: nodeResolveId(tOpr.NodeDelete.Id, &idMap),
					}

					tr.Clear(nodeIRI.getKey(s))

					if DebugServerRequests {
						s.logEvent(fmt.Sprintf("T.NodeDelete(%v) = %v", tAct, nodeIRI.getPath()))
					}

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
					})

					if err != nil{
						return nil, err
					}

				// ReadCheck
				case *pb.TransactionAction_ReadCheck:
					// tOpr.ReadCheck
					break

				default:
					return nil, status.Error(codes.Unimplemented, RPCErrorInvalidAction)
			}
		}
	})

	if err != nil{
		return err
	}

	return nil
}

func (s *CDSCabinetServer) ReadCheck(ctx context.Context, readRq *pb.ReadCheckRequest) (*pb.ReadCheckResponse, error){
	return nil, nil
}
