// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package main

import (
	pb "cds.ikigai.net/cabinet.v1/rpc"
	"errors"
	"log"
	"strings"
)

func resolveMetaIRI(tMeta *pb.Meta, nMap *map[string]string) (IRI, error){
	switch mType := tMeta.Object.(type) {

	case *pb.Meta_Edge:
		return (&IRIEdgeMeta{
			Property: 	uint16(tMeta.Key),
			Subject: 	nodeResolveId(mType.Edge.Subject, nMap),
			Predicate: 	uint16(mType.Edge.Predicate),
			Target: 	nodeResolveId(mType.Edge.Target, nMap),
		}), nil

	case *pb.Meta_Node:
		return (&IRINodeMeta{
			Property: 	uint16(tMeta.Key),
			Node: 		nodeResolveId(mType.Node, nMap),
		}), nil

	default:
		return nil, errors.New("bad mType")
	}
}

func resolveCounterIRI(tCounter *pb.Counter, nMap *map[string]string) (IRICounter, error){
	switch cType := tCounter.Object.(type) {

	case *pb.Counter_Edge:
		return (&IRIEdgeCounter{
			Counter:   uint16(tCounter.Counter),
			Subject:   nodeResolveId(cType.Edge.Subject, nMap),
			Predicate: uint16(cType.Edge.Predicate),
			Target:    nodeResolveId(cType.Edge.Target, nMap),
		}), nil

	case *pb.Counter_Node:
		return (&IRINodeCounter{
			Counter: uint16(tCounter.Counter),
			Node:    nodeResolveId(cType.Node, nMap),
		}), nil

	default:
		return nil, errors.New("bad cType")
	}
}


func nodeResolveId(nID string, nMap *map[string]string) string{
	if strings.HasPrefix(nID, "tmp:") {
		if val, ok := (*nMap)[strings.TrimLeft(nID, "tmp:")]; ok {
			return val
		}

		log.Panicf("unable to map %s in %v", nID, *nMap)
	}

	return nID
}

