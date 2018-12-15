// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package server

import (
	pb "cds.ikigai.net/cabinet.v1/rpc"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
)

type TransactionOperation struct{
	IdMap map[string]string
	UsedActions map[uint32]bool

	action *pb.TransactionAction
	stream pb.CDSCabinet_TransactionServer
	tr fdb.Transaction
	server *CDSCabinetServer
}
