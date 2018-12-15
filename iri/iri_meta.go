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

/* Edges*/
type IRIEdgeMeta struct{
	IRI

	Subject string
	Predicate uint16
	Target string

	Property uint16
}

func (m *IRIEdgeMeta) GetPath() string{
	return fmt.Sprintf("/m/e/%s/%d/%s/%d", m.Subject, m.Predicate, m.Target, m.Property)
}

func (m *IRIEdgeMeta) getPropertyK() string{
	return intToKeyElement(m.Property)
}

func (m *IRIEdgeMeta) getPredicateK() string{
	return intToKeyElement(m.Predicate)
}

func (m *IRIEdgeMeta) GetKey(db subspace.Subspace) fdb.Key{
	return db.Sub("e").Pack(tuple.Tuple{m.Subject, m.getPredicateK(), m.Target, m.getPropertyK()})
}

func (m *IRIEdgeMeta) GetClearRange(db subspace.Subspace) fdb.ExactRange{
	if m.Target != ""{
		return db.Sub("e").Sub(m.Subject).Sub(m.getPredicateK()).Sub(m.Target)
	}

	if m.Predicate > 0{
		return db.Sub("e").Sub(m.Subject).Sub(m.Predicate)
	}

	return db.Sub("e").Sub(m.Subject)
}

func (e *IRIEdgeMeta) ValidateIRI() error{
	return nil
}

func (e *IRIEdgeMeta) ValidatePermission() error{
	return nil
}


/* Nodes */
type IRINodeMeta struct{
	Node string
	Property uint16
}

func (m *IRINodeMeta) GetPath() string{
	return fmt.Sprintf("/m/n/%s/%d", m.Node, m.Property)
}

func (m *IRINodeMeta) getPropertyK() string{
	return intToKeyElement(m.Property)
}

func (m *IRINodeMeta) GetKey(db subspace.Subspace) fdb.Key{
	return db.Sub("n").Pack(tuple.Tuple{m.Node, m.getPropertyK()})
}

func (m *IRINodeMeta) GetClearRange(db subspace.Subspace) fdb.ExactRange{
	return db.Sub("n").Sub(m.Node)
}

func (e *IRINodeMeta) ValidateIRI() error{
	return nil
}

func (e *IRINodeMeta) ValidatePermission() error{
	return nil
}
