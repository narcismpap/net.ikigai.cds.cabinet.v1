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
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
	"github.com/segmentio/ksuid"
	"io"
)

func (s *CDSCabinetServer) Transaction(bStream pb.CDSCabinet_TransactionServer) error{
	_, err := s.fDb.Transact(func (tr fdb.Transaction) (ret interface{}, err error) {
		for {
			tAct, err := bStream.Recv()

			if err == io.EOF {
				return nil, bStream.Send(&pb.TransactionActionResponse{})

			}else if err != nil {
				return nil, err
			}

			switch tOpr := tAct.Action.(type) {

				// Counter
				case *pb.TransactionAction_CounterIncrement:
					// tOpr.CounterIncrement
					break

				case *pb.TransactionAction_CounterDelete:
					// tOpr.CounterDelete
					break

				case *pb.TransactionAction_CounterRegister:
					// tOpr.CounterRegister
					break


				// Edge
				case *pb.TransactionAction_EdgeCreate:
					// tOpr.EdgeCreate
					break

				case *pb.TransactionAction_EdgeUpdate:
					// tOpr.EdgeUpdate
					break

				case *pb.TransactionAction_EdgeDelete:
					// tOpr.EdgeDelete
					break


				// Index
				case *pb.TransactionAction_IndexCreate:
					// tOpr.IndexCreate
					break

				case *pb.TransactionAction_IndexUpdate:
					// tOpr.IndexUpdate
					break

				case *pb.TransactionAction_IndexDelete:
					// tOpr.IndexDelete
					break

				// Meta
				case *pb.TransactionAction_MetaCreate:
					// tOpr.MetaCreate
					break

				case *pb.TransactionAction_MetaUpdate:
					// tOpr.MetaUpdate
					break

				case *pb.TransactionAction_MetaDelete:
					// tOpr.MetaDelete
					break

				// Node
				case *pb.TransactionAction_NodeCreate:
					new_id, err := ksuid.New().MarshalBinary()

					if err != nil{
						return nil, bStream.Send(&pb.TransactionActionResponse{
							Status: pb.MutationStatus_GENERIC_FAILURE,
							ActionId: tAct.ActionId,
							Error: "cannot create node ID",
						})
					}

					ob_key := s.dbNode.Sub(tOpr.NodeCreate.Type).Pack(tuple.Tuple{new_id})
					fmt.Print(ob_key)

					break

				case *pb.TransactionAction_NodeUpdate:
					// tOpr.NodeUpdate
					break

				case *pb.TransactionAction_NodeDelete:
					// tOpr.NodeDelete
					break

				// ReadCheck
				case *pb.TransactionAction_ReadCheck:
					// tOpr.ReadCheck
					break

				case nil:
					return nil, bStream.Send(&pb.TransactionActionResponse{
						Status: pb.MutationStatus_GENERIC_FAILURE,
						ActionId: tAct.ActionId,
						Error: "Missing action",
					})

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

