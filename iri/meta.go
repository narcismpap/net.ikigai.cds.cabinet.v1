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
type EdgeMeta struct{
	IRI

	Subject string
	Predicate uint16
	Target string

	Property uint16
}

func (m *EdgeMeta) GetPath() string{
	return fmt.Sprintf("/m/e/%s/%d/%s/%d", m.Subject, m.Predicate, m.Target, m.Property)
}

func (m *EdgeMeta) getPropertyK() string{
	return intToKeyElement(m.Property)
}

func (m *EdgeMeta) getPredicateK() string{
	return intToKeyElement(m.Predicate)
}

func (m *EdgeMeta) GetKey(db subspace.Subspace) fdb.Key{
	return db.Sub("e").Pack(tuple.Tuple{m.Subject, m.getPredicateK(), m.Target, m.getPropertyK()})
}

func (m *EdgeMeta) GetClearRange(db subspace.Subspace) fdb.ExactRange{
	if m.Target != ""{
		return db.Sub("e").Sub(m.Subject).Sub(m.getPredicateK()).Sub(m.Target)
	}

	if m.Predicate > 0{
		return db.Sub("e").Sub(m.Subject).Sub(m.Predicate)
	}

	return db.Sub("e").Sub(m.Subject)
}

func (e *EdgeMeta) ValidateIRI() error{
	return nil
}

func (e *EdgeMeta) ValidatePermission() error{
	return nil
}


/* Nodes */
type NodeMeta struct{
	Node string
	Property uint16
}

func (m *NodeMeta) GetPath() string{
	return fmt.Sprintf("/m/n/%s/%d", m.Node, m.Property)
}

func (m *NodeMeta) getPropertyK() string{
	return intToKeyElement(m.Property)
}

func (m *NodeMeta) GetKey(db subspace.Subspace) fdb.Key{
	return db.Sub("n").Pack(tuple.Tuple{m.Node, m.getPropertyK()})
}

func (m *NodeMeta) GetClearRange(db subspace.Subspace) fdb.ExactRange{
	return db.Sub("n").Sub(m.Node)
}

func (e *NodeMeta) ValidateIRI() error{
	return nil
}

func (e *NodeMeta) ValidatePermission() error{
	return nil
}
