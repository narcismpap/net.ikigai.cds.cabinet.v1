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
)

/* Edges*/
type EdgeMeta struct{
	IRI

	Subject string
	Predicate uint16
	Target string

	subjectKSUID ksuid.KSUID
	targetKSUID ksuid.KSUID

	Property uint16
}

func (m *EdgeMeta) GetPath() string{
	return fmt.Sprintf("/m/e/%s/%d/%s/%d", m.Subject, m.Predicate, m.Target, m.Property)
}

func (m *EdgeMeta) getPropertyK() string{
	return IntToKeyElement(m.Property)
}

func (m *EdgeMeta) getPredicateK() string{
	return IntToKeyElement(m.Predicate)
}

func (m *EdgeMeta) GetKey(db subspace.Subspace) fdb.Key{
	return db.Sub("e").Pack(tuple.Tuple{m.Subject, m.getPredicateK(), m.Target, m.getPropertyK()})
}

func (m *EdgeMeta) GetClearRange(db subspace.Subspace) fdb.ExactRange{
	if m.Target != ""{
		return db.Sub("e").Sub(m.Subject).Sub(m.getPredicateK()).Sub(m.Target)
	}

	if m.Predicate > 0{
		return db.Sub("e").Sub(m.Subject).Sub(m.Predicate)
	}

	return db.Sub("e").Sub(m.Subject)
}

func (m *EdgeMeta) GetListRange(db subspace.Subspace, rtr fdb.ReadTransaction, opt *pb.ListOptions) fdb.RangeResult{
	readRange := db.Sub("e").Sub(m.Subject).Sub(m.getPredicateK()).Sub(m.Target)

	return rtr.GetRange(readRange, fdb.RangeOptions{
		Limit: 	 int(opt.PageSize),
		Reverse: opt.Reverse,
	})
}

func (m *EdgeMeta) ValidateIRI(p *perms.Meta) error{
	var err error

	if !validateSequence(m.Property) && !p.AllowWildcardProperty{
		return &ParsingError{msg: "null record", field: "meta.property"}
	}else if !validateSequence(m.Predicate){
		return &ParsingError{msg: "null record", field: "meta.edge.predicate"}
	}

	if m.subjectKSUID, err = validateNodeID(m.Subject); err != nil{
		return &ParsingError{msg: "invalid Node ID", field: "meta.edge.subject"}
	}

	if m.targetKSUID, err = validateNodeID(m.Target); err != nil{
		return &ParsingError{msg: "invalid Node ID", field: "meta.edge.target"}
	}

	return nil
}

func (e *EdgeMeta) ValidatePermission(p perms.Meta) error{
	return nil
}


/* Nodes */
type NodeMeta struct{
	Node string
	Property uint16

	nodeKSUID ksuid.KSUID
}

func (m *NodeMeta) GetPath() string{
	return fmt.Sprintf("/m/n/%s/%d", m.Node, m.Property)
}

func (m *NodeMeta) getPropertyK() string{
	return IntToKeyElement(m.Property)
}

func (m *NodeMeta) GetKey(db subspace.Subspace) fdb.Key{
	return db.Sub("n").Pack(tuple.Tuple{m.Node, m.getPropertyK()})
}

func (m *NodeMeta) GetClearRange(db subspace.Subspace) fdb.ExactRange{
	return db.Sub("n").Sub(m.Node)
}

func (m *NodeMeta) GetListRange(db subspace.Subspace, rtr fdb.ReadTransaction, opt *pb.ListOptions) fdb.RangeResult{
	readRange := db.Sub("n").Sub(m.Node)

	return rtr.GetRange(readRange, fdb.RangeOptions{
		Limit: 	 int(opt.PageSize),
		Reverse: opt.Reverse,
	})
}

func (m *NodeMeta) ValidateIRI(p *perms.Meta) error{
	var err error

	if !validateSequence(m.Property) && !p.AllowWildcardProperty{
		return &ParsingError{msg: "null record", field: "meta.property"}
	}

	if m.nodeKSUID, err = validateNodeID(m.Node); err != nil{
		return &ParsingError{msg: "invalid Node ID", field: "meta.node"}
	}

	return nil
}

func (e *NodeMeta) ValidatePermission(p perms.Meta) error{
	return nil
}
