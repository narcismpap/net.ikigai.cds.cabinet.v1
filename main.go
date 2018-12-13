// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.


package main

import (
	pb "cds.ikigai.net/cabinet.v1/rpc"
	"flag"
	"fmt"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

var (
	port = flag.Int("port", 8888, "The server port")
)

type CDSCabinetServer struct{
	version int32

	fDb fdb.Transactor
	dbContainer directory.DirectorySubspace

	dbNode subspace.Subspace
	dbEdge subspace.Subspace
	dbIndex subspace.Subspace
	dbMeta subspace.Subspace
	dbCnt subspace.Subspace
	dbSeq subspace.Subspace
}

func newCDSServer() *CDSCabinetServer {
	fdb.MustAPIVersion(600)
	db := fdb.MustOpenDefault()

	var activeContainer = "test"
	container, err := directory.CreateOrOpen(db, []string{activeContainer}, nil)

	if err != nil {
		log.Fatal(err)
	}

	s := &CDSCabinetServer{
		version: 1,
		fDb: db,
		dbContainer: container,

		dbNode: container.Sub("n"),
		dbEdge: container.Sub("e"),
		dbIndex: container.Sub("i"),
		dbMeta: container.Sub("m"),
		dbCnt: container.Sub("c"),
		dbSeq: container.Sub("s"),
	}

	// install db
	var coreSeq = []string{
		"n", "e", "i", "m", "c",
	}

	_, err = s.fDb.Transact(func (tr fdb.Transaction) (interface{}, error) {
		tr.ClearRange(s.dbContainer)

		for i := range coreSeq {
			tr.Set(s.dbSeq.Pack(tuple.Tuple{coreSeq[i], "l"}), []byte(strconv.FormatUint(uint64(1), 10)))
		}

		log.Printf("[I] Container [%s] is now initialized", activeContainer)
		return nil, nil
	})

	return s
}

const (
	CDSErrFieldInvalid = iota
	CDSErrorFieldRequired
	CDSErrorFieldUnexpected

	CDSErrorNotFound
	CDSErrorPermission
	CDSErrorBadRecord

	CDSListNoPagination
)

type CabinetError struct {
	err  	string
	code 	int16
	field 	string
}

func (e *CabinetError) Error() string {
	return fmt.Sprintf("Cabinet Error #%d on %s: %s", e.code, e.field, e.err)
}

func main(){
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))

	if err != nil {
		log.Fatalf("[E] Failed to bind on :%d: %v", *port, err)
	}

	log.Printf("[I] net.ikigai.cds.cabinet.v1 running on port %d", *port)

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	pb.RegisterCDSCabinetServer(grpcServer, newCDSServer())
	err = grpcServer.Serve(lis)

	if err != nil{
		log.Fatalf("[E] Error during gRPC execution: %v", err)
	}
}
