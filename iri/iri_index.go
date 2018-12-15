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

type IRINodeIndex struct{
	IRI

	Node string
	IndexId uint16
	Value string
}

func (i *IRINodeIndex) GetPath() string{
	return fmt.Sprintf("/i/%d/%s/%s", i.IndexId, i.Value, i.Node)
}

func (m *IRINodeIndex) getIndexK() string{
	return intToKeyElement(m.IndexId)
}

func (i *IRINodeIndex) GetKey(server *cds.CDSCabinetServer) fdb.Key{
	return server.DbIndex.Pack(tuple.Tuple{i.getIndexK(), i.Value, i.Node})
}

func (i *IRINodeIndex) GetClearRange(server *cds.CDSCabinetServer) fdb.ExactRange{
	return nil
}

func (e *IRINodeIndex) ValidateIRI() error{
	return nil
}

func (e *IRINodeIndex) ValidatePermission() error{
	return nil
}
