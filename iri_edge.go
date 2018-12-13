// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package main

import (
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

func (e *IRIEdge) getPath() string{
	return fmt.Sprintf("/e/%s/%d/%s", e.Subject, e.Predicate, e.Target)
}

func (e *IRIEdge) getPathProperty(prop int) string{
	return fmt.Sprintf("/e/%s/%d/%s/p/%d", e.Subject, e.Predicate, e.Target, prop)
}

func (e *IRIEdge) getKey(server *CDSCabinetServer) fdb.Key{
	return server.dbEdge.Pack(tuple.Tuple{e.Subject, e.Predicate, e.Target})
}

func (e *IRIEdge) getClearRange(server *CDSCabinetServer) fdb.ExactRange{
	if e.Predicate == 0{
		return server.dbEdge.Sub(e.Subject)
	}else{
		return server.dbEdge.Sub(e.Subject).Sub(e.Predicate)
	}
}
