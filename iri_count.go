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
	getPath() string
	getKey(server *CDSCabinetServer, cntGroup string) fdb.Key
	getKeyRange(server *CDSCabinetServer) fdb.ExactRange
}

/* Edges */
type IRIEdgeCounter struct{
	IRICounter

	Subject string
	Predicate uint16
	Target string

	Counter uint16
}

func (c *IRIEdgeCounter) getPath() string{
	return fmt.Sprintf("/c/e/%d/%s/%d/%s", c.Counter, c.Subject, c.Predicate, c.Target)
}

func (c *IRIEdgeCounter) getCounterK() string{
	return intToKeyElement(c.Counter)
}

func (c *IRIEdgeCounter) getPredicateK() string{
	return intToKeyElement(c.Predicate)
}

func (c *IRIEdgeCounter) getKey(server *CDSCabinetServer, cntGroup string) fdb.Key{
	return server.dbCnt.Sub("e").Pack(tuple.Tuple{c.getCounterK(), c.Subject, c.getPredicateK(), c.Target, cntGroup})
}

func (c *IRIEdgeCounter) getKeyRange(server *CDSCabinetServer) fdb.ExactRange{
	return fdb.KeyRange{
		Begin: server.dbCnt.Sub("e").Pack(tuple.Tuple{c.getCounterK(), c.Subject, c.getPredicateK(), c.Target, "0"}),
		End: server.dbCnt.Sub("e").Pack(tuple.Tuple{c.getCounterK(), c.Subject, c.getPredicateK(), c.Target, "f"}),
	}
}

/* Nodes */
type IRINodeCounter struct{
	IRICounter
	Counter uint16

	Node string
}

func (c *IRINodeCounter) getPath() string{
	return fmt.Sprintf("/c/n/%d/%s", c.Counter, c.Node)
}

func (c *IRINodeCounter) getCounterK() string{
	return intToKeyElement(c.Counter)
}

func (c *IRINodeCounter) getKey(server *CDSCabinetServer, cntGroup string) fdb.Key{
	return server.dbCnt.Sub("n").Pack(tuple.Tuple{c.getCounterK(), c.Node, cntGroup})
}

func (c *IRINodeCounter) getKeyRange(server *CDSCabinetServer) fdb.ExactRange{
	return fdb.ExactRange(fdb.KeyRange{
		Begin: server.dbCnt.Sub("n").Pack(tuple.Tuple{c.getCounterK(), c.Node, "0"}),
		End: server.dbCnt.Sub("n").Pack(tuple.Tuple{c.getCounterK(), c.Node, "f"}),
	})
}
