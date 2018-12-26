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

type NodeIndex struct {
	IRI

	Node    string
	IndexId uint16
	Value   string

	nodeKSUID ksuid.KSUID
}

func (i *NodeIndex) GetPath() string {
	return fmt.Sprintf("i/%d/%s/%s", i.IndexId, i.Value, i.Node)
}

func (i *NodeIndex) Parse(path string) error {
	parts := strings.Split(path, "/") // i/{INDEX}/{VAL}/{NODE}
	var err error

	if i.IndexId, err = ParseCoreSequence(parts[1]); err != nil {
		return &ParsingError{msg: "invalid index", field: "index.index"}
	}

	i.Value = parts[2]
	i.Node = parts[3]

	return nil
}

func (i *NodeIndex) getIndexK() []byte {
	return SequenceToSmallKey(i.IndexId)
}

func (i *NodeIndex) GetKey(db subspace.Subspace) fdb.Key {
	return db.Pack(tuple.Tuple{i.getIndexK(), i.Value, i.Node})
}

func (i *NodeIndex) GetClearRange(db subspace.Subspace) fdb.ExactRange {
	return nil
}

func (i *NodeIndex) GetListRange(db subspace.Subspace, rtr fdb.ReadTransaction, opt *pb.ListOptions) fdb.RangeResult {
	readRange := db.Sub(i.getIndexK())

	if i.Value != "*" {
		readRange = readRange.Sub(i.Value)
	}

	return rtr.GetRange(readRange, fdb.RangeOptions{
		Limit:   int(opt.PageSize),
		Reverse: opt.Reverse,
	})
}

func (i *NodeIndex) ValidateIRI(p *perms.Index) error {
	var err error

	if !validateSequence(i.IndexId) {
		return &ParsingError{msg: "null record", field: "index.IndexId"}
	}

	if len(i.Value) == 0 {
		return &ParsingError{msg: "null record", field: "index.value"}
	}else if len(i.Value) > 256 {
		return &ParsingError{msg: "len > 256", field: "index.value"}
	}

	if i.nodeKSUID, err = validateNodeID(i.Node); err != nil {
		return &ParsingError{msg: "invalid Node ID", field: "index.node"}
	}

	return nil
}

func (i *NodeIndex) ValidatePermission(p perms.Index) error {
	return nil
}
