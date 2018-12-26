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


func TestIRISequenceIdCompose(t *testing.T) {
	x := NewIRITester(t)

	s1 := &iri.Sequence{Type: "q", SeqID: 66}
	x.path(s1.GetPath(), "s/q/66")

	x.key(s1.Type, "q", "type")
	x.key(s1.UUID, "", "uuid")
	x.seqKey32(s1.SeqID, 66, "seqId")
}

func TestIRISequenceIdParse(t *testing.T) {
	x := NewIRITester(t)

	s2 := &iri.Sequence{}
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

	s3 := &iri.Sequence{UUID: "19AEA360-E595-4B15-83A2-2CAB7FE31767"}
	x.path(s3.GetPath(), "su/19AEA360-E595-4B15-83A2-2CAB7FE31767")

	x.key(s3.Type, "", "type")
	x.key(s3.UUID, "19AEA360-E595-4B15-83A2-2CAB7FE31767", "uuid")
	x.seqKey32(s3.SeqID, 0, "seqId")
}

func TestIRISequenceUuidParse(t *testing.T) {
	x := NewIRITester(t)

	s4 := &iri.Sequence{}
	if err := s4.Parse("su/19AEA360-E595-4B15-83A2-2CAB7FE31767"); err != nil{
		x.t.Log(err)
	}

	x.path(s4.GetPath(), "su/19AEA360-E595-4B15-83A2-2CAB7FE31767")

	x.key(s4.Type, "", "type")
	x.key(s4.UUID, "19AEA360-E595-4B15-83A2-2CAB7FE31767", "uuid")
	x.seqKey32(s4.SeqID, 0, "seqId")
}


func TestIRISequenceBadSignature(t *testing.T) {
	x := NewIRITester(t)
	p := &perms.Sequence{}

	// empty
	x.error((&iri.Sequence{

	}).ValidateIRI(p), "null record on seq.type")

	// bad type, missing & over-sized
	x.error((&iri.Sequence{
		SeqID: 100,
	}).ValidateIRI(p), "null record on seq.type")

	x.error((&iri.Sequence{
		SeqID: 100, Type: "",
	}).ValidateIRI(p), "null record on seq.type")

	x.error((&iri.Sequence{
		SeqID: 100, Type: "cats are quite fun",
	}).ValidateIRI(p), "len > 10 on seq.type")

	// seqid & uuid cannot co-exist
	x.error((&iri.Sequence{
		SeqID: 100, UUID: "5D5D892A-C4EA-40F9-9724-5B1CD5A68057", Type: "test",
	}).ValidateIRI(p), "mutually exclusive on seq.uuid,seq.id")

	// missing ID & UUID
	x.error((&iri.Sequence{
		Type: "test",
	}).ValidateIRI(p), "id required on seq.uuid|seq.id")

	// bad UUID
	x.error((&iri.Sequence{
		Type: "test", UUID: "cats are fun",
	}).ValidateIRI(p), "invalid UUID on seq.uuid")
}