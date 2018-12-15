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

func (m *IRIEdgeMeta) GetKey(server *CDSCabinetServer) fdb.Key{
	return server.dbMeta.Sub("e").Pack(tuple.Tuple{m.Subject, m.getPredicateK(), m.Target, m.getPropertyK()})
}

func (m *IRIEdgeMeta) GetClearRange(server *CDSCabinetServer) fdb.ExactRange{
	if m.Target != ""{
		return server.dbMeta.Sub("e").Sub(m.Subject).Sub(m.getPredicateK()).Sub(m.Target)
	}

	if m.Predicate > 0{
		return server.dbMeta.Sub("e").Sub(m.Subject).Sub(m.Predicate)
	}

	return server.dbMeta.Sub("e").Sub(m.Subject)
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

func (m *IRINodeMeta) GetKey(server *CDSCabinetServer) fdb.Key{
	return server.dbMeta.Sub("n").Pack(tuple.Tuple{m.Node, m.getPropertyK()})
}

func (m *IRINodeMeta) GetClearRange(server *CDSCabinetServer) fdb.ExactRange{
	return server.dbMeta.Sub("n").Sub(m.Node)
}

func (e *IRINodeMeta) ValidateIRI() error{
	return nil
}

func (e *IRINodeMeta) ValidatePermission() error{
	return nil
}