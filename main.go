package main

import (
	pb "./cabinet/cabinet.pb.go"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port       = flag.Int("port", 8888, "The server port")
)

type CDSCabinetServer struct{

}

func newServer() *CabinetServiceServer {
	s := &CabinetServiceServer{}
	return s
}

func main(){
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	pb.RegisterRouteGuideServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}