// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.


package main

import (
	pb "cds.ikigai.net/cabinet.v1/rpc"
	"errors"
)

func resolveMetaIRI(tMeta *pb.Meta) (IRI, error){
	switch mType := tMeta.Object.(type) {

	case *pb.Meta_Edge:
		return (&IRIEdgeMeta{
			Property: 	uint16(tMeta.Key),
			Subject: 	mType.Edge.Subject,
			Predicate: 	uint16(mType.Edge.Predicate),
			Target: 	mType.Edge.Target,
		}), nil

	case *pb.Meta_Node:
		return (&IRINodeMeta{
			Property: 	uint16(tMeta.Key),
			Node: 		mType.Node,
		}), nil

	default:
		return nil, errors.New("bad mType")
	}
}

func resolveCounterIRI(tCounter *pb.Counter) (IRICounter, error){
	switch cType := tCounter.Object.(type) {

	case *pb.Counter_Edge:
		return (&IRIEdgeCounter{
			Counter:   uint16(tCounter.Counter),
			Subject:   cType.Edge.Subject,
			Predicate: uint16(cType.Edge.Predicate),
			Target:    cType.Edge.Target,
		}), nil

	case *pb.Counter_Node:
		return (&IRINodeCounter{
			Counter: uint16(tCounter.Counter),
			Node:    cType.Node,
		}), nil

	default:
		return nil, errors.New("bad cType")
	}
}
