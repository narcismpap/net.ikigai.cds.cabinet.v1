// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package iri_test

import "testing"

type IRITester struct{
	t *testing.T
}

func NewIRITester(t *testing.T) *IRITester {
	ir := &IRITester{t: t}
	t.Parallel()
	return ir
}


func (i *IRITester) path(a string, b string){
	if a != b{
		i.t.Errorf("Path mismatch: %s != %s", a, b)
	}
}

func (i *IRITester) key(a string, b string, name string){
	if a != b{
		i.t.Errorf("Key mismatch (.%s): %s != %s", name, a, b)
	}
}

func (i *IRITester) seqKey(a uint16, b uint16, name string){
	if a != b{
		i.t.Errorf("Key mismatch (.%s): %d != %d", name, a, b)
	}
}

func (i *IRITester) seqKey32(a uint32, b uint32, name string){
	if a != b{
		i.t.Errorf("Key mismatch (.%s): %d != %d", name, a, b)
	}
}

func (i *IRITester) error(err error, expected string){
	if err == nil || string(err.Error()) != expected{
		i.t.Errorf("Expected [%s] error, got [%v]", expected, err)
		panic(err)
	}
}

func (i *IRITester) nil(err error){
	if err != nil{
		i.t.Errorf("Expected nil error, got [%v]", err)
		//panic(err)
	}
}
