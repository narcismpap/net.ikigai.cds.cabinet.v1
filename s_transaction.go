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
	"io"
	"log"
	"math/rand"
	"time"
)

func (s *CDSCabinetServer) Transaction(bStream pb.CDSCabinet_TransactionServer) error{
	_, err := s.fDb.Transact(func (tr fdb.Transaction) (ret interface{}, err error) {
		idMap := make(map[string]string)

		for {
			tAct, err := bStream.Recv()

			if err == io.EOF {
				return nil, nil
			}else if err != nil {
				return nil, err
			}

			switch tOpr := tAct.Action.(type) {

				// Counter
				case *pb.TransactionAction_CounterIncrement:
					iriCnt, err := resolveCounterIRI(tOpr.CounterIncrement, &idMap)

					if err != nil {
						return nil, bStream.Send(&pb.TransactionActionResponse{
							Status:   pb.MutationStatus_GENERIC_FAILURE,
							ActionId: tAct.ActionId,
							Error:    "Unknown object as Counter Type",
						})
					}

					incVal, err := intToBytes(int64(tOpr.CounterIncrement.Value))

					if err != nil {
						return nil, bStream.Send(&pb.TransactionActionResponse{
							Status:   pb.MutationStatus_GENERIC_FAILURE,
							ActionId: tAct.ActionId,
							Error:    "Unable to build increment",
						})
					}

					// increment a random position in the counter
					rand.Seed(time.Now().UnixNano())
					tr.Add(iriCnt.getKey(s, CounterKeys[rand.Intn(16)]), incVal)

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
					})

					if err != nil{
						return nil, err
					}

				case *pb.TransactionAction_CounterDelete:
					iriCnt, err := resolveCounterIRI(tOpr.CounterDelete, &idMap)

					if err != nil {
						return nil, bStream.Send(&pb.TransactionActionResponse{
							Status:   pb.MutationStatus_GENERIC_FAILURE,
							ActionId: tAct.ActionId,
							Error:    "Unknown object as Counter Type",
						})
					}

					tr.ClearRange(iriCnt.getKeyRange(s))

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
					})

					if err != nil{
						return nil, err
					}

				case *pb.TransactionAction_CounterRegister:
					iriCnt, err := resolveCounterIRI(tOpr.CounterRegister, &idMap)

					if err != nil {
						return nil, bStream.Send(&pb.TransactionActionResponse{
							Status:   pb.MutationStatus_GENERIC_FAILURE,
							ActionId: tAct.ActionId,
								Error:    "Unknown object as Counter Type",
						})
					}

					incVal, err := intToBytes(int64(0))

					if err != nil{
						log.Fatalf("Unable to intToBytes from known value: 0")
					}

					for x := range CounterKeys{
						tr.Set(iriCnt.getKey(s, CounterKeys[x]), incVal)
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
					iriEdge := &IRIEdge{
						Subject: nodeResolveId(tOpr.EdgeUpdate.Subject, &idMap),
						Predicate: uint16(tOpr.EdgeUpdate.Predicate),
						Target: nodeResolveId(tOpr.EdgeUpdate.Target, &idMap),
					}

					tr.Set(iriEdge.getKey(s), prepareProperties(tOpr.EdgeUpdate.Properties))

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
					})

					if err != nil{
						return nil, err
					}

				case *pb.TransactionAction_EdgeDelete:
					iriEdge := &IRIEdge{
						Subject: nodeResolveId(tOpr.EdgeDelete.Subject, &idMap),
						Predicate: uint16(tOpr.EdgeDelete.Predicate),
						Target: nodeResolveId(tOpr.EdgeDelete.Target, &idMap),
					}

					tr.Clear(iriEdge.getKey(s))

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
					})

					if err != nil{
						return nil, err
					}

				case *pb.TransactionAction_EdgeClear:
					iriEdge := &IRIEdge{
						Subject: nodeResolveId(tOpr.EdgeClear.Subject, &idMap),
						Predicate: uint16(tOpr.EdgeClear.Predicate),
					}

					tr.ClearRange(iriEdge.getClearRange(s))

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
					})

					if err != nil{
						return nil, err
					}

				// Indexes
				case *pb.TransactionAction_IndexUpdate:
					iriIndex := &IRINodeIndex{
						Node: nodeResolveId(tOpr.IndexUpdate.Node, &idMap),
						IndexId: uint16(tOpr.IndexUpdate.Type),
						Value: tOpr.IndexUpdate.Value,
					}

					tr.Set(iriIndex.getKey(s), prepareProperties(tOpr.IndexUpdate.Properties))

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
					})

					if err != nil{
						return nil, err
					}

				case *pb.TransactionAction_IndexDelete:
					iriIndex := &IRINodeIndex{
						Node: nodeResolveId(tOpr.IndexDelete.Node, &idMap),
						IndexId: uint16(tOpr.IndexDelete.Type),
						Value: tOpr.IndexDelete.Value,
					}

					tr.Clear(iriIndex.getKey(s))

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
					})

					if err != nil{
						return nil, err
					}

				// Meta
				case *pb.TransactionAction_MetaUpdate:
					iri, err := resolveMetaIRI(tOpr.MetaUpdate, &idMap)

					if err != nil {
						return nil, bStream.Send(&pb.TransactionActionResponse{
							Status:   pb.MutationStatus_GENERIC_FAILURE,
							ActionId: tAct.ActionId,
							Error:    "Unknown object for Meta",
						})
					}

					tr.Set(iri.getKey(s), prepareProperties(tOpr.MetaUpdate.Val))

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
					})

					if err != nil{
						return nil, err
					}

				case *pb.TransactionAction_MetaDelete:
					iri, err := resolveMetaIRI(tOpr.MetaDelete, &idMap)

					if err != nil {
						return nil, bStream.Send(&pb.TransactionActionResponse{
							Status:   pb.MutationStatus_GENERIC_FAILURE,
							ActionId: tAct.ActionId,
							Error:    "Unknown object for Meta",
						})
					}

					tr.Clear(iri.getKey(s))

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
					})

					if err != nil{
						return nil, err
					}

				case *pb.TransactionAction_MetaClear:
					iri, err := resolveMetaIRI(tOpr.MetaClear, &idMap)

					if err != nil {
						return nil, bStream.Send(&pb.TransactionActionResponse{
							Status:   pb.MutationStatus_GENERIC_FAILURE,
							ActionId: tAct.ActionId,
							Error:    "Unknown object for Meta",
						})
					}

					tr.ClearRange(iri.getClearRange(s))

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
					})

					if err != nil{
						return nil, err
					}

				// Node
				case *pb.TransactionAction_NodeCreate:
					newID, err := ksuid.New().MarshalText()

					if err != nil{
						return nil, bStream.Send(&pb.TransactionActionResponse{
							Status: pb.MutationStatus_GENERIC_FAILURE,
							ActionId: tAct.ActionId,
							Error: "cannot create node ID",
						})
					}

					nodeIRI := &IRINode{Type: uint16(tOpr.NodeCreate.Type), Id: string(newID)}
					tr.Set(nodeIRI.getKey(s), prepareProperties(tOpr.NodeCreate.Properties))

					err = bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_SUCCESS,
						ActionId: tAct.ActionId,
						Response: &pb.TransactionActionResponse_NodeCreate{NodeCreate: &pb.NodeCreateResponse{Id: nodeIRI.Id}},
					})

					if err != nil{
						return nil, err
					}

					idMap[tOpr.NodeCreate.Id] = string(newID)

				case *pb.TransactionAction_NodeUpdate:
					nodeIRI := &IRINode{
						Type: uint16(tOpr.NodeUpdate.Type),
						Id: nodeResolveId(tOpr.NodeUpdate.Id, &idMap),
					}

					tr.Set(nodeIRI.getKey(s), prepareProperties(tOpr.NodeUpdate.Properties))

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
					return nil, bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_GENERIC_FAILURE,
						ActionId: tAct.ActionId,
						Error: fmt.Sprintf("Unknown trx action %s", tAct.Action),
					})
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