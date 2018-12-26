// Package: net.ikigai.cds
// Module: cabinet.services.test
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package iri_test

import (
	"cds.ikigai.net/cabinet.v1/iri"
	"testing"
)


func TestIRIIndexCompose(t *testing.T) {
	x := NewIRITester(t)

	i1 := iri.NodeIndex{Node: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", IndexId: 95, Value: "cat"}
	x.path(i1.GetPath(), "i/95/cat/1EsJ4OwOAdywg8iM3dnH2ODHfjq")

	x.key(i1.Node, "1EsJ4OwOAdywg8iM3dnH2ODHfjq", "node")
	x.key(i1.Value, "cat", "value")
	x.seqKey(i1.IndexId, 95, "index")
}

func TestIRIIndexParse(t *testing.T) {
	x := NewIRITester(t)

	i2 := iri.NodeIndex{}
	if err := i2.Parse("i/95/cat/1EsJ4OwOAdywg8iM3dnH2ODHfjq"); err != nil{
		x.t.Log(err)
	}

	x.path(i2.GetPath(), "i/95/cat/1EsJ4OwOAdywg8iM3dnH2ODHfjq")

	x.key(i2.Node, "1EsJ4OwOAdywg8iM3dnH2ODHfjq", "node")
	x.key(i2.Value, "cat", "value")
	x.seqKey(i2.IndexId, 95, "index")
}
