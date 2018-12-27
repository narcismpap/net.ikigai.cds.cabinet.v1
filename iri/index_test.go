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

var tIndexBytes ExpectedBytes

func init(){
	// IRI: i/95/cat/1EsJ4OwOAdywg8iM3dnH2ODHfjq
	tIndexBytes.key = []byte{21, 18, 2, 105, 0, 1, 0, 255, 95, 0, 2, 99, 97, 116, 0, 2, 49, 69, 115, 74, 52, 79, 119, 79, 65, 100, 121, 119, 103, 56, 105, 77, 51, 100, 110, 72, 50, 79, 68, 72, 102, 106, 113, 0}
	tIndexBytes.start = []byte{}
	tIndexBytes.end = []byte{}
}

func TestIRIIndexCompose(t *testing.T) {
	x := NewIRITester(t)

	i1 := &iri.NodeIndex{Node: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", IndexId: 95, Value: "cat"}
	x.path(i1.GetPath(), "i/95/cat/1EsJ4OwOAdywg8iM3dnH2ODHfjq")

	x.key(i1.Node, "1EsJ4OwOAdywg8iM3dnH2ODHfjq", "node")
	x.key(i1.Value, "cat", "value")
	x.seqKey(i1.IndexId, 95, "index")

	x.bytes([]byte(i1.GetKey(testDb.DbIndex)), tIndexBytes.key, "GetKey()")
}

func TestIRIIndexParse(t *testing.T) {
	x := NewIRITester(t)

	i2 := &iri.NodeIndex{}
	if err := i2.Parse("i/95/cat/1EsJ4OwOAdywg8iM3dnH2ODHfjq"); err != nil{
		t.Error(err)
	}

	x.path(i2.GetPath(), "i/95/cat/1EsJ4OwOAdywg8iM3dnH2ODHfjq")

	x.key(i2.Node, "1EsJ4OwOAdywg8iM3dnH2ODHfjq", "node")
	x.key(i2.Value, "cat", "value")
	x.seqKey(i2.IndexId, 95, "index")

	x.bytes([]byte(i2.GetKey(testDb.DbIndex)), tIndexBytes.key, "GetKey()")
}

func TestIRIIndexBadSignature(t *testing.T) {
	x := NewIRITester(t)
	p := &perms.Index{}

	// nothing
	x.error((&iri.NodeIndex{

	}).ValidateIRI(p), "null record on index.IndexId")

	// bad + missing index
	x.error((&iri.NodeIndex{
		IndexId: 0, Value: "cat", Node: "1EsJ4OwOAdywg8iM3dnH2ODHfjq",
	}).ValidateIRI(p), "null record on index.IndexId")

	x.error((&iri.NodeIndex{
		Value: "cat", Node: "1EsJ4OwOAdywg8iM3dnH2ODHfjq",
	}).ValidateIRI(p), "null record on index.IndexId")

	// bad + missing node
	x.error((&iri.NodeIndex{
		IndexId: 100, Value: "dog", Node: "test",
	}).ValidateIRI(p), "invalid Node ID on index.node")

	x.error((&iri.NodeIndex{
		IndexId: 123, Value: "dog",
	}).ValidateIRI(p), "invalid Node ID on index.node")

	// bad + missing value
	x.error((&iri.NodeIndex{
		IndexId: 100, Node: "1EsJ4OwOAdywg8iM3dnH2ODHfjq",
	}).ValidateIRI(p), "null record on index.value")

	x.error((&iri.NodeIndex{
		IndexId: 123, Node: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Value: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam ut iaculis dolor. Donec convallis quam in convallis placerat. Aenean mollis et sapien vitae aliquam. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque habitant morbi tristique sectum.",
	}).ValidateIRI(p), "len > 256 on index.value")
}
