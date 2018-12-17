// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package server

import (
	pb "cds.ikigai.net/cabinet.v1/rpc"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

func (s *CDSCabinetServer) Transaction(bStream pb.CDSCabinet_TransactionServer) error{
	_, err := s.fdb.Transact(func (tr fdb.Transaction) (ret interface{}, err error) {

		trx := TransactionOperation{
			IdMap: make(map[string]string),
			UsedActions: make(map[uint32]bool),
			stream: bStream,
			tr: tr,
			server: s,
		}

		for {
			trx.action, err = bStream.Recv()

			if err == io.EOF {
				return nil, nil
			}else if err != nil {
				return nil, err
			}

			if _, ok := trx.UsedActions[trx.action.ActionId]; ok {
				return nil, status.Error(codes.Unimplemented, RPCErrorRepeatAction)
			}else{
				trx.UsedActions[trx.action.ActionId] = true
			}

			switch tOpr := trx.action.Action.(type) {
				// Counter
				case *pb.TransactionAction_CounterIncrement:
					err = trx.CounterIncrement(tOpr.CounterIncrement)

				case *pb.TransactionAction_CounterDelete:
					err = trx.CounterDelete(tOpr.CounterDelete)

				case *pb.TransactionAction_CounterRegister:
					err = trx.CounterRegister(tOpr.CounterRegister)


				// Edge
				case *pb.TransactionAction_EdgeUpdate:
					err = trx.EdgeUpdate(tOpr.EdgeUpdate)

				case *pb.TransactionAction_EdgeDelete:
					err = trx.EdgeDelete(tOpr.EdgeDelete)

				case *pb.TransactionAction_EdgeClear:
					err = trx.EdgeClear(tOpr.EdgeClear)


				// Index
				case *pb.TransactionAction_IndexUpdate:
					err = trx.IndexUpdate(tOpr.IndexUpdate)

				case *pb.TransactionAction_IndexDelete:
					err = trx.IndexDelete(tOpr.IndexDelete)


				// Meta
				case *pb.TransactionAction_MetaUpdate:
					err = trx.MetaUpdate(tOpr.MetaUpdate)

				case *pb.TransactionAction_MetaDelete:
					err = trx.MetaDelete(tOpr.MetaDelete)

				case *pb.TransactionAction_MetaClear:
					err = trx.MetaClear(tOpr.MetaClear)


				// Node
				case *pb.TransactionAction_NodeCreate:
					err = trx.NodeCreate(tOpr.NodeCreate)

				case *pb.TransactionAction_NodeUpdate:
					err = trx.NodeUpdate(tOpr.NodeUpdate)

				case *pb.TransactionAction_NodeDelete:
					err = trx.NodeDelete(tOpr.NodeDelete)


				// Read Check
				case *pb.TransactionAction_ReadCheck:
					err = trx.ReadCheck(tOpr.ReadCheck)

				default:
					return nil, status.Error(codes.Unimplemented, RPCErrorInvalidAction)
			}

			if err != nil{
				return nil, err
			}
		}
	})

	if err != nil{
		return err
	}

	return nil
}
