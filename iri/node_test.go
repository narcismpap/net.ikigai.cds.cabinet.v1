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

var tNodeBytes ExpectedBytes

func init(){
	// IRI: n/128/1EsJ4OwOAdywg8iM3dnH2ODHfjq
	tNodeBytes.key = []byte{21, 18, 2, 110, 0, 1, 0, 255, 128, 0, 2, 49, 69, 115, 74, 52, 79, 119, 79, 65, 100, 121, 119, 103, 56, 105, 77, 51, 100, 110, 72, 50, 79, 68, 72, 102, 106, 113, 0}
	tNodeBytes.start = []byte{21, 18, 2, 110, 0, 1, 0, 255, 128, 0, 2, 49, 69, 115, 74, 52, 79, 119, 79, 65, 100, 121, 119, 103, 56, 105, 77, 51, 100, 110, 72, 50, 79, 68, 72, 102, 106, 113, 0, 0}
	tNodeBytes.end = []byte{21, 18, 2, 110, 0, 1, 0, 255, 128, 0, 2, 49, 69, 115, 74, 52, 79, 119, 79, 65, 100, 121, 119, 103, 56, 105, 77, 51, 100, 110, 72, 50, 79, 68, 72, 102, 106, 113, 0, 255}
}

func TestIRINodeCompose(t *testing.T) {
	x := NewIRITester(t)

	n1 := &iri.Node{Id: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Type: 128}
	x.path(n1.GetPath(), "n/128/1EsJ4OwOAdywg8iM3dnH2ODHfjq")

	x.key(n1.Id, "1EsJ4OwOAdywg8iM3dnH2ODHfjq", "node")
	x.seqKey(n1.Type, 128, "type")

	x.bytes([]byte(n1.GetKey(testDb.DbNode)), tNodeBytes.key, "GetKey()")

	rStart, rEnd := n1.GetClearRange(testDb.DbNode).FDBRangeKeys()
	x.bytes([]byte(rStart.FDBKey()), tNodeBytes.start, "GetClearRange(start)")
	x.bytes([]byte(rEnd.FDBKey()), tNodeBytes.end, "GetClearRange(end)")
}

func TestIRINodeParse(t *testing.T) {
	x := NewIRITester(t)

	n2 := &iri.Node{}
	if err := n2.Parse("n/128/1EsJ4OwOAdywg8iM3dnH2ODHfjq"); err != nil{
		t.Error(err)
	}

	x.path(n2.GetPath(), "n/128/1EsJ4OwOAdywg8iM3dnH2ODHfjq")

	x.key(n2.Id, "1EsJ4OwOAdywg8iM3dnH2ODHfjq", "node")
	x.seqKey(n2.Type, 128, "type")

	x.bytes([]byte(n2.GetKey(testDb.DbNode)), tNodeBytes.key, "GetKey()")

	rStart, rEnd := n2.GetClearRange(testDb.DbNode).FDBRangeKeys()
	x.bytes([]byte(rStart.FDBKey()), tNodeBytes.start, "GetClearRange(start)")
	x.bytes([]byte(rEnd.FDBKey()), tNodeBytes.end, "GetClearRange(end)")
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
