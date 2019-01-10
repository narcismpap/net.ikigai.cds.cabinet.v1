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

func TestIRIMetaEdgeCompose(t *testing.T) {
	x := NewIRITester(t)

	expBytes := &ExpectedBytes{
		[]byte{2, 101, 0, 2, 49, 69, 115, 104, 73, 113, 81, 65, 102, 121, 77, 113, 85, 71, 77, 52, 107, 106, 51, 119, 103, 111, 55, 118, 55, 90, 81, 0, 1, 0, 255, 12, 0, 2, 49, 69, 115, 104, 73, 112, 107, 69, 69, 116, 88, 81, 97, 52, 75, 86, 105, 100, 67, 107, 65, 50, 110, 81, 80, 86, 100, 0, 1, 0, 255, 243, 0},
		[]byte{2, 101, 0, 2, 49, 69, 115, 104, 73, 113, 81, 65, 102, 121, 77, 113, 85, 71, 77, 52, 107, 106, 51, 119, 103, 111, 55, 118, 55, 90, 81, 0, 1, 0, 255, 12, 0, 2, 49, 69, 115, 104, 73, 112, 107, 69, 69, 116, 88, 81, 97, 52, 75, 86, 105, 100, 67, 107, 65, 50, 110, 81, 80, 86, 100, 0, 0},
		[]byte{2, 101, 0, 2, 49, 69, 115, 104, 73, 113, 81, 65, 102, 121, 77, 113, 85, 71, 77, 52, 107, 106, 51, 119, 103, 111, 55, 118, 55, 90, 81, 0, 1, 0, 255, 12, 0, 2, 49, 69, 115, 104, 73, 112, 107, 69, 69, 116, 88, 81, 97, 52, 75, 86, 105, 100, 67, 107, 65, 50, 110, 81, 80, 86, 100, 0, 255},
	}

	c1 := &iri.EdgeMeta{Subject: "1EshIqQAfyMqUGM4kj3wgo7v7ZQ", Predicate: 12, Target: "1EshIpkEEtXQa4KVidCkA2nQPVd", Property: 243}
	x.path(c1.GetPath(), "m/e/1EshIqQAfyMqUGM4kj3wgo7v7ZQ/12/1EshIpkEEtXQa4KVidCkA2nQPVd/243")

	x.key(c1.Subject, "1EshIqQAfyMqUGM4kj3wgo7v7ZQ", "subject")
	x.key(c1.Target, "1EshIpkEEtXQa4KVidCkA2nQPVd", "target")
	x.seqKey(c1.Predicate, 12, "predicate")
	x.seqKey(c1.Property, 243, "property")

	x.bytes([]byte(c1.GetKey(testDb.DbMeta)), expBytes.getKey(testDb.DbMetaBytes), "GetKey()")

	rStart, rEnd := c1.GetClearRange(testDb.DbMeta).FDBRangeKeys()
	x.bytes([]byte(rStart.FDBKey()), expBytes.getStart(testDb.DbMetaBytes), "GetClearRange(start)")
	x.bytes([]byte(rEnd.FDBKey()), expBytes.GetEnd(testDb.DbMetaBytes), "GetClearRange(end)")
}

func TestIRIMetaEdgeParse(t *testing.T) {
	x := NewIRITester(t)

	expBytes := &ExpectedBytes{
		[]byte{2, 101, 0, 2, 49, 69, 115, 104, 73, 113, 99, 69, 109, 107, 52, 72, 120, 110, 119, 112, 90, 83, 100, 83, 70, 102, 77, 113, 109, 97, 116, 0, 1, 0, 255, 56, 0, 2, 49, 69, 115, 104, 73, 110, 116, 90, 101, 74, 78, 49, 117, 98, 68, 67, 98, 85, 89, 83, 53, 122, 65, 49, 110, 111, 78, 0, 1, 20, 124, 0},
		[]byte{2, 101, 0, 2, 49, 69, 115, 104, 73, 113, 99, 69, 109, 107, 52, 72, 120, 110, 119, 112, 90, 83, 100, 83, 70, 102, 77, 113, 109, 97, 116, 0, 1, 0, 255, 56, 0, 2, 49, 69, 115, 104, 73, 110, 116, 90, 101, 74, 78, 49, 117, 98, 68, 67, 98, 85, 89, 83, 53, 122, 65, 49, 110, 111, 78, 0, 0},
		[]byte{2, 101, 0, 2, 49, 69, 115, 104, 73, 113, 99, 69, 109, 107, 52, 72, 120, 110, 119, 112, 90, 83, 100, 83, 70, 102, 77, 113, 109, 97, 116, 0, 1, 0, 255, 56, 0, 2, 49, 69, 115, 104, 73, 110, 116, 90, 101, 74, 78, 49, 117, 98, 68, 67, 98, 85, 89, 83, 53, 122, 65, 49, 110, 111, 78, 0, 255},
	}

	c2 := &iri.EdgeMeta{}
	if err := c2.Parse("m/e/1EshIqcEmk4HxnwpZSdSFfMqmat/56/1EshIntZeJN1ubDCbUYS5zA1noN/5244"); err != nil {
		t.Error(err)
	}

	x.path(c2.GetPath(), "m/e/1EshIqcEmk4HxnwpZSdSFfMqmat/56/1EshIntZeJN1ubDCbUYS5zA1noN/5244")

	x.key(c2.Subject, "1EshIqcEmk4HxnwpZSdSFfMqmat", "subject")
	x.key(c2.Target, "1EshIntZeJN1ubDCbUYS5zA1noN", "target")
	x.seqKey(c2.Predicate, 56, "predicate")
	x.seqKey(c2.Property, 5244, "property")

	x.bytes([]byte(c2.GetKey(testDb.DbMeta)), expBytes.getKey(testDb.DbMetaBytes), "GetKey()")

	rStart, rEnd := c2.GetClearRange(testDb.DbMeta).FDBRangeKeys()
	x.bytes([]byte(rStart.FDBKey()), expBytes.getStart(testDb.DbMetaBytes), "GetClearRange(start)")
	x.bytes([]byte(rEnd.FDBKey()), expBytes.GetEnd(testDb.DbMetaBytes), "GetClearRange(end)")
}

func TestIRIMetaEdgeBadSignature(t *testing.T) {
	x := NewIRITester(t)
	p := &perms.Meta{}

	// empty
	x.error((&iri.EdgeMeta{}).ValidateIRI(p), "null record on meta.property")

	// no prop
	x.error((&iri.EdgeMeta{
		Subject: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Predicate: 100, Target: "1EsJ4O4FlJXmKdNxLk52Go4x0uE",
	}).ValidateIRI(p), "null record on meta.property")

	// wildcard prop (not allowed)
	x.error((&iri.EdgeMeta{
		Subject: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Predicate: 100, Target: "1EsJ4O4FlJXmKdNxLk52Go4x0uE", Property: 0,
	}).ValidateIRI(p), "null record on meta.property")

	// no subject + bad
	x.error((&iri.EdgeMeta{
		Predicate: 100, Target: "1EsJ4O4FlJXmKdNxLk52Go4x0uE", Property: 143,
	}).ValidateIRI(p), "invalid Node ID on meta.edge.subject")

	x.error((&iri.EdgeMeta{
		Subject: "i am a test", Predicate: 100, Target: "1EsJ4O4FlJXmKdNxLk52Go4x0uE", Property: 255,
	}).ValidateIRI(p), "invalid Node ID on meta.edge.subject")

	// no predicate + bad
	x.error((&iri.EdgeMeta{
		Subject: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Target: "1EsJ4O4FlJXmKdNxLk52Go4x0uE", Property: 5243,
	}).ValidateIRI(p), "null record on meta.edge.predicate")

	x.error((&iri.EdgeMeta{
		Subject: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Predicate: 0, Target: "1EsJ4O4FlJXmKdNxLk52Go4x0uE", Property: 524,
	}).ValidateIRI(p), "null record on meta.edge.predicate")

	// no target + bad
	x.error((&iri.EdgeMeta{
		Subject: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Predicate: 100, Property: 24,
	}).ValidateIRI(p), "invalid Node ID on meta.edge.target")

	x.error((&iri.EdgeMeta{
		Subject: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Predicate: 100, Property: 876, Target: "I like cats",
	}).ValidateIRI(p), "invalid Node ID on meta.edge.target")
}

func TestIRIMetaEdgeBadSignatureWildcard(t *testing.T) {
	x := NewIRITester(t)

	p1 := &perms.Meta{AllowWildcardProperty: false}
	p2 := &perms.Meta{AllowWildcardProperty: true}

	x.error((&iri.EdgeMeta{
		Subject: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Predicate: 100, Target: "1EsJ4O4FlJXmKdNxLk52Go4x0uE", Property: 0,
	}).ValidateIRI(p1), "null record on meta.property")

	x.nil((&iri.EdgeMeta{
		Subject: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Predicate: 100, Target: "1EsJ4O4FlJXmKdNxLk52Go4x0uE", Property: 0,
	}).ValidateIRI(p2))
}

func TestIRIMetaNodeCompose(t *testing.T) {
	x := NewIRITester(t)

	expBytes := &ExpectedBytes{
		[]byte{2, 110, 0, 2, 49, 69, 115, 104, 73, 112, 107, 80, 113, 100, 84, 48, 50, 81, 54, 77, 80, 107, 104, 119, 52, 108, 87, 80, 77, 74, 107, 0, 1, 2, 12, 0},
		[]byte{2, 110, 0, 2, 49, 69, 115, 104, 73, 112, 107, 80, 113, 100, 84, 48, 50, 81, 54, 77, 80, 107, 104, 119, 52, 108, 87, 80, 77, 74, 107, 0, 0},
		[]byte{2, 110, 0, 2, 49, 69, 115, 104, 73, 112, 107, 80, 113, 100, 84, 48, 50, 81, 54, 77, 80, 107, 104, 119, 52, 108, 87, 80, 77, 74, 107, 0, 255},
	}

	c3 := &iri.NodeMeta{Node: "1EshIpkPqdT02Q6MPkhw4lWPMJk", Property: 524}
	x.path(c3.GetPath(), "m/n/1EshIpkPqdT02Q6MPkhw4lWPMJk/524")

	x.key(c3.Node, "1EshIpkPqdT02Q6MPkhw4lWPMJk", "node")
	x.seqKey(c3.Property, 524, "property")

	x.bytes([]byte(c3.GetKey(testDb.DbMeta)), expBytes.getKey(testDb.DbMetaBytes), "GetKey()")

	rStart, rEnd := c3.GetClearRange(testDb.DbMeta).FDBRangeKeys()
	x.bytes([]byte(rStart.FDBKey()), expBytes.getStart(testDb.DbMetaBytes), "GetClearRange(start)")
	x.bytes([]byte(rEnd.FDBKey()), expBytes.GetEnd(testDb.DbMetaBytes), "GetClearRange(end)")
}

func TestIRIMetaNodeParse(t *testing.T) {
	x := NewIRITester(t)

	expBytes := &ExpectedBytes{
		[]byte{2, 110, 0, 2, 49, 69, 115, 104, 73, 109, 116, 119, 119, 54, 122, 71, 100, 120, 103, 49, 99, 115, 66, 89, 80, 106, 103, 113, 57, 89, 89, 0, 1, 211, 227, 0},
		[]byte{2, 110, 0, 2, 49, 69, 115, 104, 73, 109, 116, 119, 119, 54, 122, 71, 100, 120, 103, 49, 99, 115, 66, 89, 80, 106, 103, 113, 57, 89, 89, 0, 0},
		[]byte{2, 110, 0, 2, 49, 69, 115, 104, 73, 109, 116, 119, 119, 54, 122, 71, 100, 120, 103, 49, 99, 115, 66, 89, 80, 106, 103, 113, 57, 89, 89, 0, 255},
	}

	c4 := &iri.NodeMeta{}
	if err := c4.Parse("m/n/1EshImtww6zGdxg1csBYPjgq9YY/54243"); err != nil {
		t.Error(err)
	}

	x.path(c4.GetPath(), "m/n/1EshImtww6zGdxg1csBYPjgq9YY/54243")

	x.key(c4.Node, "1EshImtww6zGdxg1csBYPjgq9YY", "node")
	x.seqKey(c4.Property, 54243, "property")

	x.bytes([]byte(c4.GetKey(testDb.DbMeta)), expBytes.getKey(testDb.DbMetaBytes), "GetKey()")

	rStart, rEnd := c4.GetClearRange(testDb.DbMeta).FDBRangeKeys()
	x.bytes([]byte(rStart.FDBKey()), expBytes.getStart(testDb.DbMetaBytes), "GetClearRange(start)")
	x.bytes([]byte(rEnd.FDBKey()), expBytes.GetEnd(testDb.DbMetaBytes), "GetClearRange(end)")
}

func TestIRIMetaNodeBadSignature(t *testing.T) {
	x := NewIRITester(t)
	p := &perms.Meta{}

	// empty
	x.error((&iri.NodeMeta{}).ValidateIRI(p), "null record on meta.property")

	// no property
	x.error((&iri.NodeMeta{
		Node: "1EsJ4OwOAdywg8iM3dnH2ODHfjq",
	}).ValidateIRI(p), "null record on meta.property")

	// wildcard property (not allowed)
	x.error((&iri.NodeMeta{
		Node: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Property: 0,
	}).ValidateIRI(p), "null record on meta.property")

	// no valid ID
	x.error((&iri.NodeMeta{
		Node: "not good", Property: 100,
	}).ValidateIRI(p), "invalid Node ID on meta.node")

	// invalid size
	x.error((&iri.NodeMeta{
		Node: "1EsJ4OwOAdywg8iM3dnH2ODHfjqQ", Property: 100,
	}).ValidateIRI(p), "invalid Node ID on meta.node")

	// UUID
	x.error((&iri.NodeMeta{
		Node: "151E0C62-1ADD-4BA0-BFC5-840E9370592F", Property: 100,
	}).ValidateIRI(p), "invalid Node ID on meta.node")
}

func TestIRIMetaNodeBadSignatureWildcard(t *testing.T) {
	x := NewIRITester(t)

	p1 := &perms.Meta{AllowWildcardProperty: false}
	p2 := &perms.Meta{AllowWildcardProperty: true}

	x.error((&iri.NodeMeta{
		Node: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Property: 0,
	}).ValidateIRI(p1), "null record on meta.property")

	x.nil((&iri.NodeMeta{
		Node: "1EshIohhwTq3pya0tLVNCST07gN", Property: 0,
	}).ValidateIRI(p2))
}
