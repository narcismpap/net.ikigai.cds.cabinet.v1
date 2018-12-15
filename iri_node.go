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

type IRINode struct{
	IRI
	Type uint16
	Id string
}

func (n *IRINode) getTypeK() string{
	return intToKeyElement(n.Type)
}

func (n *IRINode) GetPath() string{
	return fmt.Sprintf("/n/%d/%s", n.Type, n.Id)
}

func (n *IRINode) GetKey(server *CDSCabinetServer) fdb.Key{
	return server.dbNode.Sub(n.getTypeK()).Pack(tuple.Tuple{n.Id})
}

func (n *IRINode) GetClearRange(server *CDSCabinetServer) fdb.ExactRange{
	if n.Id == ""{
		return server.dbNode.Sub(n.getTypeK())
	}else{
		return server.dbEdge.Sub(n.getTypeK()).Sub(n.Id)
	}
}

func (e *IRINode) ValidateIRI() error{
	return nil
}

func (e *IRINode) ValidatePermission() error{
	return nil
}