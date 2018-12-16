// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package iri

import (
	pb "cds.ikigai.net/cabinet.v1/rpc"
	"log"
	"strings"
)

func ResolveMetaIRI(tMeta *pb.Meta, nMap *map[string]string) (IRI, error){
	switch mType := tMeta.Object.(type) {

	case *pb.Meta_Edge:
		meta := &EdgeMeta{
			Property: 	uint16(tMeta.Key),
			Subject: 	NodeResolveId(mType.Edge.Subject, nMap),
			Predicate: 	uint16(mType.Edge.Predicate),
			Target: 	NodeResolveId(mType.Edge.Target, nMap),
		}

		return meta, meta.ValidateIRI()

	case *pb.Meta_Node:
		meta := &NodeMeta{
			Property: 	uint16(tMeta.Key),
			Node: 		NodeResolveId(mType.Node, nMap),
		}

		return meta, meta.ValidateIRI()

	default:
		return nil, &ParsingError{msg: "unimplemented object type", field: "meta.object"}
	}
}

func ResolveCounterIRI(tCounter *pb.Counter, nMap *map[string]string) (BaseCounter, error){
	switch cType := tCounter.Object.(type) {

	case *pb.Counter_Edge:
		cnt := &EdgeCounter{
			Counter:   uint16(tCounter.Counter),
			Subject:   NodeResolveId(cType.Edge.Subject, nMap),
			Predicate: uint16(cType.Edge.Predicate),
			Target:    NodeResolveId(cType.Edge.Target, nMap),
		}

		return cnt, cnt.ValidateIRI()

	case *pb.Counter_Node:
		cnt := &NodeCounter{
			Counter: uint16(tCounter.Counter),
			Node:    NodeResolveId(cType.Node, nMap),
		}

		return cnt, cnt.ValidateIRI()

	default:
		return nil, &ParsingError{msg: "unimplemented object type", field: "counter.object"}
	}
}


func NodeResolveId(nID string, nMap *map[string]string) string{
	if strings.HasPrefix(nID, "tmp:") {
		if val, ok := (*nMap)[strings.TrimLeft(nID, "tmp:")]; ok {
			return val
		}

		log.Panicf("unable to map %s in %v", nID, *nMap)
	}

	return nID
}

