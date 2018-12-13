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

type IRI interface {
	getPath() string
	getKey(server *CDSCabinetServer) fdb.Key
}

type IRISequential struct{
	IRI
	Type string
	SeqID uint32
}

func (s *IRISequential) dbSeqID() string{
	return fmt.Sprintf("%05d", s.SeqID)
}

func (s *IRISequential) getPath() string{
	return fmt.Sprintf("/s/%s/%d", s.Type, s.SeqID)
}

func (s *IRISequential) getKey(server *CDSCabinetServer) fdb.Key{
	return server.dbSeq.Pack(tuple.Tuple{s.Type, s.dbSeqID()})
}

func (s *IRISequential) getIncrementKey(server *CDSCabinetServer) fdb.Key{
	return server.dbSeq.Pack(tuple.Tuple{"l", s.Type})
}


type IRINode struct{
	IRI
	Type uint16
	Id string
	Property int
}

func (n *IRINode) getPath() string{
	return fmt.Sprintf("/n/%d/%s", n.Type, n.Id)
}

func (n *IRINode) getPathProperty(prop int) string{
	return fmt.Sprintf("/n/%d/%s/p/%d", n.Type, n.Id, prop)
}

func (n *IRINode) getKey(server *CDSCabinetServer) fdb.Key{
	return server.dbNode.Sub(n.Type).Pack(tuple.Tuple{[]byte(n.Id)})
}

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


type IRINodeIndex struct{
	IRI

	Node string
	IndexId uint16
	Value string
}

func (i *IRINodeIndex) getPath() string{
	return fmt.Sprintf("/i/%d/%s/%s", i.IndexId, i.Value, i.Node)
}

func (i *IRINodeIndex) getKey(server *CDSCabinetServer) fdb.Key{
	return server.dbIndex.Pack(tuple.Tuple{i.IndexId, i.Value, i.Node})
}

type IRINodeMeta struct{
	Node string
	Property uint16
}

func (m *IRINodeMeta) getPath() string{
	return fmt.Sprintf("/m/n/%s/%d", m.Node, m.Property)
}

func (m *IRINodeMeta) getKey(server *CDSCabinetServer) fdb.Key{
	return server.dbMeta.Sub("n").Pack(tuple.Tuple{m.Node, m.Property})
}


type IRIEdgeMeta struct{
	IRI

	Subject string
	Predicate uint16
	Target string

	Property uint16
}

func (m *IRIEdgeMeta) getPath() string{
	return fmt.Sprintf("/m/e/%s/%d/%s/%d", m.Subject, m.Predicate, m.Target, m.Property)
}

func (m *IRIEdgeMeta) getKey(server *CDSCabinetServer) fdb.Key{
	return server.dbMeta.Sub("e").Pack(tuple.Tuple{m.Subject, m.Predicate, m.Target, m.Property})
}


type IRICounter interface {
	getPath() string
	getKey(server *CDSCabinetServer, cntGroup string) fdb.Key
	getKeyRange(server *CDSCabinetServer) fdb.ExactRange
}

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

func (c *IRIEdgeCounter) getKey(server *CDSCabinetServer, cntGroup string) fdb.Key{
	return server.dbCnt.Sub("e").Pack(tuple.Tuple{c.Counter, c.Subject, c.Predicate, c.Target, cntGroup})
}

func (c *IRIEdgeCounter) getKeyRange(server *CDSCabinetServer) fdb.ExactRange{
	return fdb.KeyRange{
		Begin: server.dbCnt.Sub("e").Pack(tuple.Tuple{c.Counter, c.Subject, c.Predicate, c.Target, "0"}),
		End: server.dbCnt.Sub("e").Pack(tuple.Tuple{c.Counter, c.Subject, c.Predicate, c.Target, "f"}),
	}
}

type IRINodeCounter struct{
	IRICounter
	Counter uint16

	Node string
}

func (c *IRINodeCounter) getPath() string{
	return fmt.Sprintf("/c/n/%d/%s", c.Counter, c.Node)
}

func (c *IRINodeCounter) getKey(server *CDSCabinetServer, cntGroup string) fdb.Key{
	return server.dbCnt.Sub("n").Pack(tuple.Tuple{c.Counter, c.Node, cntGroup})
}

func (c *IRINodeCounter) getKeyRange(server *CDSCabinetServer) fdb.ExactRange{
	return fdb.ExactRange(fdb.KeyRange{
		Begin: server.dbCnt.Sub("n").Pack(tuple.Tuple{c.Counter, c.Node, "0"}),
		End: server.dbCnt.Sub("n").Pack(tuple.Tuple{c.Counter, c.Node, "f"}),
	})
}