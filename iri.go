// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package main

import (
	"fmt"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
)

type IRISequential struct{
	Type string
	SeqID uint32
}

func (s *IRISequential) dbSeqID() string{
	return fmt.Sprintf("%05d", s.SeqID)
}

func (s *IRISequential) getPath() string{
	return fmt.Sprintf("/s/%s/%d", s.Type, s.SeqID)
}

func (s *IRISequential) getKey(server *CDSCabinetServer) fdb.Key{
	return server.dbSeq.Pack(tuple.Tuple{s.Type, s.dbSeqID()})
}

func (s *IRISequential) getIncrementKey(server *CDSCabinetServer) fdb.Key{
	return server.dbSeq.Pack(tuple.Tuple{"l", s.Type})
}
