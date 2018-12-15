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

type IRICounter interface {
	GetPath() string
	GetKey(dbCnt subspace.Subspace, cntGroup string) fdb.Key
	GetKeyRange(dbCnt subspace.Subspace) fdb.ExactRange

	ValidateIRI() error
	ValidatePermission() error
}

/* Edges */
type IRIEdgeCounter struct{
	IRICounter

	Subject string
	Predicate uint16
	Target string

	Counter uint16
}

func (c *IRIEdgeCounter) GetPath() string{
	return fmt.Sprintf("/c/e/%d/%s/%d/%s", c.Counter, c.Subject, c.Predicate, c.Target)
}

func (c *IRIEdgeCounter) GetCounterK() string{
	return intToKeyElement(c.Counter)
}

func (c *IRIEdgeCounter) GetPredicateK() string{
	return intToKeyElement(c.Predicate)
}

func (c *IRIEdgeCounter) GetKey(dbCnt subspace.Subspace, cntGroup string) fdb.Key{
	return dbCnt.Sub("e").Pack(tuple.Tuple{c.GetCounterK(), c.Subject, c.GetPredicateK(), c.Target, cntGroup})
}

func (c *IRIEdgeCounter) GetKeyRange(dbCnt subspace.Subspace) fdb.ExactRange{
	return fdb.KeyRange{
		Begin: dbCnt.Sub("e").Pack(tuple.Tuple{c.GetCounterK(), c.Subject, c.GetPredicateK(), c.Target, "0"}),
		End: dbCnt.Sub("e").Pack(tuple.Tuple{c.GetCounterK(), c.Subject, c.GetPredicateK(), c.Target, "f"}),
	}
}

func (e *IRIEdgeCounter) ValidateIRI() error{
	return nil
}

func (e *IRIEdgeCounter) ValidatePermission() error{
	return nil
}


/* Nodes */
type IRINodeCounter struct{
	IRICounter
	Counter uint16

	Node string
}

func (c *IRINodeCounter) GetPath() string{
	return fmt.Sprintf("/c/n/%d/%s", c.Counter, c.Node)
}

func (c *IRINodeCounter) getCounterK() string{
	return intToKeyElement(c.Counter)
}

func (c *IRINodeCounter) GetKey(dbCnt subspace.Subspace, cntGroup string) fdb.Key{
	return dbCnt.Sub("n").Pack(tuple.Tuple{c.getCounterK(), c.Node, cntGroup})
}

func (c *IRINodeCounter) GetKeyRange(dbCnt subspace.Subspace) fdb.ExactRange{
	return fdb.ExactRange(fdb.KeyRange{
		Begin: dbCnt.Sub("n").Pack(tuple.Tuple{c.getCounterK(), c.Node, "0"}),
		End: dbCnt.Sub("n").Pack(tuple.Tuple{c.getCounterK(), c.Node, "f"}),
	})
}

func (e *IRINodeCounter) ValidateIRI() error{
	return nil
}

func (e *IRINodeCounter) ValidatePermission() error{
	return nil
}
