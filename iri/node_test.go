// Package: net.ikigai.cds
// Module: cabinet.services.test
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package iri_test

import (
	"cds.ikigai.net/cabinet.v1/iri"
	"cds.ikigai.net/cabinet.v1/perms"
	"testing"
)


func TestIRINodeCompose(t *testing.T) {
	x := NewIRITester(t)

	n1 := &iri.Node{Id: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Type: 128}
	x.path(n1.GetPath(), "n/128/1EsJ4OwOAdywg8iM3dnH2ODHfjq")

	x.key(n1.Id, "1EsJ4OwOAdywg8iM3dnH2ODHfjq", "node")
	x.seqKey(n1.Type, 128, "type")
}

func TestIRINodeParse(t *testing.T) {
	x := NewIRITester(t)

	n2 := &iri.Node{}
	if err := n2.Parse("n/128/1EsJ4OwOAdywg8iM3dnH2ODHfjq"); err != nil{
		x.t.Log(err)
	}

	x.path(n2.GetPath(), "n/128/1EsJ4OwOAdywg8iM3dnH2ODHfjq")

	x.key(n2.Id, "1EsJ4OwOAdywg8iM3dnH2ODHfjq", "node")
	x.seqKey(n2.Type, 128, "type")
}


func TestIRINodeBadSignature(t *testing.T) {
	x := NewIRITester(t)
	p := &perms.Node{}

	// empty
	x.error((&iri.Node{

	}).ValidateIRI(p), "null record on node.type")

	// missing + bad type
	x.error((&iri.Node{
		Id: "1EsJ4OwOAdywg8iM3dnH2ODHfjq",
	}).ValidateIRI(p), "null record on node.type")

	x.error((&iri.Node{
		Id: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Type: 0,
	}).ValidateIRI(p), "null record on node.type")

	// missing + bad Id
	x.error((&iri.Node{
		Type: 153,
	}).ValidateIRI(p), "invalid Node ID on node.id")

	x.error((&iri.Node{
		Id: "cat", Type: 123,
	}).ValidateIRI(p), "invalid Node ID on node.id")

	x.error((&iri.Node{
		Id: "010064E9-E09D-43BF-AC85-BE265EB85886", Type: 123,
	}).ValidateIRI(p), "invalid Node ID on node.id")
}
