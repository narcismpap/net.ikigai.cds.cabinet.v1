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

func TestIRIEdgeCompose(t *testing.T) {
	x := NewIRITester(t)

	e1 := iri.Edge{Subject: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Predicate: 1024, Target: "1EsJ4O4FlJXmKdNxLk52Go4x0uE"}
	x.path(e1.GetPath(), "e/1EsJ4OwOAdywg8iM3dnH2ODHfjq/1024/1EsJ4O4FlJXmKdNxLk52Go4x0uE")

	x.key(e1.Subject, "1EsJ4OwOAdywg8iM3dnH2ODHfjq", "subject")
	x.key(e1.Target, "1EsJ4O4FlJXmKdNxLk52Go4x0uE", "target")
	x.seqKey(e1.Predicate, 1024, "predicate")
}

func TestIRIEdgeParse(t *testing.T) {
	x := NewIRITester(t)

	e2 := iri.Edge{}
	if err := e2.Parse("e/1EsJ4OwOAdywg8iM3dnH2ODHfjq/1024/1EsJ4O4FlJXmKdNxLk52Go4x0uE"); err != nil{
		x.t.Log(err)
	}

	x.path(e2.GetPath(), "e/1EsJ4OwOAdywg8iM3dnH2ODHfjq/1024/1EsJ4O4FlJXmKdNxLk52Go4x0uE")

	x.key(e2.Subject, "1EsJ4OwOAdywg8iM3dnH2ODHfjq", "subject")
	x.key(e2.Target, "1EsJ4O4FlJXmKdNxLk52Go4x0uE", "target")
	x.seqKey(e2.Predicate, 1024, "predicate")
}
