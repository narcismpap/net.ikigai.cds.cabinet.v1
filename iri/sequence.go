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
)

type Sequence struct{
	IRI
	Type string
	SeqID uint32
	UUID string
}

func (s *Sequence) DbSeqID() string{
	return fmt.Sprintf("%05d", s.SeqID)
}

func (s *Sequence) GetPath() string{
	if s.SeqID > 0 {
		return fmt.Sprintf("/s/%s/i/%d", s.Type, s.SeqID)
	}else{
		return fmt.Sprintf("/s/%s/u/%d", s.Type, s.SeqID)
	}
}

func (s *Sequence) GetKey(db subspace.Subspace) fdb.Key{
	return db.Pack(tuple.Tuple{s.Type, "i", s.DbSeqID()})
}

func (s *Sequence) GetReverseKey(db subspace.Subspace) fdb.Key{
	return db.Pack(tuple.Tuple{s.Type, "u", s.UUID})
}

func (s *Sequence) GetIncrementKey(db subspace.Subspace) fdb.Key{
	return db.Pack(tuple.Tuple{"l", s.Type})
}
