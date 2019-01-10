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

var tEdgeBytes ExpectedBytes

func init() {
	// IRI: e/1EsJ4OwOAdywg8iM3dnH2ODHfjq/1024/1EsJ4O4FlJXmKdNxLk52Go4x0uE
	tEdgeBytes.key = []byte{2, 49, 69, 115, 74, 52, 79, 119, 79, 65, 100, 121, 119, 103, 56, 105, 77, 51, 100, 110, 72, 50, 79, 68, 72, 102, 106, 113, 0, 1, 4, 0, 255, 0, 2, 49, 69, 115, 74, 52, 79, 52, 70, 108, 74, 88, 109, 75, 100, 78, 120, 76, 107, 53, 50, 71, 111, 52, 120, 48, 117, 69, 0}
	tEdgeBytes.start = []byte{2, 49, 69, 115, 74, 52, 79, 119, 79, 65, 100, 121, 119, 103, 56, 105, 77, 51, 100, 110, 72, 50, 79, 68, 72, 102, 106, 113, 0, 1, 4, 0, 255, 0, 0}
	tEdgeBytes.end = []byte{2, 49, 69, 115, 74, 52, 79, 119, 79, 65, 100, 121, 119, 103, 56, 105, 77, 51, 100, 110, 72, 50, 79, 68, 72, 102, 106, 113, 0, 1, 4, 0, 255, 0, 255}
}

func TestIRIEdgeCompose(t *testing.T) {
	x := NewIRITester(t)

	e1 := &iri.Edge{Subject: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Predicate: 1024, Target: "1EsJ4O4FlJXmKdNxLk52Go4x0uE"}
	x.path(e1.GetPath(), "e/1EsJ4OwOAdywg8iM3dnH2ODHfjq/1024/1EsJ4O4FlJXmKdNxLk52Go4x0uE")

	x.key(e1.Subject, "1EsJ4OwOAdywg8iM3dnH2ODHfjq", "subject")
	x.key(e1.Target, "1EsJ4O4FlJXmKdNxLk52Go4x0uE", "target")
	x.seqKey(e1.Predicate, 1024, "predicate")

	x.bytes([]byte(e1.GetKey(testDb.DbEdge)), tEdgeBytes.getKey(testDb.DbEdgeBytes), "GetKey()")

	rStart, rEnd := e1.GetClearRange(testDb.DbEdge).FDBRangeKeys()
	x.bytes([]byte(rStart.FDBKey()), tEdgeBytes.getStart(testDb.DbEdgeBytes), "GetClearRange(start)")
	x.bytes([]byte(rEnd.FDBKey()), tEdgeBytes.GetEnd(testDb.DbEdgeBytes), "GetClearRange(end)")
}

func TestIRIEdgeParse(t *testing.T) {
	x := NewIRITester(t)

	e2 := &iri.Edge{}
	if err := e2.Parse("e/1EsJ4OwOAdywg8iM3dnH2ODHfjq/1024/1EsJ4O4FlJXmKdNxLk52Go4x0uE"); err != nil {
		t.Error(err)
	}

	x.path(e2.GetPath(), "e/1EsJ4OwOAdywg8iM3dnH2ODHfjq/1024/1EsJ4O4FlJXmKdNxLk52Go4x0uE")

	x.key(e2.Subject, "1EsJ4OwOAdywg8iM3dnH2ODHfjq", "subject")
	x.key(e2.Target, "1EsJ4O4FlJXmKdNxLk52Go4x0uE", "target")
	x.seqKey(e2.Predicate, 1024, "predicate")

	x.bytes([]byte(e2.GetKey(testDb.DbEdge)), tEdgeBytes.getKey(testDb.DbEdgeBytes), "GetKey()")

	rStart, rEnd := e2.GetClearRange(testDb.DbEdge).FDBRangeKeys()
	x.bytes([]byte(rStart.FDBKey()), tEdgeBytes.getStart(testDb.DbEdgeBytes), "GetClearRange(start)")
	x.bytes([]byte(rEnd.FDBKey()), tEdgeBytes.GetEnd(testDb.DbEdgeBytes), "GetClearRange(end)")
}

func TestIRIEdgeBadSignature(t *testing.T) {
	x := NewIRITester(t)
	p := &perms.Edge{}

	// nothing
	x.error((&iri.Edge{}).ValidateIRI(p), "null record on edge.predicate")

	// no subject + bad
	x.error((&iri.Edge{
		Predicate: 100, Target: "1EsJ4O4FlJXmKdNxLk52Go4x0uE",
	}).ValidateIRI(p), "invalid Node ID on edge.subject")

	x.error((&iri.Edge{
		Subject: "i am a test", Predicate: 100, Target: "1EsJ4O4FlJXmKdNxLk52Go4x0uE",
	}).ValidateIRI(p), "invalid Node ID on edge.subject")

	// no predicate + bad
	x.error((&iri.Edge{
		Subject: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Target: "1EsJ4O4FlJXmKdNxLk52Go4x0uE",
	}).ValidateIRI(p), "null record on edge.predicate")

	// wildcard predicate
	x.error((&iri.Edge{
		Subject: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Predicate: 0, Target: "1EsJ4O4FlJXmKdNxLk52Go4x0uE",
	}).ValidateIRI(p), "null record on edge.predicate")

	// no target + bad
	x.error((&iri.Edge{
		Subject: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Predicate: 100,
	}).ValidateIRI(p), "invalid Node ID on edge.target")

	x.error((&iri.Edge{
		Subject: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Predicate: 100, Target: "I like cats",
	}).ValidateIRI(p), "invalid Node ID on edge.target")

	// wildcard target (not allowed)
	x.error((&iri.Edge{
		Subject: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Predicate: 100, Target: "*",
	}).ValidateIRI(p), "invalid Node ID on edge.target")

	// wildcard pred + target
	x.error((&iri.Edge{
		Subject: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Predicate: 0, Target: "*",
	}).ValidateIRI(p), "null record on edge.predicate")
}

func TestIRIEdgeBadSignatureWildcard(t *testing.T) {
	x := NewIRITester(t)

	// wildcard target
	p1 := &perms.Edge{AllowPredicateWildcard: false, AllowTargetWildcard: true}

	x.nil((&iri.Edge{
		Subject: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Predicate: 100, Target: "*",
	}).ValidateIRI(p1))

	x.error((&iri.Edge{
		Subject: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Predicate: 0, Target: "1EsJ4O4FlJXmKdNxLk52Go4x0uE",
	}).ValidateIRI(p1), "null record on edge.predicate")

	// wildcard predicate
	p2 := &perms.Edge{AllowPredicateWildcard: true, AllowTargetWildcard: false}

	x.nil((&iri.Edge{
		Subject: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Predicate: 0, Target: "1EsJ4OwOAdywg8iM3dnH2ODHfjq",
	}).ValidateIRI(p2))

	x.error((&iri.Edge{
		Subject: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Predicate: 0, Target: "*",
	}).ValidateIRI(p2), "invalid Node ID on edge.target")

	// wildcard pred + target
	p3 := &perms.Edge{AllowPredicateWildcard: true, AllowTargetWildcard: true}

	x.nil((&iri.Edge{
		Subject: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Predicate: 0, Target: "*",
	}).ValidateIRI(p3))

	x.error((&iri.Edge{
		Predicate: 0, Target: "*",
	}).ValidateIRI(p3), "invalid Node ID on edge.subject")
}
