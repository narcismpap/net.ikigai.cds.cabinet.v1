// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package iri

import (
	"cds.ikigai.net/cabinet.v1/perms"
	pb "cds.ikigai.net/cabinet.v1/rpc"
	"errors"
	"strings"
)

func ResolveMetaIRI(tMeta *pb.Meta, nMap *map[string]string, p *perms.Meta) (IRI, error) {
	switch mType := tMeta.Object.(type) {

	case *pb.Meta_Edge:
		subjectID, err := NodeResolveId(mType.Edge.Subject, nMap)
		if err != nil {
			return nil, &ParsingError{msg: "tmp:X is invalid", field: "meta.edge.subject"}
		}

		targetID, err := NodeResolveId(mType.Edge.Target, nMap)
		if err != nil {
			return nil, &ParsingError{msg: "tmp:X is invalid", field: "meta.edge.target"}
		}

		meta := &EdgeMeta{
			Property:  uint16(tMeta.Key),
			Subject:   subjectID,
			Predicate: uint16(mType.Edge.Predicate),
			Target:    targetID,
		}

		return meta, meta.ValidateIRI(p)

	case *pb.Meta_Node:
		nID, err := NodeResolveId(mType.Node, nMap)
		if err != nil {
			return nil, &ParsingError{msg: "tmp:X is invalid", field: "meta.node"}
		}

		meta := &NodeMeta{
			Property: uint16(tMeta.Key),
			Node:     nID,
		}

		return meta, meta.ValidateIRI(p)

	default:
		return nil, &ParsingError{msg: "unimplemented object type", field: "meta.object"}
	}
}

func ResolveCounterIRI(tCounter *pb.Counter, nMap *map[string]string, p *perms.Count) (BaseCounter, error) {
	switch cType := tCounter.Object.(type) {

	case *pb.Counter_Edge:
		subjectID, err := NodeResolveId(cType.Edge.Subject, nMap)
		if err != nil {
			return nil, &ParsingError{msg: "tmp:X is invalid", field: "counter.edge.subject"}
		}

		targetID, err := NodeResolveId(cType.Edge.Target, nMap)
		if err != nil {
			return nil, &ParsingError{msg: "tmp:X is invalid", field: "counter.edge.target"}
		}

		cnt := &EdgeCounter{
			Counter:   uint16(tCounter.Counter),
			Subject:   subjectID,
			Predicate: uint16(cType.Edge.Predicate),
			Target:    targetID,
		}

		return cnt, cnt.ValidateIRI(p)

	case *pb.Counter_Node:
		nID, err := NodeResolveId(cType.Node, nMap)
		if err != nil {
			return nil, &ParsingError{msg: "tmp:X is invalid", field: "counter.node"}
		}

		cnt := &NodeCounter{
			Counter: uint16(tCounter.Counter),
			Node:    nID,
		}

		return cnt, cnt.ValidateIRI(p)

	default:
		return nil, &ParsingError{msg: "unimplemented object type", field: "counter.object"}
	}
}

func NodeResolveId(nID string, nMap *map[string]string) (string, error) {
	if strings.HasPrefix(nID, "tmp:") {
		if val, ok := (*nMap)[strings.TrimLeft(nID, "tmp:")]; ok {
			return val, nil
		}

		return "", errors.New("node map is incomplete")
	}

	return nID, nil
}
