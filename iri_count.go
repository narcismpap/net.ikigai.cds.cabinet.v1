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

type IRICounter interface {
	GetPath() string
	GetKey(server *CDSCabinetServer, cntGroup string) fdb.Key
	GetKeyRange(server *CDSCabinetServer) fdb.ExactRange

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

func (c *IRIEdgeCounter) GetKey(server *CDSCabinetServer, cntGroup string) fdb.Key{
	return server.dbCnt.Sub("e").Pack(tuple.Tuple{c.GetCounterK(), c.Subject, c.GetPredicateK(), c.Target, cntGroup})
}

func (c *IRIEdgeCounter) GetKeyRange(server *CDSCabinetServer) fdb.ExactRange{
	return fdb.KeyRange{
		Begin: server.dbCnt.Sub("e").Pack(tuple.Tuple{c.GetCounterK(), c.Subject, c.GetPredicateK(), c.Target, "0"}),
		End: server.dbCnt.Sub("e").Pack(tuple.Tuple{c.GetCounterK(), c.Subject, c.GetPredicateK(), c.Target, "f"}),
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

func (c *IRINodeCounter) GetKey(server *CDSCabinetServer, cntGroup string) fdb.Key{
	return server.dbCnt.Sub("n").Pack(tuple.Tuple{c.getCounterK(), c.Node, cntGroup})
}

func (c *IRINodeCounter) GetKeyRange(server *CDSCabinetServer) fdb.ExactRange{
	return fdb.ExactRange(fdb.KeyRange{
		Begin: server.dbCnt.Sub("n").Pack(tuple.Tuple{c.getCounterK(), c.Node, "0"}),
		End: server.dbCnt.Sub("n").Pack(tuple.Tuple{c.getCounterK(), c.Node, "f"}),
	})
}

func (e *IRINodeCounter) ValidateIRI() error{
	return nil
}

func (e *IRINodeCounter) ValidatePermission() error{
	return nil
}
