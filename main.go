// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.


package main

import (
	"flag"
	"fmt"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "cds.ikigai.net/cabinet.v1/cds_go"
)

var (
	port = flag.Int("port", 8888, "The server port")
)

type CDSCabinetServer struct{
	version int32
	fdb fdb.Transactor
}

func newCDSServer() *CDSCabinetServer {
	// Different API versions may expose different runtime behaviors.
	fdb.MustAPIVersion(600)

	// Open the default database from the system cluster
	db := fdb.MustOpenDefault()

	s := &CDSCabinetServer{
		version:1,
		fdb:db,
	}

	return s
}

func main(){
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))

	if err != nil {
		log.Fatalf("Failed to bind on :%d: %v", *port, err)
	}

	log.Printf("I am running! on port %d", *port)

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	pb.RegisterCDSCabinetServer(grpcServer, newCDSServer())
	err = grpcServer.Serve(lis)

	if err != nil{
		log.Fatalf("Error during gRPC execution: %v", err)
	}
}
