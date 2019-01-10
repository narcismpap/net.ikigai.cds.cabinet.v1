// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package iri

import (
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
)

type IRI interface {
	GetPath() string
	GetKey(db subspace.Subspace) fdb.Key
	GetClearRange(db subspace.Subspace) fdb.ExactRange
	GetListRange(db subspace.Subspace, rtr fdb.ReadTransaction, opt *ListOptions) fdb.RangeResult
}

type ListOptions struct {
	PageSize int
	Reverse  bool
}
