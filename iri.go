// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package main

import (
	"github.com/apple/foundationdb/bindings/go/src/fdb"
)

type IRI interface {
	GetPath() string
	GetKey(server *CDSCabinetServer) fdb.Key
	GetClearRange(server *CDSCabinetServer) fdb.ExactRange

	ValidateIRI() error
	ValidatePermission() error
}

