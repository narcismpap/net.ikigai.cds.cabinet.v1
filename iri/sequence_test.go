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


func TestIRISequenceIdCompose(t *testing.T) {
	x := NewIRITester(t)

	s1 := iri.Sequence{Type: "q", SeqID: 66}
	x.path(s1.GetPath(), "s/q/66")

	x.key(s1.Type, "q", "type")
	x.key(s1.UUID, "", "uuid")
	x.seqKey32(s1.SeqID, 66, "seqId")
}

func TestIRISequenceIdParse(t *testing.T) {
	x := NewIRITester(t)

	s2 := iri.Sequence{}
	if err := s2.Parse("s/q/66"); err != nil{
		x.t.Log(err)
	}

	x.path(s2.GetPath(), "s/q/66")

	x.key(s2.Type, "q", "type")
	x.key(s2.UUID, "", "uuid")
	x.seqKey32(s2.SeqID, 66, "seqId")
}

func TestIRISequenceUuidCompose(t *testing.T) {
	x := NewIRITester(t)

	s3 := iri.Sequence{UUID: "19AEA360-E595-4B15-83A2-2CAB7FE31767"}
	x.path(s3.GetPath(), "su/19AEA360-E595-4B15-83A2-2CAB7FE31767")

	x.key(s3.Type, "", "type")
	x.key(s3.UUID, "19AEA360-E595-4B15-83A2-2CAB7FE31767", "uuid")
	x.seqKey32(s3.SeqID, 0, "seqId")
}

func TestIRISequenceUuidParse(t *testing.T) {
	x := NewIRITester(t)

	s4 := iri.Sequence{}
	if err := s4.Parse("su/19AEA360-E595-4B15-83A2-2CAB7FE31767"); err != nil{
		x.t.Log(err)
	}

	x.path(s4.GetPath(), "su/19AEA360-E595-4B15-83A2-2CAB7FE31767")

	x.key(s4.Type, "", "type")
	x.key(s4.UUID, "19AEA360-E595-4B15-83A2-2CAB7FE31767", "uuid")
	x.seqKey32(s4.SeqID, 0, "seqId")
}
