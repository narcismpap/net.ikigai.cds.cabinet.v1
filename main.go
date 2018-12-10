// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.


package main

import (
	"flag"
	"fmt"
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
}

func newCDSServer() *CDSCabinetServer {
	s := &CDSCabinetServer{
		version:1,
	}

	return s
}

func main(){
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))

	if err != nil {
		log.Fatalf("Failed to bind on :%d: %v", port, err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	pb.RegisterCDSCabinetServer(grpcServer, newCDSServer())
	err = grpcServer.Serve(lis)

	if err != nil{
		log.Fatalf("Error during gRPC execution: %v", err)
	}
}
