// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package iri

import (
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"strconv"
)

type IRI interface {
	GetPath() string
	GetKey(db subspace.Subspace) fdb.Key
	GetClearRange(db subspace.Subspace) fdb.ExactRange
}

func intToKeyElement(v uint16) string{
	return strconv.FormatUint(uint64(v), 36)
}

func KeyElementToInt(k string) (uint16, error){
	v, e := strconv.ParseUint(k, 36, 32)

	if e != nil{
		return 0, e
	}

	return uint16(v), nil
}
