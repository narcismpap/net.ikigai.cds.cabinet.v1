// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package iri

import (
	cds "cds.ikigai.net/cabinet.v1/server"
	"fmt"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
)

type IRIEdge struct{
	IRI

	Subject string
	Predicate uint16
	Target string
	Property int
}

func (e *IRIEdge) GetPath() string{
	return fmt.Sprintf("/e/%s/%d/%s", e.Subject, e.Predicate, e.Target)
}

func (e *IRIEdge) GetPathProperty(prop int) string{
	return fmt.Sprintf("/e/%s/%d/%s/p/%d", e.Subject, e.Predicate, e.Target, prop)
}

func (e *IRIEdge) getPredicateK() string{
	return intToKeyElement(e.Predicate)
}

func (e *IRIEdge) GetKey(server *cds.CDSCabinetServer) fdb.Key{
	return server.DbEdge.Pack(tuple.Tuple{e.Subject, e.getPredicateK(), e.Target})
}

func (e *IRIEdge) GetClearRange(server *cds.CDSCabinetServer) fdb.ExactRange{
	if e.Predicate == 0{
		return server.DbEdge.Sub(e.Subject)
	}else{
		return server.DbEdge.Sub(e.Subject).Sub(e.getPredicateK())
	}
}

func (e *IRIEdge) ValidateIRI() error{
	return nil
}

func (e *IRIEdge) ValidatePermission() error{
	return nil
}
