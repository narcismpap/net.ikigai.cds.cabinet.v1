// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package iri

import (
	"fmt"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
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

func (e *IRIEdge) GetKey(db subspace.Subspace) fdb.Key{
	return db.Pack(tuple.Tuple{e.Subject, e.getPredicateK(), e.Target})
}

func (e *IRIEdge) GetClearRange(db subspace.Subspace) fdb.ExactRange{
	if e.Predicate == 0{
		return db.Sub(e.Subject)
	}else{
		return db.Sub(e.Subject).Sub(e.getPredicateK())
	}
}

func (e *IRIEdge) ValidateIRI() error{
	return nil
}

func (e *IRIEdge) ValidatePermission() error{
	return nil
}
