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
	"github.com/segmentio/ksuid"
)

type BaseCounter interface {
	GetPath() string
	GetKey(dbCnt subspace.Subspace, cntGroup string) fdb.Key
	GetKeyRange(dbCnt subspace.Subspace) fdb.ExactRange

	ValidateIRI() error
	ValidatePermission() error
}

/* Edges */
type EdgeCounter struct{
	BaseCounter

	Subject string
	Predicate uint16
	Target string

	subjectKSUID ksuid.KSUID
	targetKSUID ksuid.KSUID

	Counter uint16
}

func (c *EdgeCounter) GetPath() string{
	return fmt.Sprintf("/c/e/%d/%s/%d/%s", c.Counter, c.Subject, c.Predicate, c.Target)
}

func (c *EdgeCounter) GetCounterK() string{
	return intToKeyElement(c.Counter)
}

func (c *EdgeCounter) GetPredicateK() string{
	return intToKeyElement(c.Predicate)
}

func (c *EdgeCounter) GetKey(dbCnt subspace.Subspace, cntGroup string) fdb.Key{
	return dbCnt.Sub("e").Pack(tuple.Tuple{c.GetCounterK(), c.Subject, c.GetPredicateK(), c.Target, cntGroup})
}

func (c *EdgeCounter) GetKeyRange(dbCnt subspace.Subspace) fdb.ExactRange{
	return fdb.KeyRange{
		Begin: dbCnt.Sub("e").Pack(tuple.Tuple{c.GetCounterK(), c.Subject, c.GetPredicateK(), c.Target, "0"}),
		End: dbCnt.Sub("e").Pack(tuple.Tuple{c.GetCounterK(), c.Subject, c.GetPredicateK(), c.Target, "f"}),
	}
}

func (e *EdgeCounter) ValidateIRI() error{
	var err error

	if !validateSequence(e.Counter){
		return &ParsingError{msg: "null record", field: "counter.counter"}
	}else if !validateSequence(e.Predicate){
		return &ParsingError{msg: "null record", field: "counter.edge.predicate"}
	}

	if e.subjectKSUID, err = validateNodeID(e.Subject); err != nil{
		return &ParsingError{msg: "invalid Node ID", field: "counter.edge.subject"}
	}

	if e.targetKSUID, err = validateNodeID(e.Target); err != nil{
		return &ParsingError{msg: "invalid Node ID", field: "counter.edge.target"}
	}

	return nil
}

func (e *EdgeCounter) ValidatePermission() error{
	return nil
}


/* Nodes */
type NodeCounter struct{
	BaseCounter
	Counter uint16

	Node string
	nodeKSUID ksuid.KSUID
}

func (c *NodeCounter) GetPath() string{
	return fmt.Sprintf("/c/n/%d/%s", c.Counter, c.Node)
}

func (c *NodeCounter) getCounterK() string{
	return intToKeyElement(c.Counter)
}

func (c *NodeCounter) GetKey(dbCnt subspace.Subspace, cntGroup string) fdb.Key{
	return dbCnt.Sub("n").Pack(tuple.Tuple{c.getCounterK(), c.Node, cntGroup})
}

func (c *NodeCounter) GetKeyRange(dbCnt subspace.Subspace) fdb.ExactRange{
	return fdb.ExactRange(fdb.KeyRange{
		Begin: dbCnt.Sub("n").Pack(tuple.Tuple{c.getCounterK(), c.Node, "0"}),
		End: dbCnt.Sub("n").Pack(tuple.Tuple{c.getCounterK(), c.Node, "f"}),
	})
}

func (c *NodeCounter) ValidateIRI() error{
	var err error

	if !validateSequence(c.Counter) {
		return &ParsingError{msg: "null record", field: "counter.counter"}
	}

	if c.nodeKSUID, err = validateNodeID(c.Node); err != nil{
		return &ParsingError{msg: "invalid Node ID", field: "counter.node"}
	}

	return nil
}

func (e *NodeCounter) ValidatePermission() error{
	return nil
}
