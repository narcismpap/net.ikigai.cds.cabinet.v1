// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package main

import(
	pb "cds.ikigai.net/cabinet.v1/cds_go"
	"context"
)

func (s *CDSCabinetServer) CounterGet(ctx context.Context, counter *pb.Counter) (*pb.CounterValueResponse, error) {
	//val := pb.CounterValueResponse{Value:1}
	//return &val, nil

	return nil, nil
}

func (s *CDSCabinetServer) EdgeGet(ctx context.Context, edge *pb.EdgeGetRequest) (*pb.Edge, error){
	return nil, nil
}

func (s *CDSCabinetServer) EdgeList(edgeRq *pb.EdgeListRequest, stream pb.CDSCabinet_EdgeListServer) error{
	return nil
}


func (s *CDSCabinetServer) IndexGet(ctx context.Context, indexGet *pb.IndexGetRequest) (*pb.Index, error){
	return nil, nil
}


func (s *CDSCabinetServer) IndexList(indexRq *pb.IndexListRequest, stream pb.CDSCabinet_IndexListServer) error{
	return nil
}


func (s *CDSCabinetServer) MetaGet(ctx context.Context, meta *pb.Meta) (*pb.MetaGetResponse, error){
	return nil, nil
}

func (s *CDSCabinetServer) MetaList(metaRq *pb.MetaListRequest, stream pb.CDSCabinet_MetaListServer) error{
	return nil
}

func (s *CDSCabinetServer) NodeGet(ctx context.Context, noteRq *pb.NodeGetRequest) (*pb.Node, error){
	return nil, nil
}

func (s *CDSCabinetServer) NodeList(nodeRq *pb.NodeListRequest, stream pb.CDSCabinet_NodeListServer) error{
	return nil
}

func (s *CDSCabinetServer) Transaction(bStream pb.CDSCabinet_TransactionServer) error{
	return nil
}

func (s *CDSCabinetServer) ReadCheck(ctx context.Context, readRq *pb.ReadCheckRequest) (*pb.ReadCheckResponse, error){
	return nil, nil
}

func (s *CDSCabinetServer) SequentialCreate(ctx context.Context, seq *pb.Sequential) (*pb.Sequential, error){
	return nil, nil
}

func (s *CDSCabinetServer) SequentialUpdate(ctx context.Context, seq *pb.Sequential) (*pb.MutationResponse, error){
	return nil, nil
}

func (s *CDSCabinetServer) SequentialDelete(ctx context.Context, seq *pb.Sequential) (*pb.MutationResponse, error){
	return nil, nil
}

func (s *CDSCabinetServer) SequentialGet(ctx context.Context, seq *pb.Sequential) (*pb.Sequential, error){
	return nil, nil
}

func (s *CDSCabinetServer) SequentialList(seq *pb.SequentialListRequest, stream pb.CDSCabinet_SequentialListServer) error{
	return nil
}
