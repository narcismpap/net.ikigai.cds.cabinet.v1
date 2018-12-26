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

func TestIRICounterEdgeCompose(t *testing.T) {
	x := NewIRITester(t)

	c1 := iri.EdgeCounter{Subject: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Predicate: 10, Target: "1EsJ4O4FlJXmKdNxLk52Go4x0uE", Counter: 100}
	x.path(c1.GetPath(), "c/e/100/1EsJ4OwOAdywg8iM3dnH2ODHfjq/10/1EsJ4O4FlJXmKdNxLk52Go4x0uE")

	x.key(c1.Subject, "1EsJ4OwOAdywg8iM3dnH2ODHfjq", "subject")
	x.key(c1.Target, "1EsJ4O4FlJXmKdNxLk52Go4x0uE", "target")
	x.seqKey(c1.Predicate, 10, "predicate")
	x.seqKey(c1.Counter, 100, "counter")
}

func TestIRICounterEdgeParse(t *testing.T) {
	x := NewIRITester(t)

	c2 := iri.EdgeCounter{}
	if err := c2.Parse("c/e/100/1EsJ4OwOAdywg8iM3dnH2ODHfjq/10/1EsJ4O4FlJXmKdNxLk52Go4x0uE"); err != nil{
		x.t.Log(err)
	}

	x.path(c2.GetPath(), "c/e/100/1EsJ4OwOAdywg8iM3dnH2ODHfjq/10/1EsJ4O4FlJXmKdNxLk52Go4x0uE")

	x.key(c2.Subject, "1EsJ4OwOAdywg8iM3dnH2ODHfjq", "subject")
	x.key(c2.Target, "1EsJ4O4FlJXmKdNxLk52Go4x0uE", "target")
	x.seqKey(c2.Predicate, 10, "predicate")
	x.seqKey(c2.Counter, 100, "counter")
}


func TestIRICounterNodeCompose(t *testing.T) {
	x := NewIRITester(t)

	c3 := iri.NodeCounter{Node: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Counter: 100}
	x.path(c3.GetPath(), "c/n/100/1EsJ4OwOAdywg8iM3dnH2ODHfjq")

	x.key(c3.Node, "1EsJ4OwOAdywg8iM3dnH2ODHfjq", "node")
	x.seqKey(c3.Counter, 100, "counter")
}

func TestIRICounterNodeParse(t *testing.T) {
	x := NewIRITester(t)

	c4 := iri.NodeCounter{}
	if err := c4.Parse("c/n/100/1EsJ4OwOAdywg8iM3dnH2ODHfjq"); err != nil{
		x.t.Log(err)
	}

	x.path(c4.GetPath(), "c/n/100/1EsJ4OwOAdywg8iM3dnH2ODHfjq")

	x.key(c4.Node, "1EsJ4OwOAdywg8iM3dnH2ODHfjq", "node")
	x.seqKey(c4.Counter, 100, "counter")
}
