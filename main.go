// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package main

import (
	pb "cds.ikigai.net/cabinet.v1/rpc"
	cdsServer "cds.ikigai.net/cabinet.v1/server"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	ServerPort          = 8888
)

func main(){
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", ServerPort))

	if err != nil {
		log.Fatalf("[E] Failed to bind on :%d: %v", ServerPort, err)
	}

	log.Printf("[I] net.ikigai.cds.cabinet.v1 running on port %d", ServerPort)

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	pb.RegisterCDSCabinetServer(grpcServer, cdsServer.StartServer())
	err = grpcServer.Serve(lis)

	if err != nil{
		log.Fatalf("[E] Error during gRPC execution: %v", err)
	}
}
