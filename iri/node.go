// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package iri

import (
	"cds.ikigai.net/cabinet.v1/perms"
	pb "cds.ikigai.net/cabinet.v1/rpc"
	"fmt"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
	"github.com/segmentio/ksuid"
	"strings"
)

type Node struct {
	IRI
	Type uint16
	Id   string

	nodeKSUID ksuid.KSUID
}

func (n *Node) getTypeK() []byte {
	return SequenceToSmallKey(n.Type)
}

func (n *Node) GetPath() string {
	return fmt.Sprintf("n/%d/%s", n.Type, n.Id)
}

func (n *Node) Parse(path string) error {
	parts := strings.Split(path, "/") // n/{TYPE}/{ID}
	var err error

	if n.Type, err = ParseCoreSequence(parts[1]); err != nil {
		return &ParsingError{msg: "invalid type", field: "node.type"}
	}

	n.Id = parts[2]
	return nil
}

func (n *Node) GetKey(db subspace.Subspace) fdb.Key {
	return db.Sub(n.getTypeK()).Pack(tuple.Tuple{n.Id})
}

func (n *Node) GetClearRange(db subspace.Subspace) fdb.ExactRange {
	if n.Id == "" {
		return db.Sub(n.getTypeK())
	} else {
		return db.Sub(n.getTypeK()).Sub(n.Id)
	}
}

func (n *Node) GetListRange(db subspace.Subspace, rtr fdb.ReadTransaction, opt *pb.ListOptions) fdb.RangeResult {
	readRange := db.Sub(n.getTypeK())

	return rtr.GetRange(readRange, fdb.RangeOptions{
		Limit:   int(opt.PageSize),
		Reverse: opt.Reverse,
	})
}

func (n *Node) ValidateIRI(p *perms.Node) error {
	var err error

	if !validateSequence(n.Type) {
		return &ParsingError{msg: "null record", field: "node.type"}
	}

	if n.nodeKSUID, err = validateNodeID(n.Id); err != nil {
		return &ParsingError{msg: "invalid Node ID", field: "node.id"}
	}

	return nil
}

func (n *Node) ValidatePermission(p perms.Node) error {
	return nil
}
