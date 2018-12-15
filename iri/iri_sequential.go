// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package iri

import (
	cds "cds.ikigai.net/cabinet.v1/server"
	"fmt"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
)

type IRISequential struct{
	IRI
	Type string
	SeqID uint32
}

func (s *IRISequential) DbSeqID() string{
	return fmt.Sprintf("%05d", s.SeqID)
}

func (s *IRISequential) GetPath() string{
	return fmt.Sprintf("/s/%s/%d", s.Type, s.SeqID)
}

func (s *IRISequential) GetKey(server *cds.CDSCabinetServer) fdb.Key{
	return server.DbSeq.Pack(tuple.Tuple{s.Type, s.DbSeqID()})
}

func (s *IRISequential) GetIncrementKey(server *cds.CDSCabinetServer) fdb.Key{
	return server.DbSeq.Pack(tuple.Tuple{"l", s.Type})
}
