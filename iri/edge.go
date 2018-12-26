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

type Edge struct {
	IRI

	Subject   string
	Predicate uint16
	Target    string

	subjectKSUID ksuid.KSUID
	targetKSUID  ksuid.KSUID
}

func (e *Edge) GetPath() string {
	return fmt.Sprintf("e/%s/%d/%s", e.Subject, e.Predicate, e.Target)
}

func (e *Edge) Parse(path string) error {
	parts := strings.Split(path, "/") // e/{SUBJECT}/{PREDICATE}/{TARGET}
	var err error

	if e.Predicate, err = ParseCoreSequence(parts[2]); err != nil {
		return &ParsingError{msg: "invalid predicate", field: "edge.predicate"}
	}

	e.Subject = parts[1]
	e.Target = parts[3]

	return nil
}

func (e *Edge) getPredicateK() []byte {
	return SequenceToSmallKey(e.Predicate)
}

func (e *Edge) GetKey(db subspace.Subspace) fdb.Key {
	return db.Pack(tuple.Tuple{e.Subject, e.getPredicateK(), e.Target})
}

func (e *Edge) GetClearRange(db subspace.Subspace) fdb.ExactRange {
	if e.Predicate == 0 {
		return db.Sub(e.Subject)
	} else {
		return db.Sub(e.Subject).Sub(e.getPredicateK())
	}
}

func (e *Edge) GetListRange(db subspace.Subspace, rtr fdb.ReadTransaction, opt *pb.ListOptions) fdb.RangeResult {
	readRange := db.Sub(e.Subject)

	if e.Predicate > 0 {
		readRange = readRange.Sub(e.getPredicateK())
	}

	return rtr.GetRange(readRange, fdb.RangeOptions{
		Limit:   int(opt.PageSize),
		Reverse: opt.Reverse,
	})
}

func (e *Edge) ValidateIRI(p *perms.Edge) error {
	var err error

	if !validateSequence(e.Predicate) && !p.AllowPredicateWildcard {
		return &ParsingError{msg: "null record", field: "edge.predicate"}
	}

	if e.subjectKSUID, err = validateNodeID(e.Subject); err != nil {
		return &ParsingError{msg: "invalid Node ID", field: "edge.subject"}
	}

	if p.AllowTargetWildcard && e.Target == "*" {

	} else if e.targetKSUID, err = validateNodeID(e.Target); err != nil {
		return &ParsingError{msg: "invalid Node ID", field: "edge.target"}
	}

	return nil
}

func (e *Edge) ValidatePermission(p *perms.Edge) error {
	return nil
}
