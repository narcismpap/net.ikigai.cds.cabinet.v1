// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package iri

import (
	"cds.ikigai.net/cabinet.v1/perms"
	pb "cds.ikigai.net/cabinet.v1/rpc"
	"errors"
	"fmt"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
)

type Sequence struct {
	IRI
	Type  string
	SeqID uint32
	UUID  string
}

func (s *Sequence) DbSeqID() string {
	return fmt.Sprintf("%05d", s.SeqID)
}

func (s *Sequence) GetPath() string {
	if s.SeqID > 0 {
		return fmt.Sprintf("/s/%s/%d", s.Type, s.SeqID)
	} else {
		return fmt.Sprintf("/su/%s", s.UUID)
	}
}

func (s *Sequence) Parse(path string) error {
	return errors.New("not implemented")
}

func (s *Sequence) GetKey(db subspace.Subspace) fdb.Key {
	return db.Pack(tuple.Tuple{s.Type, "i", s.DbSeqID()})
}

func (s *Sequence) GetReverseKey(db subspace.Subspace) fdb.Key {
	return db.Pack(tuple.Tuple{"su", s.UUID})
}

func (s *Sequence) GetIncrementKey(db subspace.Subspace) fdb.Key {
	return db.Pack(tuple.Tuple{"sl", s.Type})
}

func (s *Sequence) GetListRange(db subspace.Subspace, rtr fdb.ReadTransaction, opt *pb.ListOptions) fdb.RangeResult {
	readRange := db.Sub(s.Type).Sub("i")

	return rtr.GetRange(readRange, fdb.RangeOptions{
		Limit:   int(opt.PageSize),
		Reverse: opt.Reverse,
	})
}

func (s *Sequence) ValidateIRI(p *perms.Sequence) error {
	var err error

	if len(s.Type) == 0 {
		return &ParsingError{msg: "null record", field: "seq.type"}
	} else if len(s.Type) > 10 {
		return &ParsingError{msg: "len > 10", field: "seq.type"}
	}

	if len(s.UUID) > 0 && s.SeqID > 0 {
		return &ParsingError{msg: "mutually exclusive", field: "seq.uuid,seq.id"}
	}

	if len(s.UUID) == 0 && s.SeqID == 0 {
		return &ParsingError{msg: "id required", field: "seq.uuid|seq.id"}
	}

	if len(s.UUID) > 0 {
		if _, err = validateUUID(s.UUID); err != nil {
			return &ParsingError{msg: "invalid UUID", field: "seq.uuid"}
		}
	} else if s.SeqID == 0 {
		return errors.New("missing seqId")
	}

	return nil
}

func (s *Sequence) ValidatePermission(p perms.Sequence) error {
	return nil
}
