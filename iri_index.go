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

type IRINodeIndex struct{
	IRI

	Node string
	IndexId uint16
	Value string
}

func (i *IRINodeIndex) getPath() string{
	return fmt.Sprintf("/i/%d/%s/%s", i.IndexId, i.Value, i.Node)
}

func (m *IRINodeIndex) getIndexK() string{
	return intToKeyElement(m.IndexId)
}

func (i *IRINodeIndex) getKey(server *CDSCabinetServer) fdb.Key{
	return server.dbIndex.Pack(tuple.Tuple{i.getIndexK(), i.Value, i.Node})
}

func (i *IRINodeIndex) getClearRange(server *CDSCabinetServer) fdb.ExactRange{
	return nil
}
