// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package iri

import (
	"cds.ikigai.net/cabinet.v1/perms"
	"fmt"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
	"github.com/segmentio/ksuid"
	"strings"
)

type BaseCounter interface {
	GetPath() string
	GetKey(dbCnt subspace.Subspace, cntGroup string) fdb.Key
	GetKeyRange(dbCnt subspace.Subspace) fdb.ExactRange
}

/* Edges */
type EdgeCounter struct {
	BaseCounter

	Subject   string
	Predicate uint16
	Target    string

	subjectKSUID ksuid.KSUID
	targetKSUID  ksuid.KSUID

	Counter uint16
}

func (c *EdgeCounter) GetPath() string {
	return fmt.Sprintf("/c/e/%d/%s/%d/%s", c.Counter, c.Subject, c.Predicate, c.Target)
}

func (c *EdgeCounter) Parse(path string) error {
	parts := strings.Split(path, "/") // c/e/{COUNTER}/{SUBJECT}/{PREDICATE}/{TARGET}
	var err error

	if c.Counter, err = ParseCoreSequence(parts[2]); err != nil {
		return &ParsingError{msg: "invalid counter", field: "counter.counter"}
	}

	if c.Predicate, err = ParseCoreSequence(parts[4]); err != nil {
		return &ParsingError{msg: "invalid predicate", field: "counter.edge.predicate"}
	}

	c.Subject = parts[3]
	c.Target = parts[5]

	return nil
}

func (c *EdgeCounter) getCounterK() []byte {
	return SequenceToSmallKey(c.Counter)
}

func (c *EdgeCounter) getPredicateK() []byte {
	return SequenceToSmallKey(c.Predicate)
}

func (c *EdgeCounter) GetKey(dbCnt subspace.Subspace, cntGroup string) fdb.Key {
	return dbCnt.Sub("e").Pack(tuple.Tuple{c.getCounterK(), c.Subject, c.getPredicateK(), c.Target, cntGroup})
}

func (c *EdgeCounter) GetKeyRange(dbCnt subspace.Subspace) fdb.ExactRange {
	return fdb.KeyRange{
		Begin: dbCnt.Sub("e").Pack(tuple.Tuple{c.getCounterK(), c.Subject, c.getPredicateK(), c.Target, "0"}),
		End:   dbCnt.Sub("e").Pack(tuple.Tuple{c.getCounterK(), c.Subject, c.getPredicateK(), c.Target, "f"}),
	}
}

func (c *EdgeCounter) ValidateIRI(p *perms.Count) error {
	var err error

	if !validateSequence(c.Counter) {
		return &ParsingError{msg: "null record", field: "counter.counter"}
	} else if !validateSequence(c.Predicate) {
		return &ParsingError{msg: "null record", field: "counter.edge.predicate"}
	}

	if c.subjectKSUID, err = validateNodeID(c.Subject); err != nil {
		return &ParsingError{msg: "invalid Node ID", field: "counter.edge.subject"}
	}

	if c.targetKSUID, err = validateNodeID(c.Target); err != nil {
		return &ParsingError{msg: "invalid Node ID", field: "counter.edge.target"}
	}

	return nil
}

func (c *EdgeCounter) ValidatePermission(p perms.Count) error {
	return nil
}

/* Nodes */
type NodeCounter struct {
	BaseCounter
	Counter uint16

	Node      string
	nodeKSUID ksuid.KSUID
}

func (c *NodeCounter) GetPath() string {
	return fmt.Sprintf("/c/n/%d/%s", c.Counter, c.Node)
}

func (c *NodeCounter) Parse(path string) error {
	parts := strings.Split(path, "/") // c/n/{COUNTER}/{NODE}
	var err error

	if c.Counter, err = ParseCoreSequence(parts[2]); err != nil {
		return &ParsingError{msg: "invalid counter", field: "counter.counter"}
	}

	c.Node = parts[3]
	return nil
}

func (c *NodeCounter) getCounterK() []byte {
	return SequenceToSmallKey(c.Counter)
}

func (c *NodeCounter) GetKey(dbCnt subspace.Subspace, cntGroup string) fdb.Key {
	return dbCnt.Sub("n").Pack(tuple.Tuple{c.getCounterK(), c.Node, cntGroup})
}

func (c *NodeCounter) GetKeyRange(dbCnt subspace.Subspace) fdb.ExactRange {
	return fdb.ExactRange(fdb.KeyRange{
		Begin: dbCnt.Sub("n").Pack(tuple.Tuple{c.getCounterK(), c.Node, "0"}),
		End:   dbCnt.Sub("n").Pack(tuple.Tuple{c.getCounterK(), c.Node, "f"}),
	})
}

func (c *NodeCounter) ValidateIRI(p *perms.Count) error {
	var err error

	if !validateSequence(c.Counter) {
		return &ParsingError{msg: "null record", field: "counter.counter"}
	}

	if c.nodeKSUID, err = validateNodeID(c.Node); err != nil {
		return &ParsingError{msg: "invalid Node ID", field: "counter.node"}
	}

	return nil
}

func (c *NodeCounter) ValidatePermission(p perms.Count) error {
	return nil
}
