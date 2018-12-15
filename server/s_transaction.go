// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package server

import (
	"cds.ikigai.net/cabinet.v1/iri"
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
	_, err := s.FdbConn.Transact(func (tr fdb.Transaction) (ret interface{}, err error) {
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
					cntIRI, err := iri.ResolveCounterIRI(tOpr.CounterIncrement, &idMap)

					if err != nil {
						return nil, status.Error(codes.InvalidArgument, RPCErrorInvalidIRI)
					}

					incVal, err := Int64ToBytes(int64(tOpr.CounterIncrement.Value))


					if err != nil {
						return nil, status.Error(codes.InvalidArgument, RPCErrorArgumentInvalid)
					}

					// increment a random position in the counter
					rand.Seed(time.Now().UnixNano())
					cKey := cntIRI.GetKey(s, CounterKeys[rand.Intn(16)])

					tr.Add(cKey, incVal)

					if DebugServerRequests {
						s.logEvent(fmt.Sprintf("T.CounterIncrement(%v) = %v", tAct, cntIRI.GetPath()))
					}

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
					})

					if err != nil{
						return nil, err
					}

				case *pb.TransactionAction_CounterDelete:
					cntIRI, err := iri.ResolveCounterIRI(tOpr.CounterDelete, &idMap)

					if err != nil {
						return nil, status.Error(codes.InvalidArgument, RPCErrorInvalidIRI)
					}

					tr.ClearRange(cntIRI.GetKeyRange(s))

					if DebugServerRequests {
						s.logEvent(fmt.Sprintf("T.CounterDelete(%v) = %v", tAct, cntIRI.GetPath()))
					}

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
					})

					if err != nil{
						return nil, err
					}

				case *pb.TransactionAction_CounterRegister:
					cntIRI, err := iri.ResolveCounterIRI(tOpr.CounterRegister, &idMap)

					if err != nil {
						return nil, status.Error(codes.InvalidArgument, RPCErrorInvalidIRI)
					}

					incVal, err := Int64ToBytes(int64(0))
					CheckFatalError(err)

					for x := range CounterKeys{
						tr.Set(cntIRI.GetKey(s, CounterKeys[x]), incVal)
					}

					if DebugServerRequests {
						s.logEvent(fmt.Sprintf("T.CounterRegister(%v) = %v", tAct, cntIRI.GetPath()))
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
					edgeIRI := &iri.IRIEdge{
						Subject: iri.NodeResolveId(tOpr.EdgeUpdate.Subject, &idMap),
						Predicate: uint16(tOpr.EdgeUpdate.Predicate),
						Target: iri.NodeResolveId(tOpr.EdgeUpdate.Target, &idMap),
					}

					tr.Set(edgeIRI.GetKey(s), PreparePayload(tOpr.EdgeUpdate.Properties))

					if DebugServerRequests {
						s.logEvent(fmt.Sprintf("T.EdgeUpdate(%v) = %v", tAct, edgeIRI.GetPath()))
					}

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
					})

					if err != nil{
						return nil, err
					}

				case *pb.TransactionAction_EdgeDelete:
					edgeIRI := &iri.IRIEdge{
						Subject: iri.NodeResolveId(tOpr.EdgeDelete.Subject, &idMap),
						Predicate: uint16(tOpr.EdgeDelete.Predicate),
						Target: iri.NodeResolveId(tOpr.EdgeDelete.Target, &idMap),
					}

					tr.Clear(edgeIRI.GetKey(s))

					if DebugServerRequests {
						s.logEvent(fmt.Sprintf("T.EdgeDelete(%v) = %v", tAct, edgeIRI.GetPath()))
					}

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
					})

					if err != nil{
						return nil, err
					}

				case *pb.TransactionAction_EdgeClear:
					edgeIRI := &iri.IRIEdge{
						Subject: iri.NodeResolveId(tOpr.EdgeClear.Subject, &idMap),
						Predicate: uint16(tOpr.EdgeClear.Predicate),
					}

					tr.ClearRange(edgeIRI.GetClearRange(s))

					if DebugServerRequests {
						s.logEvent(fmt.Sprintf("T.EdgeClear(%v) = %v", tAct, edgeIRI.GetPath()))
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
					indexIRI := &iri.IRINodeIndex{
						Node: iri.NodeResolveId(tOpr.IndexUpdate.Node, &idMap),
						IndexId: uint16(tOpr.IndexUpdate.Type),
						Value: tOpr.IndexUpdate.Value,
					}

					tr.Set(indexIRI.GetKey(s), PreparePayload(tOpr.IndexUpdate.Properties))

					if DebugServerRequests {
						s.logEvent(fmt.Sprintf("T.IndexUpdate(%v) = %v", tAct, indexIRI.GetPath()))
					}

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
					})

					if err != nil{
						return nil, err
					}

				case *pb.TransactionAction_IndexDelete:
					indexIRI := &iri.IRINodeIndex{
						Node: iri.NodeResolveId(tOpr.IndexDelete.Node, &idMap),
						IndexId: uint16(tOpr.IndexDelete.Type),
						Value: tOpr.IndexDelete.Value,
					}

					tr.Clear(indexIRI.GetKey(s))

					if DebugServerRequests {
						s.logEvent(fmt.Sprintf("T.IndexDelete(%v) = %v", tAct, indexIRI.GetPath()))
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
					metaIRI, err := iri.ResolveMetaIRI(tOpr.MetaUpdate, &idMap)

					if err != nil {
						return nil, status.Error(codes.InvalidArgument, RPCErrorInvalidIRI)
					}

					tr.Set(metaIRI.GetKey(s), PreparePayload(tOpr.MetaUpdate.Val))

					if DebugServerRequests {
						s.logEvent(fmt.Sprintf("T.MetaUpdate(%v) = %v", tAct, metaIRI.GetPath()))
					}

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
					})

					if err != nil{
						return nil, err
					}

				case *pb.TransactionAction_MetaDelete:
					metaIRI, err := iri.ResolveMetaIRI(tOpr.MetaDelete, &idMap)

					if err != nil {
						return nil, status.Error(codes.InvalidArgument, RPCErrorInvalidIRI)
					}

					tr.Clear(metaIRI.GetKey(s))

					if DebugServerRequests {
						s.logEvent(fmt.Sprintf("T.MetaDelete(%v) = %v", tAct, metaIRI.GetPath()))
					}

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
					})

					if err != nil{
						return nil, err
					}

				case *pb.TransactionAction_MetaClear:
					metaIRI, err := iri.ResolveMetaIRI(tOpr.MetaClear, &idMap)

					if err != nil {
						return nil, status.Error(codes.InvalidArgument, RPCErrorInvalidIRI)
					}

					tr.ClearRange(metaIRI.GetClearRange(s))

					if DebugServerRequests {
						s.logEvent(fmt.Sprintf("T.MetaClear(%v) = %v", tAct, metaIRI.GetPath()))
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
					nodeIRI := &iri.IRINode{Type: uint16(tOpr.NodeCreate.Type), Id: newID}
					tr.Set(nodeIRI.GetKey(s), PreparePayload(tOpr.NodeCreate.Properties))

					if DebugServerRequests {
						s.logEvent(fmt.Sprintf("T.NodeCreate(%v) = %v", tAct, nodeIRI.GetPath()))
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
					nodeIRI := &iri.IRINode{
						Type: uint16(tOpr.NodeUpdate.Type),
						Id: iri.NodeResolveId(tOpr.NodeUpdate.Id, &idMap),
					}

					tr.Set(nodeIRI.GetKey(s), PreparePayload(tOpr.NodeUpdate.Properties))

					if DebugServerRequests {
						s.logEvent(fmt.Sprintf("T.NodeUpdate(%v) = %v", tAct, nodeIRI.GetPath()))
					}

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
					})

					if err != nil{
						return nil, err
					}

				case *pb.TransactionAction_NodeDelete:
					nodeIRI := &iri.IRINode{
						Type: uint16(tOpr.NodeDelete.Type),
						Id: iri.NodeResolveId(tOpr.NodeDelete.Id, &idMap),
					}

					tr.Clear(nodeIRI.GetKey(s))

					if DebugServerRequests {
						s.logEvent(fmt.Sprintf("T.NodeDelete(%v) = %v", tAct, nodeIRI.GetPath()))
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
