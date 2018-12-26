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


func TestIRINodeCompose(t *testing.T) {
	x := NewIRITester(t)

	n1 := iri.Node{Id: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Type: 128}
	x.path(n1.GetPath(), "n/128/1EsJ4OwOAdywg8iM3dnH2ODHfjq")

	x.key(n1.Id, "1EsJ4OwOAdywg8iM3dnH2ODHfjq", "node")
	x.seqKey(n1.Type, 128, "type")
}

func TestIRINodeParse(t *testing.T) {
	x := NewIRITester(t)

	n2 := iri.Node{}
	if err := n2.Parse("n/128/1EsJ4OwOAdywg8iM3dnH2ODHfjq"); err != nil{
		x.t.Log(err)
	}

	x.path(n2.GetPath(), "n/128/1EsJ4OwOAdywg8iM3dnH2ODHfjq")

	x.key(n2.Id, "1EsJ4OwOAdywg8iM3dnH2ODHfjq", "node")
	x.seqKey(n2.Type, 128, "type")
}
