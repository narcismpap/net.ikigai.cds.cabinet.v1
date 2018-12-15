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
}

func (s *Sequence) DbSeqID() string{
	return fmt.Sprintf("%05d", s.SeqID)
}

func (s *Sequence) GetPath() string{
	return fmt.Sprintf("/s/%s/%d", s.Type, s.SeqID)
}

func (s *Sequence) GetKey(db subspace.Subspace) fdb.Key{
	return db.Pack(tuple.Tuple{s.Type, s.DbSeqID()})
}

func (s *Sequence) GetIncrementKey(db subspace.Subspace) fdb.Key{
	return db.Pack(tuple.Tuple{"l", s.Type})
}
