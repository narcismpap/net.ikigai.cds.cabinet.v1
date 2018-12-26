// Package: net.ikigai.cds
// Module: cabinet.services.test
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package iri_test

import (
	"cds.ikigai.net/cabinet.v1/iri"
	"fmt"
	"testing"
	"unsafe"
)

func TestIRICounterEdge(t *testing.T) {
	x := IRITester{t: t}

	c1 := iri.EdgeCounter{Subject: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Predicate: 10, Target: "1EsJ4O4FlJXmKdNxLk52Go4x0uE", Counter: 1200}

	x.path(c1.GetPath(), "/c/e/100/1EsJ4OwOAdywg8iM3dnH2ODHfjq/10/1EsJ4O4FlJXmKdNxLk52Go4x0uE")

	c2 := c1.GetCounterBytes()


}

type IRITester struct{
	t *testing.T
}


func (i *IRITester) path(a string, b string){
	if a != b{
		i.t.Errorf("Path mismatch: %s != %s", a, b)
	}
}

func (i *IRITester) key(a string, b string){
	if a != b{
		i.t.Errorf("Key mismatch: %s != %s", a, b)
	}
}